package util

import (
	"reflect"
)

// Clone はオブジェクトを複製する。
//
// 参考: https://arhipov.net/golang/2016/03/12/shallow-copying-interface-values-in-go.html
func Clone(i interface{}) interface{} {
	// Wrap argument to reflect.Value, dereference it and return back as interface{}
	return reflect.Indirect(reflect.ValueOf(i)).Interface()
}
