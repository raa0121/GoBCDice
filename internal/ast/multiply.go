package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 乗算のノード
type Multiply struct {
	InfixExpressionImpl
}

// MultiplyがNodeを実装していることの確認
var _ Node = (*Multiply)(nil)

// MultiplyがInfixExpressionを実装していることの確認
var _ InfixExpression = (*Multiply)(nil)

// Typeはノードの種類を返す
func (n *Multiply) Type() NodeType {
	return MULTIPLY_NODE
}

// IsCommutativeは可換演算子かどうかを返す
func (n *Multiply) IsCommutative() bool {
	return true
}

// Precedenceは演算子の優先順位を返す
func (n *Multiply) Precedence() OperatorPrecedenceType {
	return PREC_MULTITIVE
}

// IsLeftAssociativeは左結合性かどうかを返す
func (n *Multiply) IsLeftAssociative() bool {
	return true
}

// IsRightAssociativeは右結合性かどうかを返す
func (n *Multiply) IsRightAssociative() bool {
	return true
}

// NewMultiplyは、乗算のノードを返す
//
// * multiplicand: 被乗数のノード
// * tok: 対応するトークン
// * multiplier: 乗数のノード
func NewMultiply(multiplicand Node, tok token.Token, multiplier Node) *Multiply {
	return &Multiply{
		InfixExpressionImpl: *NewInfixExpression(multiplicand, tok, multiplier),
	}
}
