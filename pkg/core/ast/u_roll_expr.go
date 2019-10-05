package ast

import (
	"bytes"
	"github.com/raa0121/GoBCDice/pkg/core/util"
)

// URollExpr は上方無限ロール式のノードを表す。
type URollExpr struct {
	NonNilNode

	// 上方無限ロールのリスト。
	URollList *RRollList
	// ボーナスのノード。
	Bonus InfixExpression
}

// DRollExpr がNodeを実装していることの確認。
var _ Node = (*URollExpr)(nil)

// NewURollExpr は新しい上方無限ロール式のノードを返す。
func NewURollExpr(uRollList *RRollList, bonus InfixExpression) *URollExpr {
	return &URollExpr{
		URollList: uRollList,
		Bonus:     bonus,
	}
}

// Type はノードの種類を返す。
func (n *URollExpr) Type() NodeType {
	return U_ROLL_EXPR_NODE
}

// SExp はノードのS式を返す。
func (n *URollExpr) SExp() string {
	var out bytes.Buffer

	out.WriteString("(URollExpr ")

	if n.Bonus == nil {
		out.WriteString(n.URollList.SExp())
	} else {
		// cloneしてinterface{}型に変わる
		clonedBonus := util.Clone(n.Bonus)

		var bonusInfixExpession InfixExpression
		switch b := clonedBonus.(type) {
		case Add:
			bonusInfixExpession = &b
		case Subtract:
			bonusInfixExpession = &b
		}

		bonusInfixExpession.SetLeft(n.URollList)

		out.WriteString(bonusInfixExpession.SExp())
	}

	out.WriteString(")")

	return out.String()
}

// IsPrimaryExpression は一次式かどうかを返す。
// URollExprではfalseを返す。
func (n *URollExpr) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// URollExprではtrueを返す。
func (n *URollExpr) IsVariable() bool {
	return true
}
