package test

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/rbac"
	"github.com/wingfeng/backend/routers"
	"github.com/wingfeng/backend/utils"
)

var route *gin.Engine

func setupRouter() *gin.Engine {
	if route == nil {
		connection := "root:eATq1GDhsP@tcp(localhost:3306)/idx?&parseTime=true"
		//初始化DB

		enf := rbac.InitEnforcer("mysql", connection, "../policy/rbac_model.conf")
		//初始化Gin
		route = gin.Default()
		api := route.Group("/api/v1/system")
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // 禁用彩色打印
			},
		)
		//api.Use(oidc.AuthWare())
		api.Use(func(c *gin.Context) {
			biz := utils.InitContext("mysql", connection, "员工", "1", newLogger)

			c.Set(utils.Const_UserIDKey, "1")
			c.Set(utils.Const_UserNameKey, "员工")
			c.Set(utils.Const_BizContextKey, biz)
			c.Set("casbin", enf)
		})
		fmt.Println("Hello World!")
		routers.RegisterRouter(api)
	}

	return route
}

func GetContext() *utils.BizContext {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	connection := "root:123456@tcp(localhost:3306)/OrgDb?parseTime=true"
	biz := utils.InitContext("mysql", connection, "fenggr", "1", newLogger)
	return biz
}
