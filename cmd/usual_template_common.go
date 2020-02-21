package cmd

const usualCommonValidator = `package common

import (
	"regexp"
	"reflect"
)

func ValidateEmail(email string) bool {
    Re := regexp.MustCompile(` + "`" + `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$` + "`" + `)
    return Re.MatchString(email)
}

func ValidateMobile(phone string) bool {
    Re := regexp.MustCompile(` + "`" + `^[+][0-9]{11,}` + "`" + `)
    return Re.MatchString(phone)
}

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

	listInt, ok := array.([]int)
	if ok {
		for _, i := range listInt {
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
`

const usualCommonGenerator = `package common

import (
    "math/rand"
    "time"
)

func RandomString(size int) string {
    var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
    rand.Seed(time.Now().UTC().UnixNano())
    buf := make([]byte, size)
    for i := 0; i < size; i++ {
        buf[i] = alpha[rand.Intn(len(alpha))]
    }
    return string(buf)
}
`
var usualTemplateCommonValidator = template{
	Path:    "./common/validator.go",
	Content: usualCommonValidator,
}

var usualTemplateCommonGenerator = template{
	Path:    "./common/generator.go",
	Content: usualCommonGenerator,
}
