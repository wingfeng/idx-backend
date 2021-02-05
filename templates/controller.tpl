package controller

import (
	"github.com/wingfeng/backend/system/models"
    "github.com/wingfeng/backend/utils"
	"github.com/gin-gonic/gin"
)

type {{.ShortName}}Controller struct{
    utils.BaseController
}

//RegisterRouters 注册路由信息
func (ctrl *{{.ShortName}}Controller) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)	
}
//Save 保存{{.ShortName}} 对象
func (ctrl *{{.ShortName}}Controller) Save(ctx *gin.Context) {
	 row:=&{{.Type}}{}
	ctrl.BaseController.Save(row, ctx)

}   
//Delete 删除{{.ShortName}}对象
//@id 对象的ID
func (ctrl *{{.ShortName}}Controller) Delete(ctx *gin.Context) {
	row :=&{{.Type}}{}
	ctrl.BaseController.Delete(row, ctx)
}
//Page 取得分页数据
//@page 当前页
//@rows 每页的行数
//@filter 查询条件
//@cols 查询出来的数据列
func (ctrl *{{.ShortName}}Controller) Page(ctx *gin.Context) {

	rows := make([]{{.Type}}, 0)
	ctrl.BaseController.Page(&rows, ctx)

}
//Get 获取{{.ShortName}}对象
//@id 对象的id
func(ctrl *{{.ShortName}}Controller) Get(ctx *gin.Context){
    row:=&{{.Type}}{}
    ctrl.BaseController.Get(row,ctx)
}



