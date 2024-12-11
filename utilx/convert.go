package utilx

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// MustString 类型转换-返回string
func MustString(v interface{}, defaultval ...string) string {
	val, ok := TryString(v)
	if ok {
		return val
	}
	if len(defaultval) > 0 {
		return defaultval[0]
	}
	return ""
}

// TryString 类型转换-转换为string
func TryString(v interface{}) (string, bool) {
	switch tv := v.(type) {
	case string:
		return tv, true
	case []byte:
		return string(tv), true
	case int64:
		return strconv.FormatInt(int64(tv), 10), true
	case uint64:
		return strconv.FormatUint(uint64(tv), 10), true
	case int32:
		return strconv.FormatInt(int64(tv), 10), true
	case uint32:
		return strconv.FormatUint(uint64(tv), 10), true
	case int16:
		return strconv.FormatInt(int64(tv), 10), true
	case uint16:
		return strconv.FormatUint(uint64(tv), 10), true
	case int8:
		return strconv.FormatInt(int64(tv), 10), true
	case uint8:
		return strconv.FormatUint(uint64(tv), 10), true
	case float32:
		return strconv.FormatFloat(float64(tv), 'f', -1, 64), true
	case float64:
		return strconv.FormatFloat(float64(tv), 'f', -1, 64), true
	case int:
		return strconv.Itoa(int(tv)), true
	case json.Number:
		return tv.String(), true
	case bool:
		if tv {
			return "true", true
		} else {
			return "false", true
		}
	}
	return "", false
}

// StructToMap 结构体转换为 map[string]interface{}
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	if data == nil {
		return result
	}
	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := field.Tag.Get("json") //优先读取json标签内容
		if fieldName == "" {
			fieldName = field.Name
		}
		fieldValue := val.Field(i).Interface()

		result[fieldName] = fieldValue
	}

	return result
}

// MapToStruct map[string]interface{} 转换为结构体
func MapToStruct(data map[string]interface{}, result interface{}) error {
	val := reflect.ValueOf(result)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("result must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue, ok := data[fieldName]
		if !ok {
			// 如果 map 中没有结构体字段对应的键，跳过
			continue
		}

		val.Field(i).Set(reflect.ValueOf(fieldValue))
	}
	return nil
}

// MustInt64 强制返回int64
func MustInt64(v interface{}, defaultval ...int64) int64 {
	var defaultValue int64 = 0
	if len(defaultval) > 0 {
		defaultValue = defaultval[0]
	}
	if v == nil {
		return defaultValue
	}
	switch tv := v.(type) {
	case float32:
		if tv > float32(math.MaxInt64) {
			return defaultValue
		}
		return int64(tv)
	case float64:
		if tv > float64(math.MaxInt64) {
			return defaultValue
		}
		return int64(tv)
	}
	val, ok := TryInt64(v)
	if ok {
		return val
	}
	return defaultValue
}

// TryInt64 类型转换-int64
func TryInt64(v interface{}) (int64, bool) {
	if v == nil {
		return -1, false
	}
	switch tv := v.(type) {
	case []byte:
		res, err := strconv.ParseInt(string(tv), 10, 0)
		if err != nil {
			return -1, false
		}
		return res, true
	case string:
		res, err := strconv.ParseInt(tv, 10, 0)
		if err != nil {
			return -1, false
		}
		return res, true
	case int64:
		return tv, true
	case uint64:
		if tv > uint64(math.MaxInt64) {
			return -1, false
		}
		return int64(tv), true
	case int32:
		return int64(tv), true
	case uint32:
		return int64(tv), true
	case int:
		return int64(tv), true
	case int16:
		return int64(tv), true
	case uint16:
		return int64(tv), true
	case int8:
		return int64(tv), true
	case uint8:
		return int64(tv), true
	case json.Number:
		val, err := tv.Int64()
		if err == nil {
			return val, true
		}
	}
	return -1, false
}
