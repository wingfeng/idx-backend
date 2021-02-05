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
	clients_path = "/api/v1/system/clients/"
)

func TestGetClients(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", clients_path+"get?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveClients(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.Client{
		ID:                               2,
		ClientID:                         "local_test",
		Enabled:                          true,
		ProtocolType:                     "oidc",
		RequireClientSecret:              true,
		ClientName:                       "Code Client",
		RequireConsent:                   true,
		AllowRememberConsent:             true,
		AlwaysIncludeUserClaimsInIDToken: false,
		AllowAccessTokensViaBrowser:      true,
		BackChannelLogoutSessionRequired: true,
		IDentityTokenLifetime:            300,
		AccessTokenLifetime:              3600,
		AbsoluteRefreshTokenLifetime:     2592000,
		SlidingRefreshTokenLifetime:      2592000,
		AuthorizationCodeLifetime:        300,
		RefreshTokenUsage:                1,
		RefreshTokenExpiration:           1,
		ClientClaimsPrefix:               "client_",
		DeviceCodeLifetime:               300,

		EnableLocalLogin: true,
	}
	json, _ := json.Marshal(u)
	strJson := string(json)
	_ = strJson
	req, _ := http.NewRequest("PUT", clients_path, bytes.NewReader(json))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}

func TestPageClients(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", clients_path+"page?page=1&rows=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteClients(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", clients_path+"del?id=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
