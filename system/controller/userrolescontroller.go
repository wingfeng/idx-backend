package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
)

type UserRolesController struct {
	base.BaseController
}

// RegisterRouters 注册路由信息
func (ctrl *UserRolesController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

// Save 保存UserRoles 对象
func (ctrl *UserRolesController) Save(ctx *gin.Context) {
	row := &models.UserRoles{}
	ctrl.BaseController.Save(row, ctx)

}

// Delete 删除UserRoles对象
// @id 对象的ID
func (ctrl *UserRolesController) Delete(ctx *gin.Context) {
	row := &models.UserRoles{}
	ctrl.BaseController.Delete(row, ctx)
}

// Page 取得分页数据
// @page 当前页
// @rows 每页的行数
// @filter 查询条件
// @cols 查询出来的数据列
func (ctrl *UserRolesController) Page(ctx *gin.Context) {
	data := make([]models.UserRoles, 0)

	ctrl.BaseController.Page(&data, ctx)

}

// Get 获取UserRoles对象
// @id 对象的id
func (ctrl *UserRolesController) Get(ctx *gin.Context) {
	row := &models.UserRoles{}
	ctrl.BaseController.Get(row, ctx)
}
