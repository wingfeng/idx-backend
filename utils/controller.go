package utils

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
)

type Controller interface {
	RegisterRouters(api *gin.RouterGroup)
}
type BaseController struct {
}

func (c BaseController) Prepare(ctx *gin.Context) *BizContext {
	biz, _ := ctx.Get(Const_BizContextKey)

	enf, existEnf := ctx.Get(Const_CasbinKey)
	bizContext := biz.(*BizContext)
	bizContext.User = ctx.GetString(Const_UserNameKey)
	bizContext.UserID = ctx.GetString(Const_UserIDKey)
	bizContext.OU = ctx.GetString(Const_OUKey)
	bizContext.OUID = ctx.GetString(Const_OUIDKey)
	if existEnf {
		bizContext.Enforcer = enf.(*casbin.Enforcer)
	}
	return bizContext
}

func (ctrl *BaseController) Save(row interface{}, c *gin.Context) {

	err := c.BindJSON(row)
	if err != nil {
		log.Errorf("绑定User对象错误!,%v", err.Error())
		c.AbortWithError(500, err)
		return
	}
	biz := ctrl.Prepare(c)
	err = biz.DB().Save(row).Error
	if err != nil {
		c.JSON(500, SysResult{500, "Error", err.Error()})
		return
	}
	c.JSON(200, SysResult{200, "Success", nil})
}

func (ctrl *BaseController) Update(row interface{}, c *gin.Context) {
	biz := ctrl.Prepare(c)
	//获取修改信息
	err := c.BindJSON(row)
	if err != nil {
		log.Errorf("绑定User对象错误!,%v", err.Error())
		c.AbortWithError(500, err)
		return
	}
	// if row.GetID() == nil || row.GetID() == "" {
	// 	c.JSON(500, SysResult{500, fmt.Sprintf("删除失败!"), ""})
	// 	return
	// }
	db := biz.DB().Updates(row)
	err = db.Error
	affect := db.RowsAffected
	if err != nil {
		c.JSON(500, SysResult{500, fmt.Sprintf("修改%v Error", row), err.Error()})
		return
	}
	if affect == 0 {
		c.JSON(500, SysResult{500, fmt.Sprintf("删除失败，没有相关记录，%v", row), row})
		return
	}
	c.JSON(200, SysResult{200, "", ""})
}

func (ctrl *BaseController) Get(row interface{}, c *gin.Context) {
	biz := ctrl.Prepare(c)
	id := c.Query("id")
	//	row.SetID(id)
	err := biz.DB().Where("id=?", id).First(row).Error
	if err != nil {
		c.JSON(500, SysResult{500, err.Error(), err})
		return
	}
	c.JSON(200, SysResult{200, "", row})
}
func (ctrl *BaseController) Delete(row interface{}, c *gin.Context) {
	biz := ctrl.Prepare(c)
	id := c.Query("id")

	db := biz.DB().Where("id=?", id).Delete(row)
	err := db.Error
	affect := db.RowsAffected
	if err != nil {
		c.JSON(500, SysResult{500, fmt.Sprintf("删除%v Error", row), err.Error()})
		return
	}
	if affect == 0 {
		c.JSON(500, SysResult{500, fmt.Sprintf("删除失败，没有相关记录，%v", id), id})
		return
	}
	c.JSON(200, SysResult{200, "删除成功", id})
}
func (ctrl *BaseController) Page(rows interface{}, c *gin.Context) {
	p := &Page{
		CurPage:  1,
		PageSize: 10,
		Data:     rows,
	}
	err := c.ShouldBind(p)
	biz := ctrl.Prepare(c)

	err = biz.Page(rows, p)
	if err != nil {
		c.JSON(500, SysResult{500, "Error", err.Error()})
		return
	}

	c.JSON(200, SysResult{200, "", p})
}

func (ctrl *BaseController) GeneratID(c *gin.Context) {
	id := GeneratID()
	idMap := make(map[string]string)
	idMap["id"] = id
	c.JSON(200, SysResult{200, "", idMap})
}
