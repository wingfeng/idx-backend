package utils

import "strings"

type Page struct {
	//记录总数
	Counts int64 `json:"totalSize" `
	//每页显示记录数
	PageSize int64 `json:"rows" form:"rows"`
	//总页数
	TotalPage int64           `json:"totalPageSize"`
	Cols      map[string]bool `form:"cols"`
	//当前页
	CurPage int64  `json:"page" form:"page"`
	Filters string `form:"filters"`
	Args    string `form:"args"`
	//页面显示开始记录数
	FirstResult int64
	//页面显示最后记录数
	LastResult int64
	//排序类型
	OrderType string `form:"ordertype"`
	//排序名称
	OrderName string      `form:"ordername"`
	Data      interface{} `json:"content"`
}

func (this *Page) Build(counts int64, pageSize int64) {
	this.Counts = counts
	this.PageSize = pageSize
	if counts%pageSize == 0 {
		this.TotalPage = this.Counts / this.PageSize
	} else {
		this.TotalPage = this.Counts/this.PageSize + 1
	}
}

func (this *Page) GetCounts() int64 {
	return this.Counts
}

/**
 *  Counts
 *            the Counts to set
 */
func (this *Page) SetCounts(counts int64) {
	// 计算所有的页面数
	this.Counts = counts
	// this.TotalPage = (int)Math.ceil((this.Counts + this.perPageSize - 1)
	// / this.perPageSize)
	if counts%this.PageSize == 0 {
		this.TotalPage = this.Counts / this.PageSize
	} else {
		this.TotalPage = this.Counts/this.PageSize + 1
	}
}

func (this *Page) GetPageSize() int64 {
	return this.PageSize
}

func (this *Page) SetPageSize(pageSize int64) {
	this.PageSize = pageSize
}

/**
 *  the TotalPage
 */
func (this *Page) GetTotalPage() int64 {
	if this.TotalPage < 1 {
		return 1
	}
	return this.TotalPage
}

/**
 *  TotalPage
 *            the TotalPage to set
 */
func (this *Page) SetTotalPage(totalPage int64) {
	this.TotalPage = totalPage
}

func (this *Page) GetCurPage() int64 {
	return this.CurPage
}

func (this *Page) SetCurPage(curPage int64) {
	this.CurPage = curPage
}

/**
 *  the FirstResult
 */
func (this *Page) GetFirstResult() int64 {
	temp := this.CurPage - 1
	if temp <= 0 {
		return 0
	}
	this.FirstResult = (this.CurPage - 1) * this.PageSize
	return this.FirstResult
}

/**
 *  FirstResult
 *            the FirstResult to set
 */
func (this *Page) SetFirstResult(firstResult int64) {
	this.FirstResult = firstResult
}

/**
 *  the LastResult
 */
func (this *Page) GetLastResult() int64 {
	this.LastResult = this.FirstResult + this.PageSize
	return this.LastResult
}

/**
 *  LastResult
 *            the LastResult to set
 */
func (this *Page) SetLastResult(lastResult int64) {
	this.LastResult = lastResult
}

/**
 *  the OrderName
 */
func (this *Page) GetOrderName() string {
	return this.OrderName
}

/**
 *  OrderName
 *            the OrderName to set
 */
func (this *Page) SetOrderName(orderName string) {
	this.OrderName = orderName
}

/**
 *  the orderBy
 */
func (this *Page) getOrderType() string {
	return this.OrderType
}

/**
 *  orderBy
 *            the orderBy to set
 */
func (this *Page) SetOrderType(orderType string) {
	this.OrderType = orderType
}

/**
 *  the orderBy
 */
func (this *Page) GetOrderBy() string {
	if len(this.GetOrderName()) <= 0 {
		return ""
	}
	orderBy := this.GetOrderName() + " " + this.getOrderType()
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
