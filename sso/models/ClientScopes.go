package models

import "github.com/wingfeng/backend/utils"

// ClientScopes [...]
type ClientScopes struct {
	ID           int     `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Scope        string  `gorm:"column:Scope;type:varchar(200);not null"`
	ClientID     int     `gorm:"index:IX_ClientScopes_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Clients `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

func (m *ClientScopes) TableName() string {
	return "ClientScopes"
}
