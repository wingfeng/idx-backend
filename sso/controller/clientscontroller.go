package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/utils"
	idxmodels "github.com/wingfeng/idx/models"
)

type ClientsController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *ClientsController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存Clients 对象
func (ctrl *ClientsController) Save(ctx *gin.Context) {
	row := &idxmodels.Client{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除Clients对象
//@id 对象的ID
func (ctrl *ClientsController) Delete(ctx *gin.Context) {
	row := &idxmodels.Client{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *ClientsController) Page(ctx *gin.Context) {
	data := make([]idxmodels.Client, 0)

	ctrl.BaseController.Page(&data, ctx)

}

//Get 获取Clients对象
//@id 对象的id
func (ctrl *ClientsController) Get(ctx *gin.Context) {
	row := &idxmodels.Client{}
	ctrl.BaseController.Get(row, ctx)
}
