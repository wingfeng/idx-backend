package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	orgunit_path = "/api/v1/system/orgunit/"
)

func TestGetOUTree(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", orgunit_path+"tree?parent=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestGetChildren(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", orgunit_path+"getchildren?parent=", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
