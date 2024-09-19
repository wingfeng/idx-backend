package base

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Page struct {
	//记录总数
	Counts int64 `json:"total" `
	//每页显示记录数
	PageSize int64 `json:"pageSize" form:"pageSize"`
	//总页数
	TotalPage int64           `json:"totalPages"`
	Cols      map[string]bool `form:"cols"`
	//当前页
	CurPage int64 `json:"page" form:"page"`
	//过滤条件过滤条件只支持'field operator ?'格式
	Filters []string `json:"filters"`
	//多个参数用分号分隔 如: 内容1;内容2;
	Args []string `json:"args"`
	//页面显示开始记录数
	FirstResult int64
	//页面显示最后记录数
	LastResult int64
	//排序类型
	SortOrder string `json:"sortOrder"`
	//排序名称
	SortField string      `json:"sortField"`
	Data      interface{} `json:"list"`
}

func (p *Page) Build(counts int64, pageSize int64) {
	p.Counts = counts
	p.PageSize = pageSize
	if counts%pageSize == 0 {
		p.TotalPage = p.Counts / p.PageSize
	} else {
		p.TotalPage = p.Counts/p.PageSize + 1
	}
}

func (p *Page) GetCounts() int64 {
	return p.Counts
}

/**
 *  Counts
 *            the Counts to set
 */
func (p *Page) SetCounts(counts int64) {
	// 计算所有的页面数
	p.Counts = counts
	// p.TotalPage = (int)Math.ceil((p.Counts + p.perPageSize - 1)
	// / p.perPageSize)
	if counts%p.PageSize == 0 {
		p.TotalPage = p.Counts / p.PageSize
	} else {
		p.TotalPage = p.Counts/p.PageSize + 1
	}
}

func (p *Page) GetPageSize() int64 {
	return p.PageSize
}

func (p *Page) SetPageSize(pageSize int64) {
	p.PageSize = pageSize
}

/**
 *  the TotalPage
 */
func (p *Page) GetTotalPage() int64 {
	if p.TotalPage < 1 {
		return 1
	}
	return p.TotalPage
}

/**
 *  TotalPage
 *            the TotalPage to set
 */
func (p *Page) SetTotalPage(totalPage int64) {
	p.TotalPage = totalPage
}

func (p *Page) GetCurPage() int64 {
	return p.CurPage
}

func (p *Page) SetCurPage(curPage int64) {
	p.CurPage = curPage
}

/**
 *  the FirstResult
 */
func (p *Page) GetFirstResult() int64 {
	temp := p.CurPage - 1
	if temp <= 0 {
		return 0
	}
	p.FirstResult = (p.CurPage - 1) * p.PageSize
	return p.FirstResult
}

/**
 *  FirstResult
 *            the FirstResult to set
 */
func (p *Page) SetFirstResult(firstResult int64) {
	p.FirstResult = firstResult
}

/**
 *  the LastResult
 */
func (p *Page) GetLastResult() int64 {
	p.LastResult = p.FirstResult + p.PageSize
	return p.LastResult
}

/**
 *  LastResult
 *            the LastResult to set
 */
func (p *Page) SetLastResult(lastResult int64) {
	p.LastResult = lastResult
}

/**
 *  the orderBy
 */
func (p *Page) getOrderType() string {
	return strings.TrimSuffix(p.SortOrder, "end")
}

/**
 *  the orderBy
 */
func (p *Page) GetOrderBy() string {
	if len(p.SortField) <= 0 {
		return ""
	}
	orderBy := p.SortField + " " + p.getOrderType()
	return orderBy
}

func (p *Page) ValidateFilters() error {
	pattern := `^[a-zA-Z_][a-zA-Z0-9_]*$`
	supportedOperators := map[string]bool{
		"=":    true,
		"!=":   true,
		">":    true,
		"<":    true,
		">=":   true,
		"<=":   true,
		"like": true,
		"in":   true,
	}
	compiledPattern, _ := regexp.Compile(pattern)
	for _, filter := range p.Filters {
		if strings.LastIndex(filter, "?") != len(filter)-1 {
			return errors.New("? must be at the end of filter")
		}
		tmp := strings.Split(filter, " ")
		if len(tmp) != 3 {
			return errors.New("filter format error, only support 'field operator ?' pattern")
		}
		field := tmp[0]

		//通过正则表达时校验field是否符合数据库字段格式
		matched := compiledPattern.MatchString(field)
		if !matched {
			return fmt.Errorf("field name %s format error", field)
		}
		op := tmp[1]
		if !supportedOperators[op] {
			return fmt.Errorf("operator %s not supported", op)
		}

	}

	return nil
}
