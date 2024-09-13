package backend

import (
	"github.com/gin-gonic/gin"
	//	oauth2 "github.com/wingfeng/idx-oauth2"
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/oauth2"
	"github.com/wingfeng/idxadmin/oidc"
	"github.com/wingfeng/idxadmin/rbac"
	"github.com/wingfeng/idxadmin/system"
)

type EntryOption struct {
	Driver        string
	Connection    string
	RootGroupName string

	//是否在Group内直接注入OIDC的中间件
	EnableOidc   bool
	PolicyPath   string
	JWTKeySet    string
	UserEndpoint string
}

func Init(option EntryOption, route gin.IRouter) {

	//初始化Casbin RBAC
	enf := rbac.InitEnforcer(option.Driver, option.Connection, option.PolicyPath)
	rm := oidc.NewRoleManager(10)
	enf.SetRoleManager(rm)
	enf.LoadPolicy()

	issuer := option.JWTKeySet
	userEndpoint := option.UserEndpoint

	if option.EnableOidc {
		oidc.Init(issuer, userEndpoint)
		route.Use(oidc.AuthWare())
	}
	route.Use(func(c *gin.Context) {
		//在Context内构建Biz Context保证并发时信息不回冲突
		biz := base.InitContext(option.Driver, option.Connection, "", "", nil)

		c.Set(base.Const_BizContextKey, biz)
		c.Set(base.Const_CasbinKey, enf)
	})
	//初始化Router
	apiRoot := route.Group(option.RootGroupName)
	systemGroup := apiRoot.Group("system")

	system.RegisterRouter(systemGroup)
	oauthGroup := apiRoot.Group("oauth2")
	oauth2.RegisterRouter(oauthGroup)
}
