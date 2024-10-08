package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/utils"
	"gorm.io/gorm"
)

type OrgUnitController struct {
	base.BaseController
}

func (ctrl *OrgUnitController) RegisterRouters(v1 *gin.RouterGroup) {

	v1.PUT(".", ctrl.Save)
	v1.DELETE("/del", ctrl.Delete)
	v1.GET("/page", ctrl.Page)
	v1.GET("/get", ctrl.Get)
	v1.GET("/tree", ctrl.Tree)
	v1.GET("/getchildren", ctrl.getChildren)
	v1.PUT("/update", ctrl.Update)
}

func (ctrl *OrgUnitController) Save(ctx *gin.Context) {
	row := &models.OrganizationUnit{}
	ctrl.BaseController.Save(row, ctx)

}
func (ctrl *OrgUnitController) Delete(ctx *gin.Context) {
	row := &models.OrganizationUnit{}
	ctrl.BaseController.Delete(row, ctx)
}
func (ctrl *OrgUnitController) Page(c *gin.Context) {
	data := make([]models.OrganizationUnit, 0)

	ctrl.BaseController.Page(&data, c)

}

func (ctrl *OrgUnitController) Get(c *gin.Context) {
	row := &models.OrganizationUnit{}
	ctrl.BaseController.Get(row, c)
}

func (ctrl *OrgUnitController) Tree(c *gin.Context) {
	biz := ctrl.Prepare(c)
	items := make([]models.OrganizationUnit, 0)

	err := biz.DB().WithContext(c.Request.Context()).Order("sort_order").Find(&items).Error
	input, _ := utils.ConvertToTreeSlice(items)
	ous := utils.BuildTree(input)

	if err != nil {
		c.AbortWithStatusJSON(500, base.SysResult{
			Code: 500,
			Msg:  "构建Org Tree失败,Error",
			Data: err,
		})
		return
	}

	c.JSON(200, base.SysResult{
		Code: 200,
		Msg:  "",
		Data: ous,
	})
}

func (ctrl *OrgUnitController) getChildren(c *gin.Context) {
	parent := ""
	if c.Query("parent") != "" {
		parent = c.Query("parent")
	}
	items := make([]models.OrganizationUnit, 0)
	biz := ctrl.Prepare(c)
	var tx *gorm.DB
	if strings.EqualFold(parent, "") {
		tx = biz.DB().Model(items).Where("parent is null").Find(&items)
	} else {
		tx = biz.DB().Model(items).Where("parent = ?", parent).Find(&items)
	}

	if tx.Error != nil {
		c.AbortWithStatusJSON(500, base.SysResult{
			Code: 500,
			Msg:  fmt.Sprintf("Get Children Failed,Id %s", parent),
			Data: tx.Error,
		})
		return
	}
	c.JSON(200, base.SysResult{
		Code: 200,
		Msg:  "",
		Data: items,
	})
}

// 修改组织信息
func (ctrl *OrgUnitController) Update(c *gin.Context) {
	row := &models.OrganizationUnit{}
	ctrl.BaseController.Update(row, c)
}
