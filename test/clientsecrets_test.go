package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idx/utils"

	"github.com/magiconair/properties/assert"
)

const (
	clientsecrets_path = "/api/v1/system/clientsecrets/"
)

func TestGetClientSecrets(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", clientsecrets_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveClientSecrets(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	sc := &models.ClientSecrets{
		Type:     "SHA256",
		ClientID: 2,
	}
	sc.Value = utils.HashString("local_secret")
	sc.Expiration = time.Now().AddDate(1, 0, 0)

	json, _ := json.Marshal(sc)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", clientsecrets_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageClientSecrets(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", clientsecrets_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteClientSecrets(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", clientsecrets_path+"del?id=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
