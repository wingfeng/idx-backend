package utils

import (
	"errors"
	"reflect"
)

type TreeItem interface {
	GetID() interface{}
	ParentID() interface{}
	SetChildren(children []interface{})
}

func ConvertToTreeSlice(items interface{}) ([]TreeItem, error) {
	result := make([]TreeItem, 0)
	if reflect.ValueOf(items).Kind() != reflect.Slice { //判断是否为slice
		return result, errors.New("items is not slice")
	}
	sliceItems := reflect.ValueOf(items)
	for i := 0; i < sliceItems.Len(); i++ {
		item := sliceItems.Index(i)

		ti := item.Addr().Interface()
		it := ti.(TreeItem)

		result = append(result, it)

	}
	return result, nil
}
func BuildTree(items []TreeItem) []interface{} {
	var roots []TreeItem
	roots = make([]TreeItem, 0)

	for _, it := range items {

		p := it.ParentID()
		if p == nil {
			roots = append(roots, it)
		}

	}

	result := make([]interface{}, 0)

	for _, item := range roots {
		id := item.GetID()

		deep := 0
		children := getChildren(id, items, &deep)
		item.SetChildren(children)
		result = append(result, item)
	}
	return result
}
func getChildren(id interface{}, items []TreeItem, deep *int) []interface{} {
	*deep++
	if *deep > 20 { //防止死循环，树的层次最多20层
		return make([]interface{}, 0)
	}
	var children []TreeItem

	for _, it := range items {

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
