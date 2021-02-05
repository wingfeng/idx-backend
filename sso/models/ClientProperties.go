package models

import "github.com/wingfeng/backend/utils"

// ClientProperties [...]
type ClientProperties struct {
	ID           int     `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Key          string  `gorm:"column:Key;type:varchar(250);not null"`
	Value        string  `gorm:"column:Value;type:varchar(2000);not null"`
	ClientID     int     `gorm:"index:IX_ClientProperties_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Clients `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

func (m *ClientProperties) TableName() string {
	return "ClientProperties"
}
