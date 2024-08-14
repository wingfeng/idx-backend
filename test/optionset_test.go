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
	optionset_path = "/api/v1/system/optionset/"
)

func TestGetOptionSet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", optionset_path+"get?id=10863943-e7cf-41a9-acf2-8bb33ab39960", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveOptionSet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.OptionSet{
		ID:    "10863943-e7cf-41a9-acf2-8bb33ab39960",
		Name:  "是否",
		Code:  "yes_or_no",
		Value: "[{\"value\":1,\"text\":\"是\"},{\"value\":0,\"text\":\"否\"}]",
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", optionset_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageOptionSet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", optionset_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteOptionSet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", optionset_path+"del?id=10863943-e7cf-41a9-acf2-8bb33ab39960", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
