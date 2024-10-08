package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx-oauth2/utils"
	"github.com/wingfeng/idx/models"
	"github.com/wingfeng/idxadmin/base"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	base.BaseController
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
	OldPassword  string `json:"oldpassword"`
	PasswordHash string `json:"passwordhash"`
}

type userLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (ctrl *UserController) RegisterRouters(v1 *gin.RouterGroup) {

	v1.PUT(".", ctrl.Save)
	v1.DELETE("/del", ctrl.Delete)
	v1.POST("/page", ctrl.Page)
	v1.GET("/get", ctrl.Get)
	v1.POST("/changepassword", ctrl.ChangePassword)
	v1.POST("/resetpassword", ctrl.ResetPassword)
	v1.POST("/login", ctrl.Login)
	v1.GET("/Plaintext", ctrl.Plaintext)
	v1.PUT("/update", ctrl.Update)
}

func (ctrl *UserController) Save(c *gin.Context) {

	var entity models.User
	err := c.BindJSON(&entity)
	if err != nil {
		slog.Error("绑定User对象错误!,%v", "error", err.Error())
		c.AbortWithError(500, err)
		return
	}
	biz := ctrl.Prepare(c)

	err = biz.DB().Omit("password_hash").Save(&entity).Error
	if err != nil {
		c.JSON(500, base.SysResult{500, "Error", err.Error()})
		return
	}
	c.JSON(200, base.SysResult{200, "Success", ""})
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
		ctx.JSON(200, base.SysResult{200, "绑定User对象错误!", nil})
		return
	}
	biz := ctrl.Prepare(ctx)
	//获取要修改密码的用户信息
	user := &models.User{}
	err = biz.DB().Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		ctx.JSON(500, base.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
		return
	}
	//判断密码是否一致

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPasswordHash()), []byte(u.OldPassword)); err == nil {
		err = biz.DB().Model(&user).Updates(map[string]interface{}{"PasswordHash": u.PasswordHash, "IsTemporaryPassword": false}).Error
		if err != nil {
			ctx.JSON(500, base.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
			return
		}
	} else {
		ctx.JSON(500, base.SysResult{500, fmt.Sprintf("修改密码失败!"), u.ID})
		return
	}
	ctx.JSON(200, base.SysResult{200, "", nil})
}
func (ctrl *UserController) Plaintext(c *gin.Context) {
	//u := &models.User{}
	//biz := ctrl.Prepare(c)
	//id := c.Query("id")
	//err := biz.DB().Where("id=?", id).First(&u).Error
	//if err != nil {
	//	c.JSON(500, base.SysResult{500, err.Error(), err})
	//	return
	//}
	//
	//
	//c.JSON(200, base.SysResult{200, "", row})

}

// 修改用户信息
func (ctrl *UserController) Update(c *gin.Context) {
	row := &models.User{}
	ctrl.BaseController.Update(row, c)
}
func (ctrl *UserController) Login(ctx *gin.Context) {
	var loginParam userLogin
	err := ctx.ShouldBind(&loginParam)
	if err != nil {
		ctx.JSON(400, base.SysResult{Code: 400, Msg: "绑定User对象错误!", Data: err.Error()})
		return
	}
	biz := ctrl.Prepare(ctx)
	//获取要修改密码的用户信息
	user := &models.User{}
	err = biz.DB().Where("account = ?", loginParam.Username).First(&user).Error
	if err != nil {
		slog.Error("用户登录失败!", "username", loginParam.Username, "error", err.Error())
		ctx.JSON(500, base.SysResult{Code: 500, Msg: fmt.Sprintf("用户%s登录失败!", loginParam.Username), Data: loginParam.Username})
		return
	}
	//判断密码是否一致

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPasswordHash()), []byte(loginParam.Password)); err == nil {
		claims := &CustomClaims{}
		claims.Subject = user.Id.String()
		claims.DisplayName = user.DisplayName
		claims.Email = user.Email
		claims.EmailVerified = user.EmailConfirmed
		claims.OU = user.OU
		claims.OUID = user.OUId.String()

		jwt := NewJWT()
		token, err := jwt.CreateToken(*claims)
		if err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{
			"token": token,
			"sub":   user.Id,
			"name":  user.DisplayName,
			"ou":    user.OU,
			"email": user.Email,
		})
		return

	} else {
		ctx.JSON(401, base.SysResult{401, fmt.Sprintf("用户%s登录失败!", loginParam.Username), loginParam.Username})
		return
	}

}
func (ctrl *UserController) ResetPassword(ctx *gin.Context) {
	request := struct {
		Username string `json:"username" form:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBind(&request); err != nil {

		ctx.JSON(500, base.SysResult{
			Code: 500,
			Msg:  "Reset Password fail",
			Data: "username is empty",
		})
		return
	}
	username := request.Username

	biz := ctrl.Prepare(ctx)
	//generate password with random string
	newPassword := utils.GenerateRandomString(8)
	newHash, err := utils.HashPassword(newPassword)
	if err != nil {
		ctx.JSON(500, base.SysResult{
			Code: 500,
			Msg:  "Reset Password fail",
			Data: err.Error(),
		})
		return
	}
	u := map[string]interface{}{
		"password_hash":         newHash,
		"is_temporary_password": true,
	}
	err = biz.DB().Model(&models.User{}).Where("account = ?", username).Updates(u).Error
	if err != nil {
		ctx.JSON(500, base.SysResult{
			Code: 500,
			Msg:  "Reset Password fail",
			Data: err.Error(),
		})
		return
	}
	ctx.JSON(200, base.SysResult{
		Code: 200,
		Msg:  "Reset Password success",
		Data: newPassword,
	})
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
