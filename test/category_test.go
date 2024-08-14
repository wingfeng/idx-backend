package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idxadmin/system/models"
)

const (
	category_path = "/api/v1/system/category/"
)

func TestSaveCategory(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.Category{
		ID: "1",
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", category_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestGetCategory(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", category_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageCategory(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", category_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteCategory(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", category_path+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
