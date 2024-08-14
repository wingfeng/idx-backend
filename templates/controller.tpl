package controller

import (
	"github.com/wingfeng/idxadmin/system/models"
    "github.com/wingfeng/idxadmin/base"
	"github.com/gin-gonic/gin"
)

type {{.ShortName}}Controller struct{
    base.BaseController
}

//@Description RegisterRouters 注册路由信息
func (ctrl *{{.ShortName}}Controller) RegisterRouters(r *gin.RouterGroup) {

	r.PUT(".", ctrl.Save)
	r.DELETE("/del", ctrl.Delete)
	r.GET("/page", ctrl.Page)
	r.GET("/get", ctrl.Get)	
}
//@Description Save 保存{{.ShortName}} 对象
func (ctrl *{{.ShortName}}Controller) Save(ctx *gin.Context) {
	 row:=&{{.Type}}{}
	ctrl.BaseController.Save(row, ctx)

}   
//@Description Delete 删除{{.ShortName}}对象
// @Param id query string true "对象的id"
func (ctrl *{{.ShortName}}Controller) Delete(ctx *gin.Context) {
	row :=&{{.Type}}{}
	ctrl.BaseController.Delete(row, ctx)
}
//@Description Page 取得分页数据
//@Param page query string true "当前页"
//@Param rows query string true "每页的行数"
//@Param filter query string true "查询条件"
//@Param cols query string true "查询出来的数据列"
func (ctrl *{{.ShortName}}Controller) Page(ctx *gin.Context) {

	rows := make([]{{.Type}}, 0)
	ctrl.BaseController.Page(&rows, ctx)

}
//@Description Get 获取{{.ShortName}}对象
// @Param id query string true "对象的id"
func(ctrl *{{.ShortName}}Controller) Get(ctx *gin.Context){
    row:=&{{.Type}}{}
    ctrl.BaseController.Get(row,ctx)
}



