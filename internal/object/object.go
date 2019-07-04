package object

import (
	"fmt"
)

// オブジェクトの種類を表す型
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	SUCCESS_OBJ = "SUCCESS"
)

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
type Success struct {
	// 値
	Value bool
}

func (s *Success) Type() ObjectType {
	return SUCCESS_OBJ
}

func (s *Success) Inspect() string {
	str := "failure"
	if s.Value == true {
		str = "success"
	}

	return str
}
