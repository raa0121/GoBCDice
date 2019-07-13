package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算ロールのノード
type DRoll struct {
	InfixExpressionImpl
}

// DRollがNodeを実装していることの確認
var _ Node = (*DRoll)(nil)

// DRollがInfixExpressionを実装していることの確認
var _ InfixExpression = (*DRoll)(nil)

// Typeはノードの種類を返す
func (n *DRoll) Type() NodeType {
	return D_ROLL_NODE
}

// Precedenceは演算子の優先順位を返す
func (n *DRoll) Precedence() OperatorPrecedenceType {
	return PREC_D_ROLL
}

// IsLeftAssociativeは左結合性かどうかを返す
func (n *DRoll) IsLeftAssociative() bool {
	return false
}

// IsRightAssociativeは右結合性かどうかを返す
func (n *DRoll) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpressionは一次式かどうかを返す
func (n *DRoll) IsPrimaryExpression() bool {
	return true
}

// IsVariableは可変ノードかどうかを返す。
func (n *DRoll) IsVariable() bool {
	return true
}

// NewDRollは加算ロールのノードを返す
//
// * num: 振るダイスの数のノード
// * tok: 対応するトークン
// * sides: ダイスの面数のノード
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
