package sliceutils

import (
	"github.com/pkg/errors"
	"reflect"
)

func StringSlice2InterfaceSlice(strSlice []string) []interface{} {
	ret := make([]interface{}, len(strSlice))
	for i, v := range strSlice {
		ret[i] = v
	}
	return ret
}

func InterfaceSlice2StringSlice(strSlice []interface{}) []string {
	ret := make([]string, len(strSlice))
	for i, v := range strSlice {
		ret[i] = v.(string)
	}
	return ret
}

func Contains(src []interface{}, test interface{}) bool {
	for _, item := range src {
		if item == test {
			return true
		}
	}
	return false
}

func ContainsDeep(src []interface{}, test interface{}) bool {
	for _, item := range src {
		if reflect.DeepEqual(item, test) {
			return true
		}
	}
	return false
}

func StringContains(src []string, test string) bool {
	for _, item := range src {
		if item == test {
			return true
		}
	}
	return false
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func IndexOfAny(target interface{}, anySlice interface{}) (int, error) {
	if reflect.TypeOf(anySlice).Kind() != reflect.Slice {
		return -1, errors.New("not slice")
	}
	data := reflect.ValueOf(anySlice)
	for i := 0; i < data.Len(); i++ {
		elem := data.Index(i)
		if elem.Interface() == target {
			return i, nil
		}
	}
	return -1, nil //not found.
}

func IsEmpty(src interface{}) bool {
	if slice, ok := takeSliceArg(src); ok {
		return slice == nil || len(slice) == 0
	}
	panic("not slice")
}

// https://ahmet.im/blog/golang-take-slices-of-any-type-as-input-parameter/
func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func ConvertAny2Interface(src interface{}) ([]interface{}, error) {
	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return nil, errors.New("Src not slice")
	}
	data := reflect.ValueOf(src)
	ret := make([]interface{}, data.Len())
	for i := 0; i < data.Len(); i++ {
		ret[i] = data.Index(i).Interface()
	}
	return ret, nil
}
