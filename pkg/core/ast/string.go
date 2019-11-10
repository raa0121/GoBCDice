package ast

import (
	"fmt"
)

// 文字列のノード。
type String struct {
	NodeImpl
	NonNilNode
	ConstNode

	// 文字列
	Value string
}

// String がNodeを実装していることの確認。
var _Node = (*String)(nil)

// NewString は新しい文字列のノードを返す。
//
// value: 文字列,
func NewString(value string) *String {
	return &String{
		NodeImpl: NodeImpl{
			nodeType:            STRING_NODE,
			isPrimaryExpression: true,
		},

		Value: value,
	}
}

// SExp はノードのS式を返す。
func (n *String) SExp() string {
	return fmt.Sprintf("%q", n.Value)
}
