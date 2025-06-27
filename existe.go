// Package existe possui uma função [Existe] que verifica a existência de uma
// dada chave, para tipos complexos com maps, struct e slice.
package existe

import (
	"reflect"
	"strconv"
	"strings"
)

// Existe verifica a existência de uma determinada propriedade em tipos
// complexos com maps, struct, e slice.
func Existe(v any, key string) bool {
	keys := strings.Split(key, ".")
	rv := reflect.ValueOf(v)

	for _, key := range keys {
		if !rv.IsValid() {
			return false
		}
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		switch rv.Kind() {
		case reflect.Map:
			rv = rv.MapIndex(reflect.ValueOf(key))
		case reflect.Slice:
			idx, err := strconv.Atoi(key)
			if err != nil || idx < 0 || idx >= rv.Len() {
				return false
			}
			rv = rv.Index(idx)
		case reflect.Struct:
			rv = rv.FieldByName(key)
		default:
			return false
		}
	}

	return rv.IsValid()
}
