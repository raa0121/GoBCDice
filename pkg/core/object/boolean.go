package object

import (
	"fmt"
)

// 論理値オブジェクトの構造体。
type Boolean struct {
	// 値
	Value bool
}

// NewBoolean は新しい論理値オブジェクトを返す。
func NewBoolean(v bool) *Boolean {
	return &Boolean{Value: v}
}

// Type はオブジェクトの種類を返す。
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
