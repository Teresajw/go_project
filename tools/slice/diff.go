package slice

// DiffSet 差集，只支持 comparable 类型
// 已去重
// 并且返回值的顺序是不确定的
func DiffSet[T comparable](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return []T{}
	}
	m := make(map[T]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	var diff []T
	for _, v := range b {
		_, ok := m[v]
		if !ok {
			diff = append(diff, v)
		}
	}
	return diff
}
