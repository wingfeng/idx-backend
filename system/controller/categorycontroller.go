package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/system/models"
)

type CategoryController struct {
	base.BaseController
}

// RegisterRouters 注册路由信息
func (ctrl *CategoryController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT("/", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

func (ctrl *CategoryController) Save(ctx *gin.Context) {
	row := &models.Category{}
	ctrl.BaseController.Save(row, ctx)

}

// Delete 删除Category对象
//
//	@id	对象的ID
func (ctrl *CategoryController) Delete(ctx *gin.Context) {
	row := &models.Category{}
	ctrl.BaseController.Delete(row, ctx)
}

// Page 取得分页数据
//
//	@page	当前页
//	@rows	每页的行数
//	@filter	查询条件
//	@cols	查询出来的数据列
func (ctrl *CategoryController) Page(ctx *gin.Context) {

	rows := make([]models.Category, 0)
	ctrl.BaseController.Page(&rows, ctx)

}

// Get 获取Category对象
//
//	@id	对象的id
func (ctrl *CategoryController) Get(ctx *gin.Context) {
	row := &models.Category{}
	ctrl.BaseController.Get(row, ctx)
}
