package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/utils"
	"github.com/wingfeng/idx/models"
)

type ClientGrantTypesController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *ClientGrantTypesController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存ClientGrantTypes 对象
func (ctrl *ClientGrantTypesController) Save(ctx *gin.Context) {
	row := &models.ClientGrantTypes{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除ClientGrantTypes对象
//@id 对象的ID
func (ctrl *ClientGrantTypesController) Delete(ctx *gin.Context) {
	row := &models.ClientGrantTypes{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *ClientGrantTypesController) Page(ctx *gin.Context) {

	rows := make([]models.ClientGrantTypes, 0)
	ctrl.BaseController.Page(&rows, ctx)

}

//Get 获取ClientGrantTypes对象
//@id 对象的id
func (ctrl *ClientGrantTypesController) Get(ctx *gin.Context) {
	row := &models.ClientGrantTypes{}
	ctrl.BaseController.Get(row, ctx)
}
