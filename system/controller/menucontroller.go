package controller

import (
	"log/slog"
	"strings"

	"github.com/wingfeng/backend/utils"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/system/models"
)

type MenuController struct {
	utils.BaseController
}

func (ctrl *MenuController) RegisterRouters(v *gin.RouterGroup) {
	v.PUT(".", ctrl.save)
	v.GET("/navigate", ctrl.PermissionTree)
	v.GET("/tree", ctrl.Tree)
	v.GET("/page", ctrl.Page)
	v.DELETE("/delete", ctrl.Delete)
	v.PUT("/update", ctrl.Update)
}

func (ctrl *MenuController) save(c *gin.Context) {
	row := &models.MenuItem{}
	ctrl.BaseController.Save(row, c)
}

// PermissionTree 获取带权限的菜单信息
// @id 菜单ID
func (ctrl *MenuController) PermissionTree(c *gin.Context) {
	biz := ctrl.BaseController.Prepare(c)
	items := make([]models.MenuItem, 0)

	biz.DB().Model(&models.MenuItem{}).Order("sort_order").Find(&items)
	items = filterMenu(items, biz.User, *biz.Enforcer)
	menus := utils.BuildMenuTree(items)
	c.JSON(200, utils.SysResult{
		Code: 200,
		Msg:  "",
		Data: menus,
	})
}
func (ctrl *MenuController) Page(c *gin.Context) {

	rows := make([]models.MenuItem, 0)
	ctrl.BaseController.Page(&rows, c)
}

func (ctrl *MenuController) Tree(c *gin.Context) {
	biz := ctrl.Prepare(c)
	items := make([]models.MenuItem, 0)
	biz.DB().Order("sort_order").Find(&items)
	menus := utils.BuildMenuTree(items)
	c.JSON(200, utils.SysResult{
		Code: 200,
		Msg:  "",
		Data: menus,
	})
}
func filterMenu(items []models.MenuItem, user string, enforcer casbin.Enforcer) []models.MenuItem {
	result := make([]models.MenuItem, 0)
	for _, item := range items {
		op := item.Operations
		if strings.EqualFold(op, "") {
			op = "*"
		}
		r, err := enforcer.Enforce(user, item.Code, op)
		if err != nil {
			slog.Error("判断权限错误,Err:%v ,user:%s,code:%s,operation:%s", err, user, item.Code, op)
		}
		if r {
			item.RoleOperations = "{\"toolbar\":[\"新建\",\"批量删除\"],\"table\":[\"编辑\",\"删除\"]}"
			result = append(result, item)
		}
	}
	return result
}

// func buildMenuTree(items []models.MenuItem) []interface{} {
// 	var roots []models.MenuItem
// 	linq.From(items).WhereT(func(i models.MenuItem) bool {
// 		return i.Parent == 0
// 	}).ToSlice(&roots)

// 	result := make([]interface{}, 0)

// 	for _, item := range roots {
// 		children := getChildren(item.ID, items)
// 		item.SetChildren(children)
// 		result = append(result, item)
// 	}
// 	return result
// }
// func getChildren(id int64, items []models.MenuItem) []interface{} {
// 	var children []models.MenuItem
// 	linq.From(items).WhereT(func(i models.MenuItem) bool {
// 		return i.Parent == id
// 	}).ToSlice(&children)
// 	result := make([]interface{}, 0)
// 	for _, c := range children {
// 		children := getChildren(c.ID, items)
// 		c.SetChildren(children)
// 		result = append(result, c)
// 	}
// 	return result
// }

// 删除菜单
func (ctrl *MenuController) Delete(c *gin.Context) {
	row := &models.MenuItem{}
	ctrl.BaseController.Delete(row, c)
}

// 修改菜单
func (ctrl *MenuController) Update(c *gin.Context) {
	row := &models.MenuItem{}
	ctrl.BaseController.Update(row, c)
}
