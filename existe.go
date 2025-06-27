package existe

import (
	"reflect"
	"strconv"
	"strings"
)

func Existe(v any, key string) bool {
	keys := strings.Split(key, ".")
	rv := reflect.ValueOf(v)
	return existeMesmo(rv, keys)
}

func existeMesmo(v reflect.Value, keys []string) bool {
	n := len(keys)
	if n == 0 || !v.IsValid() {
		return false
	}
	if v.Kind() == reflect.Interface {
		return existeMesmo(v.Elem(), keys)
	}

	key := keys[0]

	var kv reflect.Value

	switch v.Kind() {
	case reflect.Map:
		kv = v.MapIndex(reflect.ValueOf(key))
	case reflect.Slice:
		i, err := strconv.Atoi(key)
		if err != nil || i < 0 || i >= v.Len() {
			return false
		}
		kv = v.Index(i)
	case reflect.Struct:
		kv = v.FieldByName(key)
	default:
		return false
	}
	if !kv.IsValid() {
		return false
	}
	if n == 1 {
		return true
	}
	return existeMesmo(kv, keys[1:])
}
