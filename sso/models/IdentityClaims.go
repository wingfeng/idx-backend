package models

import "github.com/wingfeng/backend/utils"

// IDentityClaims [...]
type IDentityClaims struct {
	ID                 int               `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Type               string            `gorm:"column:Type;type:varchar(200);not null"`
	IDentityResourceID int               `gorm:"index:IX_IdentityClaims_IdentityResourceId;column:IdentityResourceId;type:int(11);not null"`
	IDentityResources  IDentityResources `gorm:"association_foreignkey:IdentityResourceId;foreignkey:Id"`
	utils.Record       `gorm:"embedded"`
}

//TableName 数据表名称
func (m *IDentityClaims) TableName() string {
	return "IdentityClaims"
}
