package tool

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToString[T any](v T) (s string) {
	t := reflect.ValueOf(v)

	switch t.Kind() {
	case reflect.String:
		s = t.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s = strconv.FormatInt(t.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		s = strconv.FormatUint(t.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		s = fmt.Sprintf("%f", t.Float())
	default:
		s = fmt.Sprintf("%#v", t.Interface())
	}
	return s
}
