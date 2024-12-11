package utilx

import "strconv"

// Fen2Yuan 分转元
func Fen2Yuan(amount int64) string {
	result := Fen2Y(amount) + "元"
	return result
}

// Fen2Y 分转元 返回 00.00  单位元
func Fen2Y(amount int64) string {
	amountInYuan := F2y(amount)

	// 将结果转换为字符串格式，并添加单位元
	result := strconv.FormatFloat(amountInYuan, 'f', 2, 64)
	return result
}

// F2y 分转元
func F2y(amount int64) float64 {
	return float64(amount) / 100
}
