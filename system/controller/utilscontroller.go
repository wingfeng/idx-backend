package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/utils"
)

type UtilsController struct {
	utils.BaseController
}

func (ctrl *UtilsController) RegisterRouters(v1 *gin.RouterGroup) {
	v1.GET("/get", ctrl.Get)
}

func (ctrl *UtilsController) Get(ctx *gin.Context) {
	ctrl.BaseController.GeneratID(ctx)
}
