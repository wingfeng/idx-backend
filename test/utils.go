package test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/rbac"
	"github.com/wingfeng/idxadmin/routers"
	"gorm.io/gorm/logger"
)

var route *gin.Engine

func setupRouter() *gin.Engine {
	if route == nil {
		connection := "host=localhost port=5432 user=root password=pass@word1 dbname=idx sslmode=disable TimeZone=Asia/Shanghai"
		//初始化DB

		enf := rbac.InitEnforcer("postgres", connection, "../policy/rbac_model.conf")
		//初始化Gin
		route = gin.Default()
		api := route.Group("/api/v1/system")

		//api.Use(oidc.AuthWare())
		api.Use(func(c *gin.Context) {
			biz := GetBizContext()

			c.Set(base.Const_BizContextKey, biz)
			c.Set("casbin", enf)
		})
		fmt.Println("Hello World!")
		routers.RegisterRouter(api)
	}

	return route
}

func GetBizContext() *base.BizContext {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	connection := "host=localhost port=5432 user=root password=pass@word1 dbname=idx sslmode=disable TimeZone=Asia/Shanghai"
	biz := base.InitContext("pgx", connection, "fenggr", "1", newLogger)
	return biz
}
