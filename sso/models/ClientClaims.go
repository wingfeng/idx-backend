package models

import "github.com/wingfeng/backend/utils"

// ClientClaims [...]
type ClientClaims struct {
	ID           int     `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Type         string  `gorm:"column:Type;type:varchar(250);not null"`
	Value        string  `gorm:"column:Value;type:varchar(250);not null"`
	ClientID     int     `gorm:"index:IX_ClientClaims_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Clients `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

func (m *ClientClaims) TableName() string {
	return "ClientClaims"
}
