package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/system/models"
	"github.com/wingfeng/backend/utils"
)

type OptionSetController struct {
	utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *OptionSetController) RegisterRouters(r *gin.RouterGroup) {

	r.PUT("/", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)
}

//Save 保存OptionSet 对象
func (ctrl *OptionSetController) Save(ctx *gin.Context) {
	row := &models.OptionSet{}
	ctrl.BaseController.Save(row, ctx)

}

//Delete 删除OptionSet对象
//@id 对象的ID
func (ctrl *OptionSetController) Delete(ctx *gin.Context) {
	row := &models.OptionSet{}
	ctrl.BaseController.Delete(row, ctx)
}

//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *OptionSetController) Page(ctx *gin.Context) {

	rows := make([]models.OptionSet, 0)
	ctrl.BaseController.Page(&rows, ctx)

}

//Get 获取OptionSet对象
//@id 对象的id
func (ctrl *OptionSetController) Get(ctx *gin.Context) {
	row := &models.OptionSet{}
	ctrl.BaseController.Get(row, ctx)
}
