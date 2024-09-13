package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wingfeng/idx/models"

	"github.com/magiconair/properties/assert"
)

const (
	userroles_path = "/api/v1/system/userroles/"
)

func TestGetUserRoles(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", userroles_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveUserRoles(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	u := &models.UserRoles{
		UserId: "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e",
		RoleId: "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b",
		OUId:   "1328680589330485248",
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", userroles_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageUserRoles(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", userroles_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteUserRoles(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", userroles_path+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
