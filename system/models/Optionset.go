package models

import "github.com/wingfeng/idxadmin/base"

//OptionSet 数据字典
type OptionSet struct {
	ID string `json:"id" gorm:"primary_key;column:id;type:varchar(36);not null"`
	//Name 显示名
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Code        string `json:"code" gorm:"index;type:varchar(255);not null"`
	Value       string `json:"value" gorm:"type:varchar(255);not null"`
	base.Record `json:"-" gorm:"embedded"`
}

//TableName 数据表名称
// func (m *OptionSet) TableName() string {
// 	return "OptionSets"
// }
