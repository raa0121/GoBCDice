package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 加算ロールのノード
type RandomNumber struct {
	InfixExpressionImpl
}

// RandomNumberがNodeを実装していることの確認
var _ Node = (*RandomNumber)(nil)

// Typeはノードの種類を返す
func (n *RandomNumber) Type() NodeType {
	return RANDOM_NUMBER_NODE
}

// IsCommutativeは可換演算子かどうかを返す
func (n *RandomNumber) IsCommutative() bool {
	return false
}

// Precedenceは演算子の優先順位を返す
func (n *RandomNumber) Precedence() OperatorPrecedenceType {
	return PREC_DOTS
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
