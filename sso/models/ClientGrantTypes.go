package models

import "github.com/wingfeng/backend/utils"

// ClientGrantTypes [...]
type ClientGrantTypes struct {
	ID           int     `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	GrantType    string  `gorm:"column:GrantType;type:varchar(250);not null"`
	ClientID     int     `gorm:"index:IX_ClientGrantTypes_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Clients `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

func (m *ClientGrantTypes) TableName() string {
	return "ClientGrantTypes"
}
