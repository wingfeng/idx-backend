package test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/backend/utils"
)

func TestBizContext_PageComplex(t *testing.T) {
	context := GetContext()
	db := context.DB().Table("Users as user").Select("user.*,org.display_name").Joins("join OrganizationUnit as org on user.ouid=org.id").Where("user.UserName='admin'")
	//	data := make(map[string]interface{}, 0)
	page := &utils.Page{
		Filters:  "",
		PageSize: 10,
		CurPage:  1,
		//	Data:     data,
	}
	context.PageComplex(db, page)
	t.Logf("%v", page)

}

func TestGetSet(t *testing.T) {
	context := GetContext()
	context.Set("wing", 1)
	obj := context.Get("wing")
	obj2 := context.Get("nothing")
	assert.Equal(t, obj, 1)
	assert.Equal(t, obj2, nil)
}
