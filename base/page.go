package base

import "strings"

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
	//以field:operator的形势拼接，如：field1:like,field2:eq， Operator:eq,like,gt,lt,ge,le
	Filters string `form:"filters"`
	//多个参数用分号分隔 如: 内容1;内容2;
	Args string `form:"args"`
	//页面显示开始记录数
	FirstResult int64
	//页面显示最后记录数
	LastResult int64
	//排序类型
	OrderType string `form:"sortOrder"`
	//排序名称
	OrderName string      `form:"sortField"`
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
 *  the OrderName
 */
func (p *Page) GetOrderName() string {
	return p.OrderName
}

/**
 *  OrderName
 *            the OrderName to set
 */
func (p *Page) SetOrderName(orderName string) {
	p.OrderName = orderName
}

/**
 *  the orderBy
 */
func (p *Page) getOrderType() string {
	return strings.TrimSuffix(p.OrderType, "end")
}

/**
 *  orderBy
 *            the orderBy to set
 */
func (p *Page) SetOrderType(orderType string) {
	p.OrderType = orderType
}

/**
 *  the orderBy
 */
func (p *Page) GetOrderBy() string {
	if len(p.GetOrderName()) <= 0 {
		return ""
	}
	orderBy := p.GetOrderName() + " " + p.getOrderType()
	return orderBy
}

// func (p *Page) GetFilters() (string, error) {
// 	if len(p.Filters) < 1 {
// 		return "", nil
// 	}
// 	result, err := url.QueryUnescape(p.Filters)
// 	return string(result), err
// }
func (p *Page) GetArgs() []interface{} {
	if len(p.Args) < 1 {
		return nil
	}

	sResult := strings.Split(p.Args, " ")
	result := make([]interface{}, 0)
	for _, s := range sResult {
		result = append(result, s)
	}
	return result
}
