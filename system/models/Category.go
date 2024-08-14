package models

type Category struct {
	ID          string `gorm:"primary_key;column:id;type:varchar(36);not null"`
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Code        string `gorm:"column:code;unique;index;type:varchar(36);not null"`
	Parent      string `gorm:"column:parent;type:varchar(36);"`
	Path        string `gorm:"column:path;type:varchar(1024);"`
	SortOrder   int    `gorm:"column:sort_order;type:int;"`
	Description string `gorm:"column:description;type:varchar(1024);"`
}

func (r *Category) GetID() interface{} {
	return r.ID
}

// func (r *Category) TableName() string {
// 	return "Categories"
// }
