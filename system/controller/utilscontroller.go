package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idxadmin/base"
)

type UtilsController struct {
	base.BaseController
}

func (ctrl *UtilsController) RegisterRouters(v1 *gin.RouterGroup) {
	v1.GET("/get", ctrl.Get)
}

func (ctrl *UtilsController) Get(ctx *gin.Context) {
	ctrl.BaseController.GeneratID(ctx)
}
