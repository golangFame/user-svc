package utils

import "reflect"

func getTypeOf(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
