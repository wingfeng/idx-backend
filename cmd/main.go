package main

import (
	"bufio"
	"html/template"
	"os"
	"reflect"
	"strings"

	"github.com/lunny/log"
	"github.com/spf13/viper"
	"github.com/wingfeng/backend"
	"github.com/wingfeng/backend/system/models"
	"github.com/wingfeng/backend/utils"

	// _ "/docs"
	"flag"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

type Opts struct {
	backend.EntryOption `mapstructure:"backend"`
	Port                int
	IP                  string
	SnowflakeNode       int
}

func main() {
	confPath := flag.String("conf", "../conf/config.yaml", "配置文件路径")
	showVersion := flag.Bool("ver", false, "程序版本")
	syncDb := flag.Bool("syncdb", false, "同步数据结构到数据库.")
	gen := flag.Bool("gen", false, "生成Controller和Test")
	flag.Parse()

	if *showVersion {
		Version()
		return
	}
	viper.SetConfigFile(*confPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AllowEmptyEnv(true)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("读取配置文件错误: %v \r\n", err))
	}
	viper.SetEnvPrefix("QQ")
	viper.AutomaticEnv()

	opts := &Opts{}
	err = viper.Unmarshal(opts)
	if err != nil {
		log.Error("读取配置错误:", err)
	}

	driver := opts.Driver

	if *syncDb {
		connection := opts.Connection
		//初始化DB
		dbEngine := utils.GetDB(driver, connection)
		models.Sync2Db(dbEngine)
		fmt.Println("同步数据库完成")
		return

	}
	if *gen {
		// row := idxmodels.ClientGrantTypes{}
		// genController(row, "")
		// apiRow := models.APIResources{}
		// genController(apiRow, "")
		// grant := models.PersistedGrants{}
		// genController(grant, "")
		// clientsecrets := models.ClientSecrets{}
		// genController(clientsecrets, "")
		// Clientproperties := models.ClientProperties{}
		// genController(Clientproperties, "")
		// Apiproperties := models.APIProperties{}
		// genController(Apiproperties, "")
		// Apisecrets := models.APISecrets{}
		// genController(Apisecrets, "")
		return
	}
	//初始化Gin
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000", "http://192.168.0.106:8080", "http://192.168.0.101:8080"},
		AllowMethods:     []string{"*", "PUT", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		//	MaxAge:           12 * time.Hour,
	}))
	group := route.Group("/")
	option := opts.EntryOption
	//  backend.EntryOption{
	// 	Driver:     driver,
	// 	Connection: connection,
	// 	Group:      group,
	// 	GroupName:  "api/v1/system",
	// 	EnableOidc: true,
	// 	PolicyPath: "../policy/rbac_model.conf",
	// }
	option.Group = group
	backend.Init(option)
	//swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	node, _ := snowflake.NewNode(1)
	fmt.Printf("程序实例启动 %v\r\n", node.Generate())
	route.Run(fmt.Sprintf("%s:%d", opts.IP, opts.Port))

}

func genController(row interface{}, shortName string) {
	tp := reflect.TypeOf(row)

	name := tp.String()
	if strings.EqualFold(shortName, "") {
		names := strings.Split(name, ".")
		shortName = names[len(names)-1]
	}
	gens := genStruct{
		Type:           name,
		ShortName:      shortName,
		LowerShortName: strings.ToLower(shortName),
	}
	gen(gens, "../templates/controller.tpl", "../system/controller/"+shortName+"controller.go")
	gen(gens, "../templates/test.tpl", "../test/"+shortName+"_test.go")
}

type genStruct struct {
	Type           string
	ShortName      string
	LowerShortName string
}

func Version() {
	fmt.Printf("App Name:\t%s\n", AppName)
	fmt.Printf("App Version:\t%s\n", AppVersion)
	fmt.Printf("Build version:\t%s\n", BuildVersion)
	fmt.Printf("Build time:\t%s\n", BuildTime)
	fmt.Printf("Git revision:\t%s\n", GitRevision)
	fmt.Printf("Git branch:\t%s\n", GitBranch)
	fmt.Printf("Golang Version: %s\n", GoVersion)
}

func gen(genS genStruct, tplFile string, outputName string) {
	tpl, err := template.ParseFiles(tplFile)
	// funcMap := template.FuncMap{"StrLower": lower}
	// tpl.Funcs(funcMap)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(strings.ToLower(outputName), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	wr := bufio.NewWriter(f)

	tpl.Execute(wr, genS)
	wr.Flush()
}
