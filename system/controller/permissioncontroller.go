package controller

import (
	"log/slog"
	"net/http"

	"github.com/wingfeng/idxadmin/base"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type PermissionParam struct {
	User   string   `json:"user"`
	Codes  []string `json:"code"`
	Action string   `json:action`
}

type PermissionController struct {
	base.BaseController
}

func (ctrl *PermissionController) RegisterRouters(r *gin.RouterGroup) {
	r.GET("/check", ctrl.CheckPermission)
	r.POST("/setmenupermission", ctrl.UpdateMenuPermission)
	r.GET("/getuserpermissions", ctrl.GetUserPermission)
}

func (ctrl *PermissionController) CheckPermission(c *gin.Context) {

}

func (ctrl *PermissionController) UpdateMenuPermission(context *gin.Context) {
	var param PermissionParam
	//	var param model.ModifyRole
	e, _ := context.Get(base.Const_CasbinKey)
	enf := e.(*casbin.Enforcer)
	if err := context.ShouldBindJSON(&param); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"messages": http.StatusText(http.StatusBadRequest),
			"errInfo":  err.Error(),
		})
		return
	}
	//biz := ctrl.prepare(context)

	//删除用户所有权限
	_, err := enf.DeletePermissionsForUser(param.User)
	if err != nil {
		slog.Error(err.Error())
	}
	//添加用户所有
	for _, role := range param.Codes {
		_, err = enf.AddPolicy(param.User, role, param.Action)
		if err != nil {
			slog.Error(err.Error())
		}
	}
	if err == nil {
		context.JSON(http.StatusOK, base.SysResult{200, "Success", nil})
	} else {
		context.JSON(http.StatusInternalServerError, base.SysResult{500, "Error", err})
	}

}

// GetUserPermission 获取用户（或角色）具备的权限
// @account 用户账号或者角色名称
func (ctrl *PermissionController) GetUserPermission(c *gin.Context) {
	e, _ := c.Get(base.Const_CasbinKey)
	enf := e.(*casbin.Enforcer)
	user := c.Query("account")
	data := enf.GetPermissionsForUser(user)
	c.JSON(http.StatusOK, base.SysResult{
		Code: 200,
		Msg:  "Success",
		Data: data,
	})
}
