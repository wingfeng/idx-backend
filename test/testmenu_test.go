package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wingfeng/idxadmin/system/models"

	"github.com/magiconair/properties/assert"
)

const (
	menu_path = "/api/v1/system/menu/"
)

func TestMenu(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/system/sidebar", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyStr := w.Body.String()
	t.Logf("Body:%v", bodyStr)
	assert.Equal(t, len(bodyStr) > 0, true)
}

func TestSaveMenu(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.MenuItem{
		ID:        "1",
		Name:      "admin",
		Icon:      "skhfalslloisjallialisf",
		Component: "koaniousa",
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", menu_path+"", bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
