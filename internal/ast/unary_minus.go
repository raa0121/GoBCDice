package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 単項マイナスのノード
type UnaryMinus struct {
	PrefixExpressionImpl
}

// UnaryMinusがNodeを実装していることの確認
var _ Node = (*UnaryMinus)(nil)

// UnaryMinusがInfixExpressionを実装していることの確認
var _ PrefixExpression = (*UnaryMinus)(nil)

// Typeはノードの種類を返す
func (n *UnaryMinus) Type() NodeType {
	return UNARY_MINUS_NODE
}

// NewUnaryMinusは、新しい単項マイナスのノードを返す
//
// * tok: トークン
// * right: 右のノード
func NewUnaryMinus(tok token.Token, right Node) *UnaryMinus {
	return &UnaryMinus{
		PrefixExpressionImpl: *NewPrefixExpression(tok, right),
	}
}
