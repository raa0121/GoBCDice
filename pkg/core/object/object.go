package object

import (
	"fmt"
)

// オブジェクトの種類を表す型
type ObjectType int

func (t ObjectType) String() string {
	if str, ok := objectTypeString[t]; ok {
		return str
	}

	return objectTypeString[ILLEGAL_OBJ]
}

const (
	ILLEGAL_OBJ ObjectType = iota
	INTEGER_OBJ
	SF_OBJ
)

var objectTypeString = map[ObjectType]string{
	ILLEGAL_OBJ: "ILLEGAL",

	INTEGER_OBJ: "INTEGER",
	SF_OBJ:      "SF",
}

// オブジェクトが持つインターフェース
type Object interface {
	// オブジェクトの種類を返す
	Type() ObjectType
	// オブジェクトの内容を文字列として返す
	Inspect() string
}

// 整数オブジェクトの構造体
type Integer struct {
	//
	Value int
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// 成功/失敗オブジェクトの構造体
type SF struct {
	// 値
	Value bool
}

func (sf *SF) Type() ObjectType {
	return SF_OBJ
}

func (sf *SF) Inspect() string {
	str := "failure"
	if sf.Value == true {
		str = "success"
	}

	return str
}
