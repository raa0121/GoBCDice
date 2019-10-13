package ast

import (
	"bytes"
	"github.com/raa0121/GoBCDice/pkg/core/util"
)

// URollExpr は上方無限ロール式のノードを表す。
type URollExpr struct {
	NodeImpl
	NonNilNode
	VariableNode

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
		NodeImpl: NodeImpl{
			nodeType:            U_ROLL_EXPR_NODE,
			isPrimaryExpression: false,
		},

		URollList: uRollList,
		Bonus:     bonus,
	}
}

// SExp はノードのS式を返す。
func (n *URollExpr) SExp() string {
	var out bytes.Buffer

	out.WriteString("(URollExpr ")

	if n.Bonus == nil {
		out.WriteString(n.URollList.SExp())
	} else {
		// cloneしてinterface{}型に変わる
		clonedBonus := util.Clone(n.Bonus).(BasicInfixExpression)
		clonedBonus.SetLeft(n.URollList)

		out.WriteString(clonedBonus.SExp())
	}

	out.WriteString(")")

	return out.String()
}
