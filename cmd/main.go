package main

import (
	"bufio"
	"html/template"
	"io/fs"
	"log/slog"
	"mime"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
	backend "github.com/wingfeng/idxadmin"

	"flag"
	"fmt"

	idxmodels "github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/system/models"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
	_ "github.com/wingfeng/idxadmin/docs"
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
	backend.EntryOption //`mapstructure:"backend"`
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
		panic(fmt.Errorf("读取配置文件错误: %s ", err))
	}
	viper.SetEnvPrefix("IDX")
	viper.AutomaticEnv()

	opts := &Opts{}
	err = viper.Unmarshal(opts)
	if err != nil {
		slog.Error("读取配置错误:", "error", err)
	}
	entryOption := &backend.EntryOption{}

	err = viper.Unmarshal(entryOption)
	// opts.EntryOption = entryOption
	if err != nil {
		slog.Error("读取配置错误:", "error", err)
	}

	driver := opts.Driver

	if *syncDb {
		connection := opts.Connection
		//初始化DB
		dbEngine := base.GetDB(driver, connection)
		models.Sync2Db(dbEngine)
		fmt.Println("同步数据库完成")
		return

	}
	if *gen {
		slog.Info("Beging generate controller and test")
		// row := idxmodels.Client{}
		// genController(row, "", "oauth2")
		apiRow := idxmodels.APIResources{}
		genController(apiRow, "", "oauth2")
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

	// route.Static("/", "../front/dist")

	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5666", "http://192.168.0.101:8080"},
		AllowMethods: []string{"PUT", "DELETE", "GET", "POST"},
		//	AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		//	MaxAge:           12 * time.Hour,
	}))

	backend.Init(entryOption, route)

	//	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	node, _ := snowflake.NewNode(1)
	fmt.Printf("程序实例启动 %v\r\n", node.Generate())
	loadFrontend(route)
	route.Run(fmt.Sprintf("%s:%d", opts.IP, opts.Port))

}
func loadFrontend(router *gin.Engine) {
	fil := "../front/"
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".mjs", "application/javascript")
	router.Static("/front", fil)
	router.GET("/front", func(c *gin.Context) {
		p := path.Join(fil, "index.html")
		c.File(p)
	})
	router.NoRoute(func(c *gin.Context) {
		p := path.Join(fil, "index.html")
		c.File(p)
	})
}
func genController(row interface{}, shortName string, module string) {
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
	o := fmt.Sprintf("../%s/controller/%scontroller.go", module, shortName)
	otest := fmt.Sprintf("../test/%s_%s_test.go", module, shortName)
	err := os.MkdirAll(path.Join("..", module, "controller"), fs.ModeAppend)
	os.MkdirAll(path.Join("..", "test", module), fs.ModeAppend)
	if err != nil {
		slog.Error("create directory error:", "err", err.Error())
	}
	slog.Info("生成文件:", "file", o)
	gen(gens, "../templates/controller.tpl", o)
	gen(gens, "../templates/test.tpl", otest)
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
