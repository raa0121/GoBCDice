package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算ロールのノード。
// 一次式、可変ノード、中置式。
type DRoll struct {
	InfixExpressionImpl
}

// DRoll がNodeを実装していることの確認。
var _ Node = (*DRoll)(nil)

// DRoll がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*DRoll)(nil)

// Type はノードの種類を返す。
func (n *DRoll) Type() NodeType {
	return D_ROLL_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *DRoll) Precedence() OperatorPrecedenceType {
	return PREC_ROLL
}

// IsLeftAssociative は左結合性かどうかを返す。
// DRollではfalseを返す。
func (n *DRoll) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// DRollではfalseを返す。
func (n *DRoll) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpression は一次式かどうかを返す。
// DRollではtrueを返す。
func (n *DRoll) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// DRollではtrueを返す。
func (n *DRoll) IsVariable() bool {
	return true
}

// NewDRoll は加算ロールのノードを返す。
//
// num: 振るダイスの数のノード,
// tok: 対応するトークン,
// sides: ダイスの面数のノード。
func NewDRoll(num Node, tok token.Token, sides Node) *DRoll {
	return &DRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            num,
			operator:        "D",
			operatorForSExp: "DRoll",
			right:           sides,
		},
	}
}

// NewDRoll2 は加算ロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewDRoll2(num Node, sides Node) *DRoll {
	return &DRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            num,
			operator:        "D",
			operatorForSExp: "DRoll",
			right:           sides,
		},
	}
}
