package cmd

const usualCommonValidator = `package common

import (
	"regexp"
	"reflect"
	"runtime"
	"strconv"
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


func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

// MyCaller returns the caller of the function that called it :)
func MyCaller() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	return getFrame(2).Function
}


func UniqueStringArray(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func UniqueIntArray(slice []int) []int {
	// create a map with all the values as key
	uniqMap := make(map[int]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]int, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func ReverseBytes(in []byte) []byte {
	res := []byte{}

	for i := len(in) - 1; i >= 0; i-- {
		res = append(res, in[i])
	}
	return res
}

func IntArrayToStringArray(in []int) (res []string) {
	for _, val := range in {
		res = append(res, strconv.Itoa(val))
	}
	return
}

func StringArrayToIntArray(in []string) (res []int) {
	for _, val := range in {
		if v, err := strconv.Atoi(val); err == nil {
			res = append(res, v)
		}
	}
	return
}

`

const usualCommonGenerator = `package common

import (
    "math/rand"
    "time"
	"reflect"
	"strings"
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

func GetTypeName(v interface{}) string {
	return strings.Replace(reflect.TypeOf(v).String(), "*", "", -1)
}

func GetFieldName(structPoint interface{}, fieldPinter interface{}) (name string) {

	val := reflect.ValueOf(structPoint).Elem()
	val2 := reflect.ValueOf(fieldPinter).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if valueField.Addr().Interface() == val2.Addr().Interface() {
			return val.Type().Field(i).Name
		}
	}

	return
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
