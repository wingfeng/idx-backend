package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wingfeng/idxadmin/system/controller"

	"github.com/magiconair/properties/assert"
)

func TestMenuPermission(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	model := &controller.PermissionParam{
		User:   "fenggr",
		Codes:  []string{"system_mng", "system_setting", "system_menu"},
		Action: "*",
	}
	data, _ := json.Marshal(model)

	req, _ := http.NewRequest("POST", "/api/v1/system/permission/setmenupermission", bytes.NewReader(data))
	router.ServeHTTP(w, req)
	model = &controller.PermissionParam{
		User:   "*",
		Codes:  []string{"home"},
		Action: "*",
	}
	data, _ = json.Marshal(model)

	req, _ = http.NewRequest("POST", "/api/v1/system/permission/setmenupermission", bytes.NewReader(data))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "success")
}
func TestGetUserPermission(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/system/permission/getuserpermissions?account=fenggr", nil)
	router.ServeHTTP(w, req)
	bodyString := w.Body.String()
	t.Logf("Body:%s", bodyString)
	assert.Equal(t, w.Code, 200)
}
