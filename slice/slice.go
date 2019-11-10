package slice

import (
	"reflect"
)

// Содержит ли слайс элемент
func IsContains(s interface{}, elem interface{}) bool {
	sVal := reflect.ValueOf(s)
	if sVal.Kind() == reflect.Slice {
		for i := 0; i < sVal.Len(); i++ {
			// !!! panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if sVal.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

//
func DeleteElements(aS []string, bS ...string) []string {
	res := make([]string, 0, len(aS))
	for _, a := range aS {
		for _, b := range bS {
			if a == b {
				goto LOOP
			}
		}
		res = append(res, a)
	LOOP:
	}
	return res
}

/*
func DelEleInSlice(arr interface{}, index int) {
    vField := reflect.ValueOf(arr)
    value := vField.Elem()
    if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
        result := reflect.AppendSlice(value.Slice(0, index), value.Slice(index+1, value.Len()))
        value.Set(result)
    }
}
*/
