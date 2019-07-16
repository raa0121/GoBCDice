package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// ランダム数値取り出しのノード。
// 一次式、可変ノード、中置式。
type RandomNumber struct {
	InfixExpressionImpl
}

// RandomNumber がNodeを実装していることの確認。
var _ Node = (*RandomNumber)(nil)

// RandomNumber がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*RandomNumber)(nil)

// Type はノードの種類を返す。
func (n *RandomNumber) Type() NodeType {
	return RANDOM_NUMBER_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *RandomNumber) Precedence() OperatorPrecedenceType {
	return PREC_DOTS
}

// IsLeftAssociative は左結合性かどうかを返す。
// RandomNumberではfalseを返す。
func (n *RandomNumber) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// RandomNumberではfalseを返す。
func (n *RandomNumber) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpression は一次式かどうかを返す。
// RandomNumberではtrueを返す。
func (n *RandomNumber) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// RandomNumberではtrueを返す。
func (n *RandomNumber) IsVariable() bool {
	return true
}

// NewRand はランダム数値取り出しのノードを返す。
//
// min: 最小値のノード,
// tok: 対応するトークン,
// max: 最大値のノード。
func NewRandomNumber(min Node, tok token.Token, max Node) *RandomNumber {
	return &RandomNumber{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            min,
			operator:        "...",
			operatorForSExp: "Rand",
			right:           max,
		},
	}
}