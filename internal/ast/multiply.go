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

// Typeはノードの種類を返す
func (n *Multiply) Type() NodeType {
	return MULTIPLY_NODE
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
