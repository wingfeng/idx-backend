package base

import (
	"fmt"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Row interface {
	GetID() interface{}
	//	SetID(id interface{})
	TableName() string
}

// Record 数据库通用结构
type Record struct {
	Creator   string         `json:"creator" gorm:"type:varchar(255)"`
	CreatorID string         `json:"creatorid" gorm:"type:varchar(36)"`
	Updator   string         `json:"updator" gorm:"type:varchar(255)"`
	UpdatorID string         `json:"updatorid" gorm:"type:varchar(36)"`
	CreatedAt null.Time      `json:"created" gorm:"autoCreateTime"`
	UpdatedAt null.Time      `json:"updated" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// GetID 获取当前记录的ID
func (r *Record) GetID() interface{} {
	panic(fmt.Errorf("model不支持GetID,%v", r))
	// val := reflect.ValueOf(r).Elem()
	// id := val.FieldByName("Id")
	// if id == reflect.ValueOf(nil) {
	// 	id = val.FieldByName("ID")
	// }
	// if id != reflect.ValueOf(nil) {
	// 	return id.Interface()
	// }
	// return nil
}

// SetID 获取当前记录的ID
func (r *Record) SetID(v interface{}) {
	panic(fmt.Errorf("请实现SetID方法,%v", r))
}
func (r *Record) BeforeCreate(tx *gorm.DB) error {
	u, _ := tx.Get(Const_UserNameKey)
	uID, _ := tx.Get(Const_UserIDKey)
	r.Creator = fmt.Sprintf("%v", u)
	r.CreatorID = fmt.Sprintf("%v", uID)
	return nil
}

func (r *Record) BeforeUpdate(tx *gorm.DB) error {
	u, _ := tx.Get(Const_UserNameKey)
	uID, _ := tx.Get(Const_UserIDKey)
	r.Updator = fmt.Sprintf("%v", u)
	r.UpdatorID = fmt.Sprintf("%v", uID)
	return nil
}
