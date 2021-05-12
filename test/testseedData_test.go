package test

import (
	"testing"

	"github.com/wingfeng/backend/system/models"
	"github.com/wingfeng/backend/utils"
	sso "github.com/wingfeng/idx/models"

	"github.com/magiconair/properties/assert"
)

var (
	context *utils.BizContext
)

func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)

	//初始化DB
	context = utils.InitContext("mysql", "root:123456@tcp(localhost:3306)/sso-v2?&parseTime=true", "fenggr", "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e", nil)
	ou := &models.OrganizationUnit{}
	ou.ID = "1328680589330485248"
	ou.Name = "XX软件"
	ou.DisplayName = "XX软件有限公司"
	db := context.DB()
	err := db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	user.ID = "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	user.UserName = "admin"
	user.DisplayName = "管理员"
	user.Email = "admin@fire.loc"
	user.OUID = ou.ID
	user.OU = ou.DisplayName

	user.PasswordHash = utils.GenHashedPWD("fire@123")
	db = context.DB()
	err = db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.ID = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b"
	role.Name = "SystemAdmin"
	addRole(role)
	addUserRole(user.ID, ou.ID, role.ID)
	role = &models.Role{}

	role.ID = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43c"
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.ID, ou.ID, role.ID)

	err = seedMenu()
	assert.Equal(t, nil, err)
	client := &sso.Client{
		ID:                               1,
		ClientID:                         "jsclient1",
		Enabled:                          true,
		ProtocolType:                     "oidc",
		RequireClientSecret:              false,
		ClientName:                       "Javascript Client",
		RequireConsent:                   true,
		AllowRememberConsent:             true,
		AlwaysIncludeUserClaimsInIDToken: false,
		AllowAccessTokensViaBrowser:      true,
		BackChannelLogoutSessionRequired: true,
		IDentityTokenLifetime:            300,
		AccessTokenLifetime:              3600,
		AbsoluteRefreshTokenLifetime:     2592000,
		SlidingRefreshTokenLifetime:      2592000,
		AuthorizationCodeLifetime:        300,
		RefreshTokenUsage:                1,
		RefreshTokenExpiration:           1,
		ClientClaimsPrefix:               "client_",
		DeviceCodeLifetime:               300,

		EnableLocalLogin: true,
		//UserSsoLifetime: , can be zero
	}
	db = context.DB()
	db.Save(client)
	cg := &sso.ClientGrantTypes{
		ID:        1,
		ClientID:  1,
		GrantType: "implicit",
	}
	db = context.DB()
	err = db.Save(cg).Error
	if err != nil {
		panic(err)
	}

	addRedirectURI(1, "http://localhost:9000/#/CallBack?", 1)
	addRedirectURI(2, "http://localhost:9000", 1)
	addClientScope(1, "openid", 1)
	addClientScope(2, "profile", 1)
	addClientScope(3, "roles", client.ID)
}
func addRedirectURI(id int, uri string, clientid int) {

	redUris := &sso.ClientRedirectURIs{
		ID:          id,
		RedirectURI: uri,
		ClientID:    clientid,
	}
	db := context.DB()
	err := db.Save(redUris).Error
	if err != nil {

		panic(err)
	}
}
func addClientScope(id int, scope string, clientid int) {
	sc := &sso.ClientScopes{
		ID:       id,
		Scope:    scope,
		ClientID: clientid,
	}
	db := context.DB()
	err := db.Save(sc).Error
	if err != nil {
		panic(err)
	}
}
func addUserRole(uid, ouid, rid string) {
	db := context.DB()
	ur := &models.UserRoles{
		RoleID: rid,
		UserID: uid,
		OUID:   ouid,
	}
	//联合主键的直接用engine来处理
	err := db.Save(ur).Error
	if err != nil {
		panic(err)
	}
}
func addRole(role *models.Role) {
	db := context.DB()
	err := db.Save(role).Error
	if err != nil {
		panic(err)
	}
}
func seedMenu() error {
	db := context.DB()
	//node, err := snowflake.NewNode(1)
	//	id := "1328680589439537153" //node.Generate().String()
	m := &models.MenuItem{}

	m.ID = "1328680589439537152"
	m.Path = "/1"
	m.Name = "系统管理"
	m.Code = "system_mng"
	m.URL = "system_mng"
	m.Icon = "menu_system2"
	m.SortOrder = 10000

	err := db.Save(m).Error
	if err != nil {
		return err
	}

	m.ID = "1328680589439537153"
	m.Path = "/1"
	m.Name = "系统设置"
	m.Code = "system_setting"
	m.URL = "system_mng"
	m.Icon = "menu_system2"
	m.Parent = "1328680589439537152"
	m.SortOrder = 10000

	err = db.Save(m).Error
	if err != nil {
		return err
	}

	//	context.Engine.ID("1328680589439537153").Get(m)

	m.ID = "1328680589452120064"
	m.Path = "/1/2"
	m.Name = "菜单管理"
	m.Code = "system_menu"
	m.URL = "/system/menu"
	m.Icon = "menu_system2"
	m.Component = "system/menu"
	m.Parent = "1328680589439537153"
	m.SortOrder = 10000
	err = db.Save(m).Error
	if err != nil {
		return err
	}
	//	context.Engine.ID("1328680589464702976").Get(m)

	m.ID = "1328680589464702976"
	m.Path = "/1/3"
	m.Name = "用户管理"
	m.Code = "system_user"
	m.URL = "/system/user"
	m.Icon = "menu_system2"
	m.Component = "system/user"
	m.Parent = "1328680589439537153"
	m.SortOrder = 10000
	err = db.Save(m).Error
	if err != nil {
		return err
	}
	// m = &models.MenuItem{
	// 	Path:      "/4",
	// 	ID:        "1328680589473091584",
	// 	Name:      "首页",
	// 	Code:      "home",
	// 	URL:       "home",
	// 	Icon:      "menu_home_white",
	// 	Component: "home/default",
	// 	SortOrder: -1,
	// }
	//	context.Engine.ID("1328680589473091584").Get(m)

	m.ID = "1328680589473091584"
	m.Path = "/4"
	m.Name = "首页"
	m.Code = "home"
	m.URL = "/home/"
	m.Icon = "menu_home_white"
	m.Component = "home/default"
	m.Parent = ""
	m.SortOrder = -1
	err = db.Save(m).Error
	return err
}
