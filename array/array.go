package array

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

// 删除重复项的内容
func RemoveDuplicate[T comparable](list []T) []T {
	if len(list) < 2 {
		return list
	}

	keys := map[T]struct{}{}
	result := make([]T, len(list))
	ri := 0
	for i := 0; i < len(list); i++ {
		if _, has := keys[list[i]]; !has {
			result[ri] = list[i]
			keys[list[i]] = struct{}{}
			ri++
		}
	}
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
