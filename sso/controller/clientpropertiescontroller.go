package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/sso/models"
	"github.com/wingfeng/backend/utils"
)

type ClientPropertiesController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *ClientPropertiesController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT("/", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存ClientProperties 对象
func (ctrl *ClientPropertiesController) Save(ctx *gin.Context) {
	row := &models.ClientProperties{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除ClientProperties对象
//@id 对象的ID
func (ctrl *ClientPropertiesController) Delete(ctx *gin.Context) {
	row := &models.ClientProperties{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *ClientPropertiesController) Page(ctx *gin.Context) {
	data := make([]models.ClientProperties, 0)

	ctrl.BaseController.Page(&data, ctx)

}

//Get 获取ClientProperties对象
//@id 对象的id
func (ctrl *ClientPropertiesController) Get(ctx *gin.Context) {
	row := &models.ClientProperties{}
	ctrl.BaseController.Get(row, ctx)
}
