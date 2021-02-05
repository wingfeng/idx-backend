package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/sso/models"
	"github.com/wingfeng/backend/utils"
)

type APIPropertiesController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *APIPropertiesController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT("/", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存APIProperties 对象
func (ctrl *APIPropertiesController) Save(ctx *gin.Context) {
	row := &models.APIProperties{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除APIProperties对象
//@id 对象的ID
func (ctrl *APIPropertiesController) Delete(ctx *gin.Context) {
	row := &models.APIProperties{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *APIPropertiesController) Page(ctx *gin.Context) {
	data := make([]models.APIProperties, 0)

	ctrl.BaseController.Page(&data, ctx)

}

//Get 获取APIProperties对象
//@id 对象的id
func (ctrl *APIPropertiesController) Get(ctx *gin.Context) {
	row := &models.APIProperties{}
	ctrl.BaseController.Get(row, ctx)
}
