package array

import (
	"reflect"

	"github.com/nlbydcg/go-tools/tool"
)

// 判断是否存在某数据
func CheckHas[T comparable](list []T, v T) bool {
	if len(list) == 0 {
		return false
	}
	for i := 0; i < len(list); i++ {
		if list[i] == v {
			return true
		}
	}
	return false
}

// 删除某项数据
func RemoveItem[T comparable](list []T, t T) []T {
	result := make([]T, len(list))
	if len(list) == 0 {
		return result
	}

	ri := 0
	for i := 0; i < len(list); i++ {
		if t != list[i] { // 如何符合删除逻辑
			result[ri] = list[i]
			ri++
		}
	}
	return result[:ri]
}

// 按照方法删除数据
func RemoveItemByFunc[T comparable](list []T, fn func(t T) bool) []T {
	result := make([]T, len(list))
	if len(list) == 0 {
		return result
	}

	ri := 0
	for i := 0; i < len(list); i++ {
		if !fn(list[i]) { // 如何符合删除逻辑
			result[ri] = list[i]
			ri++
		}
	}
	return result[:ri]
}

// 删除重复项的内容
func RemoveDuplicate[T comparable](list []T) []T {
	if len(list) < 2 {
		return list
	}

	keys := map[T]struct{}{}
	result := make([]T, len(list))
	ri := 0

	addDuplicate(result, list, keys, ri)
	return result[:ri]
}

// 取出重复项的内容
func GetDuplicate[T comparable](list []T) []T {
	if len(list) < 2 {
		return []T{}
	}

	keys := map[T]struct{}{}
	result := make([]T, len(list))
	ri := 0
	for i := 0; i < len(list); i++ {
		if _, has := keys[list[i]]; has {
			result[ri] = list[i]
			ri++
		} else {
			keys[list[i]] = struct{}{}
		}
	}
	return result[:ri]
}

// 数组转string数组
func ToStrings[T any](list []T) []string {
	if !CheckIsArray(list) {
		// if not array , to thrown
		panic("list is not array")
	}
	result := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = tool.ToString(list[i])
	}
	return result
}

// 判断是否为数组
func CheckIsArray(list interface{}) bool {
	t := reflect.ValueOf(list)
	// 如果是指针 获取到对应的数据
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		return true
	}
	return false
}

// 合并数组
func Merge[T comparable](l1 []T, l2 ...T) []T {
	r := make([]T, len(l1)+len(l2))
	copy(r, l1)
	copy(r[len(l1):], l2)
	return r
}

// 合并数组并去重
func MergeDuplicate[T comparable](l1 []T, l2 ...T) []T {
	r := make([]T, len(l1)+len(l2))
	ri := 0
	keys := map[T]struct{}{}

	addDuplicate(r, l1, keys, ri)
	addDuplicate(r, l2, keys, ri)

	return r[:ri]
}

// 添加数据并合并
func addDuplicate[T comparable](r, l []T, keys map[T]struct{}, ri int) {
	for i := 0; i < len(l); i++ {
		if _, has := keys[l[i]]; !has {
			r[ri] = l[i]
			keys[l[i]] = struct{}{}
			ri++
		}
	}
}
