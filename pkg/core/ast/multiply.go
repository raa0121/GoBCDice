package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 乗算のノード。
// 中置式。
type Multiply struct {
	InfixExpressionImpl
}

// Multiply がNodeを実装していることの確認。
var _ Node = (*Multiply)(nil)

// Multiply がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*Multiply)(nil)

// Type はノードの種類を返す。
func (n *Multiply) Type() NodeType {
	return MULTIPLY_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *Multiply) Precedence() OperatorPrecedenceType {
	return PREC_MULTITIVE
}

// IsLeftAssociative は左結合性かどうかを返す。
// Multiplyではtrueを返す。
func (n *Multiply) IsLeftAssociative() bool {
	return true
}

// IsRightAssociative は右結合性かどうかを返す。
// Multiplyではtrueを返す。
func (n *Multiply) IsRightAssociative() bool {
	return true
}

// NewMultiply は新しい乗算のノードを返す。
//
// multiplicand: 被乗数のノード,
// tok: 対応するトークン,
// multiplier: 乗数のノード。
func NewMultiply(multiplicand Node, tok token.Token, multiplier Node) *Multiply {
	return &Multiply{
		InfixExpressionImpl: *NewInfixExpression(multiplicand, tok, multiplier),
	}
}

// NewMultiply2 は新しい乗算のノードを返す。
//
// multiplicand: 被乗数のノード,
// multiplier: 乗数のノード。
func NewMultiply2(multiplicand Node, multiplier Node) *Multiply {
	return &Multiply{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            multiplicand,
			operator:        "*",
			operatorForSExp: "*",
			right:           multiplier,
		},
	}
}
