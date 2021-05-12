package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/utils"
	"github.com/wingfeng/idx/models"
)

type APIResourcesController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *APIResourcesController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT("/", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存APIResources 对象
func (ctrl *APIResourcesController) Save(ctx *gin.Context) {
	row := &models.APIResources{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除APIResources对象
//@id 对象的ID
func (ctrl *APIResourcesController) Delete(ctx *gin.Context) {
	row := &models.APIResources{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *APIResourcesController) Page(ctx *gin.Context) {
	data := make([]models.APIResources, 0)

	ctrl.BaseController.Page(&data, ctx)

}

//Get 获取APIResources对象
//@id 对象的id
func (ctrl *APIResourcesController) Get(ctx *gin.Context) {
	row := &models.APIResources{}
	ctrl.BaseController.Get(row, ctx)
}
