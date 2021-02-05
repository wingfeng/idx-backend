package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"github.com/wingfeng/backend/system/models"
	"github.com/wingfeng/backend/utils"
)

type UserController struct {
	utils.BaseController
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "KeyForJWT"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Profile       string          `json:"profile"`
	Email         string          `json:"email"`
	EmailVerified bool            `json:"email_verified"`
	Roles         json.RawMessage `json:"role,omitempty"`
	Name          string          `json:"name"`
	OU            string
	OUID          string
	DisplayName   string
	jwt.StandardClaims
}

type NewPassword struct {
	ID           string `json:"id"`
	OidPassword  string `json:"oidpassword"`
	PasswordHash string `json:"passwordhash"`
}

type userLogin struct {
	username string `json:"username"`
	password string `json:"password"`
}

func (ctrl *UserController) RegisterRouters(v1 *gin.RouterGroup) {

	v1.PUT(".", ctrl.Save)
	v1.DELETE("/del", ctrl.Delete)
	v1.GET("/page", ctrl.Page)
	v1.GET("/get", ctrl.Get)
	v1.POST("/changepassword", ctrl.ChangePassword)
	v1.GET("/Plaintext", ctrl.Plaintext)
	v1.PUT("/update", ctrl.Update)
}

func (ctrl *UserController) Save(c *gin.Context) {
	var entity models.User
	err := c.BindJSON(&entity)
	if err != nil {
		log.Errorf("绑定User对象错误!,%v", err.Error())
		c.AbortWithError(500, err)
		return
	}
	password := utils.RandSeq(8)
	entity.PasswordHash = utils.GenHashedPWD(password)
	entity.IsTemporaryPassword = true
	biz := ctrl.Prepare(c)
	err = biz.DB().Save(&entity).Error
	if err != nil {
		c.JSON(500, utils.SysResult{500, "Error", err.Error()})
		return
	}
	c.JSON(200, utils.SysResult{200, "Success", password})
}
func (ctrl *UserController) Delete(ctx *gin.Context) {
	u := &models.User{}
	ctrl.BaseController.Delete(u, ctx)
}
func (ctrl *UserController) Page(c *gin.Context) {
	data := make([]models.User, 0)
	ctrl.BaseController.Page(&data, c)
}
func (ctrl *UserController) Get(ctx *gin.Context) {
	u := &models.User{}
	ctrl.BaseController.Get(u, ctx)
}

func (ctrl *UserController) ChangePassword(ctx *gin.Context) {
	var u NewPassword
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(200, utils.SysResult{200, "绑定User对象错误!", nil})
		return
	}
	biz := ctrl.Prepare(ctx)
	//获取要修改密码的用户信息
	user := &models.User{}
	err = biz.DB().Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		ctx.JSON(500, utils.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
		return
	}
	//判断密码是否一致
	decodedHashedPassword, _ := base64.StdEncoding.DecodeString(user.PasswordHash)
	if r, _ := utils.VerifyHashedPasswordV3(decodedHashedPassword, u.OidPassword); r {
		err = biz.DB().Model(&user).Updates(map[string]interface{}{"PasswordHash": utils.GenHashedPWD(u.PasswordHash), "IsTemporaryPassword": false}).Error
		if err != nil {
			ctx.JSON(500, utils.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
			return
		}
	} else {
		ctx.JSON(500, utils.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
		return
	}
	ctx.JSON(200, utils.SysResult{200, "", nil})
}
func (ctrl *UserController) Plaintext(c *gin.Context) {
	//u := &models.User{}
	//biz := ctrl.Prepare(c)
	//id := c.Query("id")
	//err := biz.DB().Where("id=?", id).First(&u).Error
	//if err != nil {
	//	c.JSON(500, utils.SysResult{500, err.Error(), err})
	//	return
	//}
	//
	//
	//c.JSON(200, utils.SysResult{200, "", row})

}

//修改用户信息
func (ctrl *UserController) Update(c *gin.Context) {
	row := &models.User{}
	ctrl.BaseController.Update(row, c)
}
func (ctrl *UserController) Login(ctx *gin.Context) {
	var loginParam userLogin
	err := ctx.BindJSON(&loginParam)
	if err != nil {
		ctx.JSON(200, utils.SysResult{200, "绑定User对象错误!", nil})
		return
	}
	biz := ctrl.Prepare(ctx)
	//获取要修改密码的用户信息
	user := &models.User{}
	err = biz.DB().Where("UserName = ?", loginParam.username).First(&user).Error
	if err != nil {
		ctx.JSON(500, utils.SysResult{500, fmt.Sprintf("用户登录失败!"), loginParam.username})
		return
	}
	//判断密码是否一致
	decodedHashedPassword, _ := base64.StdEncoding.DecodeString(user.PasswordHash)
	if r, _ := utils.VerifyHashedPasswordV3(decodedHashedPassword, loginParam.password); r {
		claims := &CustomClaims{}
		claims.Subject = user.ID
		claims.DisplayName = user.DisplayName
		claims.Email = user.Email
		claims.EmailVerified = user.EmailConfirmed
		claims.OU = user.OU
		claims.OUID = user.OUID

		jwt := NewJWT()
		token, err := jwt.CreateToken(*claims)
		if err != nil {
			panic(err)
		}
		_ = token

	} else {
		ctx.JSON(500, utils.SysResult{500, fmt.Sprintf("用户登录失败!"), loginParam.username})
		return
	}
	ctx.JSON(200, utils.SysResult{200, "", nil})
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
