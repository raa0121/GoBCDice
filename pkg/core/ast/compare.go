package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 比較のノード。
// 中置式。
type Compare struct {
	InfixExpressionImpl
}

// Compare がNodeを実装していることの確認。
var _ Node = (*Compare)(nil)

// Compare がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*Compare)(nil)

// Type はノードの種類を返す。
func (n *Compare) Type() NodeType {
	return COMPARE_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *Compare) Precedence() OperatorPrecedenceType {
	return PREC_COMPARE
}

// IsLeftAssociative は左結合性かどうかを返す。
// Compareではfalseを返す。
func (n *Compare) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// Compareではfalseを返す。
func (n *Compare) IsRightAssociative() bool {
	return false
}

// NewCompare は新しい比較のノードを返す。
//
// left: 左辺のノード,
// tok: 対応するトークン,
// right: 右辺のノード。
func NewCompare(left Node, tok token.Token, right Node) *Compare {
	return &Compare{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
