package utility

import "reflect"

func Merge(dst interface{}, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()
	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		dstField := dstValue.Field(i)
		if dstField.CanSet() && srcField.Interface() != reflect.Zero(srcField.Type()).Interface() {
			dstField.Set(srcField)
		}
	}
}
