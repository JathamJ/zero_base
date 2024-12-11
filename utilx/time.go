package utilx

import (
	"database/sql"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

// ParseTime 文本转时间
func ParseTime(str string) time.Time {
	t, _ := TryParseTime(str)
	return t
}

func TryParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(TimeLayout, str, time.Local)
}

// NullTimeFormat 格式化sql.NullTime
func NullTimeFormat(t sql.NullTime) string {
	v, _ := t.Value()
	if v == nil {
		return ""
	}
	if ts, ok := v.(time.Time); ok {
		return ts.Format(TimeLayout)
	}
	return ""
}

func GetDays(str string) int {
	parseTime, _ := TryParseTime(str)
	// 获取当前时间
	currentTime := time.Now()
	// 计算两个时间点之间的持续时间
	duration := currentTime.Sub(parseTime)
	// 将持续时间转换为天数
	days := int(duration.Hours() / 24)
	return days
}
