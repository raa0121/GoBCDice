package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算ロールのノード
type RandomNumber struct {
	InfixExpressionImpl
}

// RandomNumberがNodeを実装していることの確認
var _ Node = (*RandomNumber)(nil)

// RandomNumberがInfixExpressionを実装していることの確認
var _ InfixExpression = (*RandomNumber)(nil)

// Typeはノードの種類を返す
func (n *RandomNumber) Type() NodeType {
	return RANDOM_NUMBER_NODE
}

// Precedenceは演算子の優先順位を返す
func (n *RandomNumber) Precedence() OperatorPrecedenceType {
	return PREC_DOTS
}

// IsLeftAssociativeは左結合性かどうかを返す
func (n *RandomNumber) IsLeftAssociative() bool {
	return false
}

// IsRightAssociativeは右結合性かどうかを返す
func (n *RandomNumber) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpressionは一次式かどうかを返す
func (n *RandomNumber) IsPrimaryExpression() bool {
	return true
}

// IsVariableは可変ノードかどうかを返す。
func (n *RandomNumber) IsVariable() bool {
	return true
}

// NewRandはランダム数値取り出しのノードを返す
//
// * min: 最小値のノード
// * tok: 対応するトークン
// * max: 最大値のノード
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
