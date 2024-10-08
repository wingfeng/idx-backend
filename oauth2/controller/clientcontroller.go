package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
)

type ClientController struct {
	base.BaseController
}

// @Description RegisterRouters 注册路由信息
func (ctrl *ClientController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.POST("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

// @Description Save 保存Client 对象
func (ctrl *ClientController) Save(ctx *gin.Context) {
	row := &models.Client{}
	ctrl.BaseController.Save(row, ctx)

}

// @Description Delete 删除Client对象
// @Param id query string true "对象的id"
func (ctrl *ClientController) Delete(ctx *gin.Context) {
	row := &models.Client{}
	ctrl.BaseController.Delete(row, ctx)
}

// @Description Page 取得分页数据
// @Param page query string true "当前页"
// @Param rows query string true "每页的行数"
// @Param filter query string true "查询条件"
// @Param cols query string true "查询出来的数据列"
func (ctrl *ClientController) Page(ctx *gin.Context) {

	rows := make([]models.Client, 0)
	ctrl.BaseController.Page(&rows, ctx)

}

// @Description Get 获取Client对象
// @Param id query string true "对象的id"
func (ctrl *ClientController) Get(ctx *gin.Context) {
	row := &models.Client{}
	ctrl.BaseController.Get(row, ctx)
}
