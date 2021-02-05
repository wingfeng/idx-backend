package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/sso/models"
	"github.com/wingfeng/backend/utils"
)

type PersistedGrantsController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *PersistedGrantsController) RegisterRouters(r *gin.RouterGroup) {

	//r.PUT("/", ctrl.Save) 只查找，不在这里实现保存供能
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存PersistedGrants 对象
func (ctrl *PersistedGrantsController) Save(ctx *gin.Context) {
	row := &models.PersistedGrants{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除PersistedGrants对象
//@id 对象的ID
func (ctrl *PersistedGrantsController) Delete(ctx *gin.Context) {
	row := &models.PersistedGrants{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *PersistedGrantsController) Page(ctx *gin.Context) {
	data := make([]models.PersistedGrants, 0)

	ctrl.BaseController.Page(&data, ctx)

}

//Get 获取PersistedGrants对象
//@id 对象的id
func (ctrl *PersistedGrantsController) Get(ctx *gin.Context) {
	row := &models.PersistedGrants{}
	ctrl.BaseController.Get(row, ctx)
}
