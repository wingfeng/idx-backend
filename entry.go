package backend

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/oidc"
	"github.com/wingfeng/idxadmin/rbac"
	"github.com/wingfeng/idxadmin/routers"
)

type EntryOption struct {
	Driver     string
	Connection string
	GroupName  string
	Group      *gin.RouterGroup
	//是否在Group内直接注入OIDC的中间件
	EnableOidc   bool
	PolicyPath   string
	JWTKeySet    string
	UserEndpoint string
}

func Init(option EntryOption) {

	//初始化Casbin RBAC
	enf := rbac.InitEnforcer(option.Driver, option.Connection, option.PolicyPath)
	rm := oidc.NewRoleManager(10)
	enf.SetRoleManager(rm)
	enf.LoadPolicy()
	route := option.Group

	if strings.EqualFold(option.GroupName, "") {
		option.GroupName = "/v1/system"
	}
	api := route.Group(option.GroupName)
	issuer := option.JWTKeySet
	userEndpoint := option.UserEndpoint

	api.Use(func(c *gin.Context) {
		//在Context内构建Biz Context保证并发时信息不回冲突
		biz := base.InitContext(option.Driver, option.Connection, "", "", nil)
		c.Set(base.Const_BizContextKey, biz)
		c.Set(base.Const_CasbinKey, enf)
	})
	if option.EnableOidc {
		oidc.Init(issuer, userEndpoint)
		api.Use(oidc.AuthWare())
	}

	routers.RegisterRouter(api)

}
