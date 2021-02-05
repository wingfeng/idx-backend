package models

import "github.com/wingfeng/backend/utils"

// ClientCorsOrigins [...]
type ClientCorsOrigins struct {
	ID           int     `gorm:"primary_key;autoIncrement;column:Id;type:int(11);not null"`
	Origin       string  `gorm:"column:Origin;type:varchar(150);not null"`
	ClientID     int     `gorm:"index:IX_ClientCorsOrigins_ClientId;column:ClientId;type:int(11);not null"`
	Clients      Clients `gorm:"association_foreignkey:ClientId;foreignkey:Id"`
	utils.Record `gorm:"embedded"`
}

func (m *ClientCorsOrigins) TableName() string {
	return "ClientCorsOrigins"
}
