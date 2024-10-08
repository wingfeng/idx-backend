package models

import (
	"log/slog"

	"gorm.io/gorm"
)

// Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	x.DisableForeignKeyConstraintWhenMigrating = true

	// 同步结构体与数据表
	err := x.AutoMigrate(

		new(Category),
		new(MenuItem),

		new(OptionSet),
	)
	if err != nil {
		slog.Error("同步数据结构错误,Error:%v", err)
	}
}
