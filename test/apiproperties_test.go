package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/wingfeng/idx/models"
)

const (
	apiproperties_path = "/api/v1/system/apiproperties/"
)

func TestGetAPIProperties(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiproperties_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveAPIProperties(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.APIProperties{
		ID: 1,
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", apiproperties_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageAPIProperties(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", apiproperties_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteAPIProperties(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", apiproperties_path+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
