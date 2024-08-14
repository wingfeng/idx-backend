package base

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"

	"database/sql"

	"github.com/lunny/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	const_CreatorField    = "Creator"
	const_CreatorIDField  = "CreatorID"
	const_ModifyerField   = "Updator"
	const_ModifyerIDField = "UpdatorID"
)

type BizContext struct {
	//UserID 当前用户ID
	UserID string
	//User 当前用户的名称
	User string
	//当前用户的部门名称
	OU string
	//当前用户的部门ID
	OUID string

	db       *gorm.DB
	Enforcer *casbin.Enforcer
	node     snowflake.Node
	store    map[string]interface{}
}

func (c *BizContext) DB() *gorm.DB {
	db := c.db.Set(Const_UserIDKey, c.UserID)
	db = db.Set(Const_UserNameKey, c.User)
	db = db.Set(Const_OUKey, c.OU)
	db = db.Set(Const_OUIDKey, c.OUID)
	return db
}
func InitContext(driver string, connection string, user string, userId string, dbLogger logger.Interface) *BizContext {
	context := &BizContext{
		User:   user,
		UserID: userId,
	}

	context.store = make(map[string]interface{})

	engine := GetDB(driver, connection)
	if dbLogger != nil {
		engine.Logger = dbLogger
	}

	// engine = engine.Set("user", user)
	// engine = engine.Set("userid", userId)
	context.db = engine
	node, err := snowflake.NewNode(1)
	if err != nil {
		slog.Error("创建ID生成器失败")
	}
	context.node = *node
	return context
}
func (c *BizContext) Set(name string, obj interface{}) {
	c.store[name] = obj
}
func (c *BizContext) Get(name string) interface{} {
	obj := c.store[name]
	return obj
}
func GetDB(driver string, connection string) *gorm.DB {
	if strings.EqualFold(driver, "") {
		driver = "mysql"
	}
	var err error
	var x *gorm.DB

	sqlDB, err := sql.Open(driver, connection)

	switch driver {
	case "mysql":

		x, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

	case "pgx":
		x, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

	}

	if nil != err {
		log.Error("init" + err.Error())
	}
	return x
}

/*
*
分页查询
*/
func (context *BizContext) Page(rows interface{}, page *Page) error {

	//声明结果变量
	var err error
	var counts int64
	//获取总记录数处理

	db := context.DB().Model(rows)

	filters := page.Filters

	args := page.GetArgs()

	if !strings.EqualFold(filters, "") {
		db = db.Where(filters, args...)
	}
	err = db.Count(&counts).Error
	if nil != err {
		return err
	} else {
		page.SetCounts(counts)
	}
	if len(page.Cols) > 0 {

		for a, c := range page.Cols {
			if c {
				db.Statement.Selects = append(db.Statement.Selects, a)
			}
		}

	}
	//排序处理
	orderBy := page.GetOrderBy()

	if len(orderBy) > 0 {
		db = db.Order(orderBy)
	}
	offset := int(page.GetFirstResult())
	limit := int(page.PageSize)
	err = db.Offset(offset).Limit(limit).Find(rows).Error
	page.Data = rows

	return err
}

// PageComplex 复杂的查询分页
// @db 通过在外面先构建好基础查询功能的DB在再来实现分页功能
//
//sample:db := context.DB().Table("wf_steps steps").Select("steps.*,inst.subject").Joins("join wf_instances inst on steps.instance_id=inst.id").Where("steps.executor=? and steps.status<3", "wing")
func (context *BizContext) PageComplex(db *gorm.DB, page *Page) error {
	//声明结果变量
	var err error
	var counts int64

	filters := page.Filters

	args := page.GetArgs()

	if !strings.EqualFold(filters, "") {
		db = db.Where(filters, args...)
	}
	err = db.Count(&counts).Error
	if nil != err {
		return err
	} else {
		page.SetCounts(counts)
	}
	if len(page.Cols) > 0 {

		for a, c := range page.Cols {
			if c {
				db.Statement.Selects = append(db.Statement.Selects, a)
			}
		}

	}
	//排序处理
	orderBy := page.GetOrderBy()

	if len(orderBy) > 0 {
		db = db.Order(orderBy)
	}
	offset := int(page.GetFirstResult())
	limit := int(page.PageSize)

	if page.Data == nil {
		var results []map[string]interface{}
		page.Data = &results
	}

	err = db.Offset(offset).Limit(limit).Find(page.Data).Error

	return err
}
