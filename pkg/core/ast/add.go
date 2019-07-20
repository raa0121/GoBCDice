package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算のノード。
// 中置式。
type Add struct {
	InfixExpressionImpl
}

// Add がNodeを実装していることの確認。
var _ Node = (*Add)(nil)

// Add がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*Add)(nil)

// Type はノードの種類を返す。
func (n *Add) Type() NodeType {
	return ADD_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *Add) Precedence() OperatorPrecedenceType {
	return PREC_ADDITIVE
}

// IsLeftAssociative は左結合性かどうかを返す。
// Addではtrueを返す。
func (n *Add) IsLeftAssociative() bool {
	return true
}

// IsRightAssociative は右結合性かどうかを返す。
// Addではtrueを返す。
func (n *Add) IsRightAssociative() bool {
	return true
}

// NewAdd は新しい加算のノードを返す。
//
// left: 加えられる数のノード,
// tok: 対応するトークン,
// right: 加える数のノード。
func NewAdd(left Node, tok token.Token, right Node) *Add {
	return &Add{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
