package object

import (
	"fmt"
)

// 論理値オブジェクトの構造体。
type Boolean struct {
	// 値
	Value bool
}

// Type はオブジェクトの種類を返す。
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
