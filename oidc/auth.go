package oidc

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/patrickmn/go-cache"
	"github.com/wingfeng/backend/utils"
	"go.uber.org/zap"
)

var (
	publicKey    *rsa.PublicKey
	userEndpoint string
	memCache     *cache.Cache
	mutex        sync.Mutex
)

type UserInfo struct {
	Subject       string          `json:"sub"`
	Profile       string          `json:"profile"`
	Email         string          `json:"email"`
	EmailVerified bool            `json:"email_verified"`
	Roles         json.RawMessage `json:"role,omitempty"`
	PreferedName  json.RawMessage `json:"preferred_username,omitempty"`
	Name          string          `json:"name"`
	OU            string
	OUID          string
	DisplayName   string
	//	claims        map[string]interface{}
}

// Init 初始化OIDC相关的参数
func Init(issuer string, userInfoEndpoint string) {
	//从Issuer中拿到PublicKey
	userEndpoint = userInfoEndpoint
	slog.Info("Auth Jwt module starting up...")
	set, err := jwk.Fetch(issuer)
	if err != nil {
		slog.Error("获取JWKS 失败!", zap.Error(err))
	}
	key, _ := set.Keys[0].Materialize()
	publicKey = key.(*rsa.PublicKey)
	memCache = cache.New(30*time.Second, 10*time.Second)
}

// AuthWare 通过OpenID的JWT验证客户身份
func AuthWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		token, err := verifyHeader(authHeader)

		if err != nil {
			slog.Error("解析token失败", zap.Error(err))
			c.AbortWithStatusJSON(401, utils.SysResult{Code: 401, Msg: "解析token失败", Data: err.Error()})
			c.Next()
			return
		}
		claims, err := extractClaims(token)
		if err != nil {
			slog.Error("解析token失败", zap.Error(err))
			c.AbortWithStatusJSON(401, utils.SysResult{Code: 401, Msg: "解析token失败", Data: err.Error()})

			return
		}

		mutex.Lock()
		u, cacheExist := memCache.Get(token)
		var user *UserInfo
		if !cacheExist {
			slog.Warn("User Cache不存在,重新写入!")
			user, err = retrieveUserInfo(token)
			dt := claims["exp"].(float64)
			intExp := int64(dt)
			expTime := time.Unix(intExp, 0)
			duration := expTime.Sub(time.Now())
			memCache.Set(token, user, duration)
			if err != nil {
				slog.Error("获取用户信息失败,Err:%v ", err)
				c.AbortWithStatusJSON(401, utils.SysResult{Code: 401, Msg: "获取用户信息失败", Data: err.Error()})
				return
			}
			//将rbac现有用户的角色清除
			e, enableCasbin := c.Get(utils.Const_CasbinKey)
			if enableCasbin {
				ef := e.(*casbin.Enforcer)
				ef.DeleteRolesForUser(user.Name)
				rm := ef.GetRoleManager()

				if user.Roles != nil {
					s := string(user.Roles)
					if strings.Index(s, "[") == 0 {

						s = strings.TrimPrefix(s, "[")
						s = strings.TrimSuffix(s, "]")
						roles := strings.Split(s, ",")
						for _, role := range roles {
							role = strings.Trim(role, "\"")
							slog.Info("Role %s added", role)
							rm.AddLink(user.Name, role)

						}
					} else {
						s = strings.Trim(s, "\"")
						slog.Info("Role %s added", s)
						rm.AddLink(user.Name, s)

					}
					slog.Info("User Roles:", s)

				}
			}
			slog.Info("User Claim:%v", user)
		} else {
			user = u.(*UserInfo)
		}

		mutex.Unlock()
		if user == nil {
			slog.Error("获取用户信息失败,Err:%v ", err)
			c.AbortWithStatusJSON(401, utils.SysResult{Code: 401, Msg: "获取用户信息失败", Data: err.Error()})
			return
		}
		//设置当前用户信息
		c.Set(utils.Const_UserIDKey, user.Subject)
		// c.Set("preferred_username", claims["preferred_username"])
		c.Set(utils.Const_UserNameKey, user.DisplayName)
		c.Set(utils.Const_OUIDKey, user.OUID)
		c.Set(utils.Const_OUKey, user.OU)
		c.Next()
	}
}
func retrieveUserInfo(token string) (*UserInfo, error) {
	req, err := http.NewRequest("GET", userEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("oidc: create GET request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	client := http.DefaultClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("oidc: failed to decode userinfo: %v ,boyd:%s", err, string(body))
	}
	// sU := string(body)
	// log.Infof("User Json:%s", sU)
	// userInfo.claims = body
	return &userInfo, nil
}
func extractClaims(tokenStr string) (jwt.MapClaims, error) {

	// 基于公钥验证Token合法性
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 基于JWT的第一部分中的alg字段值进行一次验证
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("验证Token的加密类型错误")
		}

		return publicKey, nil
	})

	if err != nil {
		slog.Error("解析JWT失败!", zap.Error(err))
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil

	} else {
		fmt.Println(err)
	}
	return nil, nil
}
func verifyHeader(header string) (string, error) {
	if strings.EqualFold(header, "") {
		return "", fmt.Errorf("Header为空")
	}
	s := strings.Split(header, " ")
	if len(s) != 2 {
		return "", fmt.Errorf("验证Header失败! Header:%s", header)
	} else if !strings.EqualFold(s[0], "Bearer") {
		return "", errors.New("验证方式错误,请使用Bearer验证")
	}
	return s[1], nil

}
