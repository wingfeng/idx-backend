package routers

import (
	sso "github.com/wingfeng/backend/sso/controller"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/backend/system/controller"
	"github.com/wingfeng/backend/utils"
)

//RegisterRouter 注册路由
func RegisterRouter(api *gin.RouterGroup) {

	ctrls := make(map[string]utils.Controller)
	uc := &controller.UserController{}
	ctrls["user"] = uc
	mc := &controller.MenuController{}
	ctrls["menu"] = mc
	ctrl := &controller.OrgUnitController{}
	ctrls["orgunit"] = ctrl
	clientCtrl := &sso.ClientsController{}
	ctrls["clients"] = clientCtrl
	RoleController := &controller.RoleController{}
	ctrls["role"] = RoleController
	UserRolesController := &controller.UserRolesController{}
	ctrls["userroles"] = UserRolesController
	OptionSetController := &controller.OptionSetController{}
	ctrls["optionset"] = OptionSetController
	UtilsController := &controller.UtilsController{}
	ctrls["utils"] = UtilsController
	// ApipropertiesController := &controller.ApipropertiesController{}
	// ctrls["apiproperties"] = ApipropertiesController
	// ApiresourcesController := &controller.ApiresourcesController{}
	// ctrls["apiresources"] = ApiresourcesController
	// ApisecretsController := &controller.ApisecretsController{}
	// ctrls["apisecrets"] = ApisecretsController
	// ClientpropertiesController := &controller.ClientpropertiesController{}
	// ctrls["clientproperties"] = ClientpropertiesController
	ClientsecretsController := &sso.ClientSecretsController{}
	ctrls["clientsecrets"] = ClientsecretsController
	ClientRedirectURIsController := &sso.ClientRedirectURIsController{}
	ctrls["clientredirecturis"] = ClientRedirectURIsController
	ClientScopesController := &sso.ClientScopesController{}
	ctrls["clientscopes"] = ClientScopesController
	ClientGrantTypesController := &sso.ClientGrantTypesController{}
	ctrls["clientgranttypes"] = ClientGrantTypesController
	// IdentityresourcesController := &controller.IdentityresourcesController{}
	// ctrls["identityresources"] = IdentityresourcesController
	// PersistedgrantsController := &controller.PersistedgrantsController{}
	//	ctrls["persistedgrants"] = PersistedgrantsController
	PermissionController := &controller.PermissionController{}
	ctrls["permission"] = PermissionController
	for c, ctrl := range ctrls {
		g := api.Group(c)
		ctrl.RegisterRouters(g)
	}

}
