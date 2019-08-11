package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 単項マイナスのノード。
// 前置式。
type UnaryMinus struct {
	PrefixExpressionImpl
}

// UnaryMinus がNodeを実装していることの確認。
var _ Node = (*UnaryMinus)(nil)

// UnaryMinus がPrefixExpressionを実装していることの確認。
var _ PrefixExpression = (*UnaryMinus)(nil)

// Type はノードの種類を返す。
func (n *UnaryMinus) Type() NodeType {
	return UNARY_MINUS_NODE
}

// NewUnaryMinus は新しい単項マイナスのノードを返す。
//
// tok: 対応するトークン,
// right: 右のノード。
func NewUnaryMinus(tok token.Token, right Node) *UnaryMinus {
	return &UnaryMinus{
		PrefixExpressionImpl: *NewPrefixExpression(tok, right),
	}
}

// NewUnaryMinus2 は新しい単項マイナスのノードを返す。
//
// right: 右のノード。
func NewUnaryMinus2(right Node) *UnaryMinus {
	return &UnaryMinus{
		PrefixExpressionImpl: PrefixExpressionImpl{
			operator:        "-",
			operatorForSExp: "-",
			right:           right,
		},
	}
}
