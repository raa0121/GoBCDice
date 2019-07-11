package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 加算のノード
type Add struct {
	InfixExpressionImpl
}

// AddがNodeを実装していることの確認
var _ Node = (*Add)(nil)

// AddがInfixExpressionを実装していることの確認
var _ InfixExpression = (*Add)(nil)

// Typeはノードの種類を返す
func (n *Add) Type() NodeType {
	return ADD_NODE
}

// IsCommutativeは可換演算子かどうかを返す
func (n *Add) IsCommutative() bool {
	return true
}

// Precedenceは演算子の優先順位を返す
func (n *Add) Precedence() OperatorPrecedenceType {
	return PREC_ADDITIVE
}

// IsLeftAssociativeは左結合性かどうかを返す
func (n *Add) IsLeftAssociative() bool {
	return true
}

// IsRightAssociativeは右結合性かどうかを返す
func (n *Add) IsRightAssociative() bool {
	return true
}

// NewAddは、加算のノードを返す
//
// * left: 加えられる数のノード
// * tok: 対応するトークン
// * right: 加える数のノード
func NewAdd(left Node, tok token.Token, right Node) *Add {
	return &Add{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
