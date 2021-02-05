package backend

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/oidc"
	"github.com/wingfeng/backend/rbac"
	"github.com/wingfeng/backend/routers"
	"github.com/wingfeng/backend/utils"
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
		option.GroupName = "/api/v2/system"
	}
	api := route.Group(option.GroupName)
	issuer := option.JWTKeySet
	userEndpoint := option.UserEndpoint

	api.Use(func(c *gin.Context) {
		//在Context内构建Biz Context保证并发时信息不回冲突
		biz := utils.InitContext(option.Driver, option.Connection, "", "", nil)
		c.Set(utils.Const_BizContextKey, biz)
		c.Set(utils.Const_CasbinKey, enf)
	})
	if option.EnableOidc {
		oidc.Init(issuer, userEndpoint)
		api.Use(oidc.AuthWare())
	}

	routers.RegisterRouter(api)

}
