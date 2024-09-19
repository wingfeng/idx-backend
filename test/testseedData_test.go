package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	idxutil "github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
	systemmodels "github.com/wingfeng/idxadmin/system/models"
)

var (
	context *base.BizContext
)

func TestSync2Db(t *testing.T) {
	context := GetBizContext()
	err := context.DB().AutoMigrate(
		&systemmodels.Category{},
		&systemmodels.MenuItem{},
		&systemmodels.OptionSet{},
	)
	if err != nil {
		t.Log(err)
		panic(err)
	}
	assert.Equal(t, nil, err)
}
func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)

	//初始化DB
	context = GetBizContext()

	ou := &models.OrganizationUnit{}
	ou.Id = "1328680589330485248"
	ou.Name = "XX软件"
	ou.DisplayName = "XX软件有限公司"
	db := context.DB()
	err := db.Save(ou).Error
	if err != nil {
		panic(err)
	}

	user := &models.User{}
	user.Id = "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	user.Account = "admin"
	user.DisplayName = "管理员"
	user.Email = "admin@fire.loc"
	user.OUId = ou.Id
	user.OU = ou.DisplayName

	user.PasswordHash, _ = idxutil.HashPassword("password1")
	db = context.DB()
	err = db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b"
	role.Name = "SystemAdmin"
	addRole(role)
	addUserRole(user.Id, ou.Id, role.Id)
	role = &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43c"
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.Id, ou.Id, role.Id)

	role = &models.Role{}
	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43d"
	role.Name = "admin"
	addRole(role)
	addUserRole(user.Id, ou.Id, role.Id)
	err = seedMenu()
	assert.Equal(t, nil, err)

	if err != nil {
		panic(err)
	}

}

func addUserRole(uid, ouid, rid string) {
	db := context.DB()
	ur := &models.UserRoles{
		RoleId: rid,
		UserId: uid,
		OUId:   ouid,
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
	m := &systemmodels.MenuItem{}

	m.ID = "1328680589439537152"
	m.Path = "/1"
	m.Name = "系统管理"
	m.Code = "system_mng"
	m.URL = "system_mng"
	m.Icon = "menu_system2"
	m.SortOrder = 10000

	tx := db.Save(m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
	}
	m = &systemmodels.MenuItem{}
	m.ID = "1328680589439537153"
	m.Path = "/1"
	m.Name = "系统设置"
	m.Code = "system_setting"
	m.URL = "system_mng"
	m.Icon = "menu_system2"
	m.Parent = "1328680589439537152"
	m.SortOrder = 10000

	tx = db.Save(m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
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
	tx = db.Save(m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
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
	tx = db.Save(m)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return tx.Error
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
	err := db.Save(m).Error
	return err
}
