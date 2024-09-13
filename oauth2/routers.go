package oauth2

import (
	"github.com/wingfeng/idxadmin/base"
	"github.com/wingfeng/idxadmin/oauth2/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(api *gin.RouterGroup) {

	ctrls := make(map[string]base.Controller)
	ctrl := &controller.ClientController{}
	ctrls["client"] = ctrl
	for c, ctrl := range ctrls {
		g := api.Group(c)
		ctrl.RegisterRouters(g)
	}

}
