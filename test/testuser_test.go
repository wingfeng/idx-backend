package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/wingfeng/idx/models"
	"gopkg.in/guregu/null.v4"

	"github.com/magiconair/properties/assert"
)

const (
	userpath = "/api/v1/system/user/"
)

func TestGetUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/system/user/get?id=7a45cb54-b0ff-4ecd-95b9-074d33aaac1e", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestSaveUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	u := &models.User{
		Id:             "eaeff30b-2d57-4732-b5a2-d3f6f0a5a710",
		Account:        "admin",
		Email:          "wing@fire.loc",
		OUId:           "1328680589330485248",
		OU:             "XX软件有限公司",
		EmailConfirmed: true,
		LockoutEnd:     null.TimeFrom(time.Now()),
	}
	jsonbyte, _ := json.Marshal(u)

	tmp := make(map[string]interface{})
	err := json.Unmarshal(jsonbyte, &tmp)
	if err != nil {
		panic(err)
	}
	tmp["lockoutend"] = "2020-11-25T22:20:00+08:00" //test local time
	jsonbyte, _ = json.Marshal(tmp)
	t.Logf("user json %s:", string(jsonbyte))
	req, _ := http.NewRequest("PUT", "/api/v1/system/user", bytes.NewReader(jsonbyte))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
func TestDeleteUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", userpath+"del?id=5fae682fffd13d74352be0a9", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(bodyString) > 0, true)
}
func TestPageUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	filters := "UserName like ? and id=?"
	strFilters := url.QueryEscape(filters)
	args := "%admin% 7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	uriArgs := url.QueryEscape(args)
	cols := url.QueryEscape("\"UserName\":true,\"ID\":true")
	req, _ := http.NewRequest("GET", userpath+"page?page=1&size=10&filters="+strFilters+"&args="+uriArgs+"&cols={"+cols+"}&ordertype=asc&ordername=username", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bodyString := w.Body.String()
	t.Logf("Body:%v", bodyString)
	assert.Equal(t, len(w.Body.String()) > 0, true)
}
