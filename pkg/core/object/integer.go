package object

import (
	"fmt"
)

// 整数オブジェクトの構造体。
type Integer struct {
	// 数値
	Value int
}

// NewInteger は新しい整数オブジェクトを返す。
func NewInteger(v int) *Integer {
	return &Integer{Value: v}
}

// Type はオブジェクトの種類を返す。
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Add は加算を行い、その結果を返す。
func (i *Integer) Add(j *Integer) *Integer {
	return NewInteger(i.Value + j.Value)
}
