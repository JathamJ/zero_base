package utilx

import (
	"math/rand"
	"time"
)

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

// Shuffle 打乱切片
func Shuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
