package utilx

import "reflect"

func IsMap(i interface{}) bool {
	t := reflect.TypeOf(i)
	return t == reflect.TypeOf(make(map[string]interface{}))
}
