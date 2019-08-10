package object

import (
	"fmt"
)

// 文字列オブジェクトの構造体。
type String struct {
	// 文字列の値
	Value string
}

// NewString は新しい文字列オブジェクトを返す。
func NewString(v string) *String {
	return &String{Value: v}
}

// Type はオブジェクトの種類を返す。
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (s *String) Inspect() string {
	return fmt.Sprintf("%q", s.Value)
}
