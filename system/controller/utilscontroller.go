package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
)

type UtilsController struct {
	base.BaseController
}

func (ctrl *UtilsController) RegisterRouters(v1 *gin.RouterGroup) {
	v1.GET("/newid", ctrl.NewId)
}

func (ctrl *UtilsController) NewId(ctx *gin.Context) {
	ctrl.BaseController.GeneratID(ctx)
}
