package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

type Int struct {
	NodeImpl

	// 数値
	Value int
}

// IntがNodeを実装していることの確認
var _ Node = (*Int)(nil)

// Typeはノードの種類を返す
func (n *Int) Type() NodeType {
	return INT_NODE
}

// SExpはノードのS式を返す
func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}

// IsPrimaryExpressionは一次式かどうかを返す
func (n *Int) IsPrimaryExpression() bool {
	return true
}

// IsVariableは可変ノードかどうかを返す。
func (n *Int) IsVariable() bool {
	return false
}

// NewIntは新しい整数ノードを返す
func NewInt(value int, tok token.Token) *Int {
	return &Int{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		Value: value,
	}
}
