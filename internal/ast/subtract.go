package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 減算のノード
type Subtract struct {
	InfixExpressionImpl
}

// SubtractがNodeを実装していることの確認
var _ Node = (*Subtract)(nil)

// Typeはノードの種類を返す
func (n *Subtract) Type() NodeType {
	return SUBTRACT_NODE
}

// IsCommutativeは可換演算子かどうかを返す
func (n *Subtract) IsCommutative() bool {
	return false
}

// Precedenceは演算子の優先順位を返す
func (n *Subtract) Precedence() OperatorPrecedenceType {
	return PREC_ADDITIVE
}

// NewSubtractは、減算のノードを返す
//
// * left: 引かれる数のノード
// * tok: 対応するトークン
// * right: 引く数のノード
func NewSubtract(left Node, tok token.Token, right Node) *Subtract {
	return &Subtract{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
