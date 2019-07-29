package common

import (
	"reflect"
)

func InArray (item interface{}, array interface{}) bool {

	rt := reflect.TypeOf(array)

	switch rt.Kind() {
	case reflect.Slice:
		return checkInArray(item, array)
	case reflect.Array:
		return checkInArray(item, array)
	default:
		return false
	}

}

func checkInArray(item interface{}, array interface{}) bool {

	listS, ok := array.([]string)
	if ok {
		for _, i := range listS {
			if i == item {
				return true
			}
		}
	}

	listI, ok := array.([]interface{})
	if ok {
		for _, i := range listI {
			if i == item {
				return true
			}
		}
	}

	return false
}