package utilx

// InSliceStr 判断slice中是否存在某元素，仅支持小切片
func InSliceStr(val string, s []string) bool {
	for _, v := range s {
		if val == v {
			return true
		}
	}
	return false
}

// InSliceInt64 判断slice中是否存在某元素，仅支持小切片
func InSliceInt64(val int64, s []int64) bool {
	for _, v := range s {
		if val == v {
			return true
		}
	}
	return false
}
