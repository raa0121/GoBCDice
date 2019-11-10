package ast

import (
	"fmt"
)

// 整数のノード。
// 一次式。
type Int struct {
	NodeImpl
	NonNilNode
	ConstNode

	// 数値
	Value int
}

// Int がNodeを実装していることの確認。
var _ Node = (*Int)(nil)

// NewInt は新しい整数のノードを返す。
//
// value: 数値,
func NewInt(value int) *Int {
	return &Int{
		NodeImpl: NodeImpl{
			nodeType:            INT_NODE,
			isPrimaryExpression: true,
		},

		Value: value,
	}
}

// SExp はノードのS式を返す。
func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}
