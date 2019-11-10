package slice

import "reflect"

// Содержит ли слайс элемент
func Contains(s interface{}, elem interface{}) bool {
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
