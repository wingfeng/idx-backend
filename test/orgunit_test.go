package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/backend/system/models"
	"github.com/wingfeng/backend/utils"
)

const (
	oupath = "/api/v1/system/orgunit/"
)

func TestGetOrgUnit(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", oupath+"get?id=2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}

func TestDeleteOrgUnit(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", oupath+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
func TestSaveOrgUnit(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", oupath+"get?id=1", nil)
	router.ServeHTTP(w, req)
	sResult := &utils.SysResult{}

	u := &models.OrganizationUnit{
		ID:          "019291210210",
		Name:        "广州XX软件",
		DisplayName: "广州XX软件",
	}
	sResult.Data = u
	err := json.Unmarshal(w.Body.Bytes(), sResult)

	// u := &models.OrganizationUnit{
	// 	ID:   1,
	// 	Name: "广州千秋软件",
	// }
	u = sResult.Data.(*models.OrganizationUnit)
	u.Name = "广州千秋软件"
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ = http.NewRequest("PUT", oupath, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, err, nil)
	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
func TestPageOrgUnit(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", oupath+"page?page=1&rows=20", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}

func TestTreeOrgUnit(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", oupath+"tree", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
