package utils

import (
	"reflect"
	"strings"
)

type TreeItem interface {
	GetID() interface{}
	ParentID() interface{}
	SetChildren(children []interface{})
}

func BuildMenuTree(items interface{}) []interface{} {
	var roots []TreeItem
	roots = make([]TreeItem, 0)

	sliceItems := reflect.ValueOf(items)
	for i := 0; i < sliceItems.Len(); i++ {
		item := sliceItems.Index(i)

		ti := item.Addr().Interface()
		it := ti.(TreeItem)

		p := it.ParentID()
		if strings.EqualFold("", p.(string)) || p == int64(0) {
			roots = append(roots, it)
		}

	}
	// linq.From(siceItems).WhereT(func(i TreeItem) bool {
	// 	return i.ParentID() == 0
	// }).ToSlice(&roots)

	result := make([]interface{}, 0)

	for _, item := range roots {
		id := item.GetID()
		if strings.EqualFold("", id.(string)) || id == int64(0) {
			continue
		}
		deep := 0
		children := getChildren(id, items, &deep)
		item.SetChildren(children)
		result = append(result, item)
	}
	return result
}
func getChildren(id interface{}, items interface{}, deep *int) []interface{} {
	*deep++
	if *deep > 20 { //防止死循环，树的层次最多20层
		return make([]interface{}, 0)
	}
	var children []TreeItem
	sliceItems := reflect.ValueOf(items)
	for i := 0; i < sliceItems.Len(); i++ {
		item := sliceItems.Index(i)

		ti := item.Addr().Interface()
		it := ti.(TreeItem)

		p := it.ParentID()
		if p == id {
			children = append(children, it)
		}

	}
	result := make([]interface{}, 0)
	for _, c := range children {
		children := getChildren(c.GetID(), items, deep)
		c.SetChildren(children)
		result = append(result, c)
	}
	return result
}
