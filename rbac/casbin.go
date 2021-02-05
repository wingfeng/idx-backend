package rbac

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/log"
	_ "github.com/mattn/go-sqlite3"
)

//InitEnforcer 初始化casbin的Enforcer
func InitEnforcer(driver string, connection string, path string) *casbin.Enforcer {
	adapter, err := xormadapter.NewAdapter(driver, connection, true)
	if err != nil {
		panic(fmt.Sprintf("adapter err is %v", err.Error()))
	}
	//初始化权限配置
	enforcer, err := casbin.NewEnforcer(path, adapter)

	if err != nil {
		panic(fmt.Sprintf("NewEnforcer err is %v", err.Error()))
	}
	return enforcer
}

//RbacHandle 权限认证中间件
func RbacHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		//rLock.RLock()
		//defer rLock.RUnlock()
		e, _ := context.Get("casbin")
		enf := e.(*casbin.Enforcer)

		//获取请求接口
		reqUrl := context.Request.URL.Path
		//获取请求方法
		name := context.GetString("name")
		//判断用户是否有权限
		ok, _ := enf.Enforce(name, reqUrl, "0")

		//判断策略中是否存在
		if ok {
			context.Next()
			return

		}

		log.Debug("权限验证不通过")
		context.JSON(http.StatusForbidden, gin.H{
			"messages": http.StatusText(http.StatusForbidden),
			"errInfo":  "抱歉, 您无权访问",
		})
		context.Abort()
	}
}
