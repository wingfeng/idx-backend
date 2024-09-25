package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idxadmin/base"
)

func TestBizContext_PageComplex(t *testing.T) {
	context := GetBizContext()
	db := context.DB().Table("users").Select("users.*,org.display_name").Joins("join organization_units as org on users.ou_id=org.id").Where("users.account='admin'")
	//	data := make(map[string]interface{}, 0)
	page := &base.Page{
		Filters:  []string{""},
		PageSize: 10,
		CurPage:  1,
		//	Data:     data,
	}
	context.PageComplex(db, page)
	t.Logf("%v", page)

}

func TestGetSet(t *testing.T) {
	context := GetBizContext()
	context.Set("wing", 1)
	obj := context.Get("wing")
	obj2 := context.Get("nothing")
	assert.Equal(t, obj, 1)
	assert.Equal(t, obj2, nil)
}
