package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
)

type RoleController struct {
	base.BaseController
}

// RegisterRouters 注册路由信息
func (ctrl *RoleController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
	r.PUT("/update", ctrl.Update)
}

// Save 保存Role 对象
func (ctrl *RoleController) Save(ctx *gin.Context) {
	row := &models.Role{}
	ctrl.BaseController.Save(row, ctx)

}

// Delete 删除Role对象
// @id 对象的ID
func (ctrl *RoleController) Delete(ctx *gin.Context) {
	row := &models.Role{}
	ctrl.BaseController.Delete(row, ctx)
}

// Page 取得分页数据
// @page 当前页
// @rows 每页的行数
// @filter 查询条件
// @cols 查询出来的数据列
func (ctrl *RoleController) Page(ctx *gin.Context) {
	data := make([]models.Role, 0)

	ctrl.BaseController.Page(&data, ctx)

}

// Get 获取Role对象
// @id 对象的id
func (ctrl *RoleController) Get(ctx *gin.Context) {
	row := &models.Role{}
	ctrl.BaseController.Get(row, ctx)
}

// 修改角色信息
func (ctrl *RoleController) Update(c *gin.Context) {
	row := &models.Role{}
	ctrl.BaseController.Update(row, c)
}
