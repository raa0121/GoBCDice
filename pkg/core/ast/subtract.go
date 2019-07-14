package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 減算のノード。
// 中置式。
type Subtract struct {
	InfixExpressionImpl
}

// Subtract がNodeを実装していることの確認。
var _ Node = (*Subtract)(nil)

// Subtract がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*Subtract)(nil)

// Type はノードの種類を返す。
func (n *Subtract) Type() NodeType {
	return SUBTRACT_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *Subtract) Precedence() OperatorPrecedenceType {
	return PREC_ADDITIVE
}

// IsLeftAssociative は左結合性かどうかを返す。
// Subtractではtrueを返す。
func (n *Subtract) IsLeftAssociative() bool {
	return true
}

// IsRightAssociative は右結合性かどうかを返す。
// Subtractではfalseを返す。
func (n *Subtract) IsRightAssociative() bool {
	return false
}

// NewSubtractは、減算のノードを返す
//
// left: 引かれる数のノード,
// tok: 対応するトークン,
// right: 引く数のノード。
func NewSubtract(left Node, tok token.Token, right Node) *Subtract {
	return &Subtract{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
