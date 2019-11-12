package slice

import (
	"reflect"
	"strings"
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

// Возвращает новый слайс из элементов aS  без bS
// если элемент bS заканчивается на  "*"  - то шаблон
func DeleteElements(aS []string, bS ...string) []string {
	res := make([]string, 0, len(aS))
	for _, a := range aS {
		for j := 0; j < len(bS); j++ {
			// если заканчивается на "*" и длина больше одного символа
			if len(bS[j]) < 2 {
				continue
			}
			if bS[j][len(bS[j])-1:] == "*" {
				// возьмем до звездочки
				if strings.HasPrefix(a, bS[j][:len(bS[j])-1]) {
					//log.Printf("!!!%v %v", a, bS[j][:len(bS[j])-1])
					goto LOOP
				}
			} else if a == bS[j] {
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
