package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 整数のノード。
// 一次式。
type Int struct {
	NodeImpl

	// 数値
	Value int
}

// Int がNodeを実装していることの確認。
var _ Node = (*Int)(nil)

// Type はノードの種類を返す。
func (n *Int) Type() NodeType {
	return INT_NODE
}

// SExp はノードのS式を返す。
func (n *Int) SExp() string {
	return fmt.Sprintf("%d", n.Value)
}

// IsPrimaryExpression は一次式かどうかを返す。
// Intではtrueを返す。
func (n *Int) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// Intではfalseを返す。
func (n *Int) IsVariable() bool {
	return false
}

// NewInt は新しい整数のノードを返す。
//
// value: 数値,
// tok: 対応するトークン。
func NewInt(value int, tok token.Token) *Int {
	return &Int{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		Value: value,
	}
}

// NewInt2 は新しい整数のノードを返す。
//
// value: 数値,
func NewInt2(value int) *Int {
	return &Int{
		Value: value,
	}
}
