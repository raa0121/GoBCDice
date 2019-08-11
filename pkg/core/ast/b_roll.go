package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// バラバラロールのノード。
// 一次式、可変ノード、中置式。
type BRoll struct {
	InfixExpressionImpl
}

// BRoll がNodeを実装していることの確認。
var _ Node = (*BRoll)(nil)

// BRoll がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*BRoll)(nil)

// Type はノードの種類を返す。
func (n *BRoll) Type() NodeType {
	return B_ROLL_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *BRoll) Precedence() OperatorPrecedenceType {
	return PREC_ROLL
}

// IsLeftAssociative は左結合性かどうかを返す。
// BRollではfalseを返す。
func (n *BRoll) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// BRollではfalseを返す。
func (n *BRoll) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpression は一次式かどうかを返す。
// BRollではtrueを返す。
func (n *BRoll) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// BRollではtrueを返す。
func (n *BRoll) IsVariable() bool {
	return true
}

// NewBRoll はバラバラロールのノードを返す。
//
// num: 振るダイスの数のノード,
// tok: 対応するトークン,
// sides: ダイスの面数のノード。
func NewBRoll(num Node, tok token.Token, sides Node) *BRoll {
	return &BRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            num,
			operator:        "B",
			operatorForSExp: "BRoll",
			right:           sides,
		},
	}
}

// NewBRoll2 はバラバラロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewBRoll2(num Node, sides Node) *BRoll {
	return &BRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            num,
			operator:        "B",
			operatorForSExp: "BRoll",
			right:           sides,
		},
	}
}
