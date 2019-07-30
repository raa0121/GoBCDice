package object

import (
	"fmt"
)

// 整数オブジェクトの構造体。
type Integer struct {
	// 数値
	Value int
}

// Type はオブジェクトの種類を返す。
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
