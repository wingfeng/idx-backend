package models

import (
	"strings"

	"github.com/wingfeng/backend/utils"

	"gorm.io/gorm"
)

type MenuItem struct {
	ID             string     `json:"id" gorm:"primary_key;column:Id;type:varchar(36);not null"`
	Name           string     `json:"text" gorm:"type:varchar(255)"`
	Icon           string     `json:"icon" gorm:"type:varchar(255)"`
	URL            string     `json:"url" gorm:"type:varchar(255)"`
	Component      string     `json:"component" gorm:"type:varchar(255)"`
	SortOrder      int        `json:"sortorder" `
	Operations     string     `json:"operations" gorm:"type:varchar(255)"`
	Path           string     `json:"path" gorm:"type:varchar(2048)"`
	Parent         string     `json:"parent" gorm:"type:varchar(36)"`
	Code           string     `json:"code" gorm:"unique"`
	Hidden         bool       `json:"hidden"`
	Children       []MenuItem `json:"nodes" gorm:"-"`
	RoleOperations string     `json:"role_operations" gorm:"-"`
	utils.Record   `gorm:"embedded"`
}

//TableName 数据表名称
func (m *MenuItem) TableName() string {
	return "Menus"
}
func (m *MenuItem) ParentID() interface{} {
	return m.Parent
}
func (m *MenuItem) SetID(id interface{}) {
	m.ID = id.(string)
}
func (m *MenuItem) GetID() interface{} {
	return m.ID
}
func (m *MenuItem) SetChildren(children []interface{}) {
	m.Children = make([]MenuItem, 0)
	for _, c := range children {
		m.Children = append(m.Children, *c.(*MenuItem))
	}
}
func (r *MenuItem) BeforeCreate(tx *gorm.DB) error {
	r.Record.BeforeCreate(tx)
	if strings.EqualFold(r.Parent, "") {
		r.Path = r.ID
	}

	return nil
}

func (r *MenuItem) BeforeUpdate(tx *gorm.DB) error {
	r.Record.BeforeUpdate(tx)
	if strings.EqualFold(r.Parent, "") {
		r.Path = r.ID
	}
	return nil
}
