package models

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

//Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	x.DisableForeignKeyConstraintWhenMigrating = true

	// 同步结构体与数据表
	err := x.AutoMigrate(

		new(Category),
		new(MenuItem),
		new(OrganizationUnit),
		new(RoleClaims),
		new(Role),
		new(User),
		new(UserRoles),
		new(OptionSet),
	)
	if err != nil {
		log.Errorf("同步数据结构错误,Error:%v", err)
	}
}
