/*
BCDiceコマンドの評価結果として生成される数値などのオブジェクトの内部表現のパッケージ。
*/
package object

import (
	"fmt"
)

// オブジェクトの種類を表す型。
type ObjectType int

// String はオブジェクトの種類を文字列として返す。
func (t ObjectType) String() string {
	if str, ok := objectTypeString[t]; ok {
		return str
	}

	return objectTypeString[ILLEGAL_OBJ]
}

const (
	ILLEGAL_OBJ ObjectType = iota
	INTEGER_OBJ
	BOOLEAN_OBJ
)

// オブジェクトの種類とそれを表す文字列との対応
var objectTypeString = map[ObjectType]string{
	ILLEGAL_OBJ: "ILLEGAL",

	INTEGER_OBJ: "INTEGER",
	BOOLEAN_OBJ: "BOOLEAN",
}

// オブジェクトが持つインターフェース。
type Object interface {
	// Type はオブジェクトの種類を返す。
	Type() ObjectType
	// Inspect はオブジェクトの内容を文字列として返す。
	Inspect() string
}

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
