package test

import (
	"encoding/base64"
	"testing"

	"github.com/wingfeng/backend/utils"

	"github.com/magiconair/properties/assert"
)

//为了兼容aspnet core的password hash
func TestPWDCompatibility(t *testing.T) {
	hashedPwd := utils.GenHashedPWD("fire@123")
	t.Logf("hashed password:%v", hashedPwd)
	decodedHashedPassword, _ := base64.StdEncoding.DecodeString(hashedPwd)
	t.Logf("decode:%v", decodedHashedPassword)
	r, _ := utils.VerifyHashedPasswordV3(decodedHashedPassword, "fire@123")

	assert.Equal(t, r, true)
}
