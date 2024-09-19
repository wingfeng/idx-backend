package test

import (
	"github.com/wingfeng/idx/models"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	client_path = "/api/v1/system/client"
)

func TestSaveClient(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	u := &models.Client{
		Id: 1,
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", client_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestGetClient(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", client_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageClient(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", client_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteClient(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", client_path+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
