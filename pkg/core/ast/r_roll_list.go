package ast

import (
	"bytes"
	"strings"
)

// 個数振り足しロール列のノード。
type RRollList struct {
	NodeImpl
	NonNilNode
	VariableNode

	// 個数振り足しロールのスライス。
	RRolls []*VariableInfixExpression
	// 個数振り足しの閾値。
	Threshold Node
}

// RRollList がNodeを実装していることの確認。
var _ Node = (*RRollList)(nil)

// NewRRollList は新しい個数振り足しロール列のノードを返す。
//
// first: 最初の個数振り足しロール
// threshold: 個数振り足しの閾値。
func NewRRollList(first *VariableInfixExpression, threshold Node) *RRollList {
	return &RRollList{
		NodeImpl: NodeImpl{
			nodeType:            R_ROLL_LIST_NODE,
			isPrimaryExpression: false,
		},

		RRolls:    []*VariableInfixExpression{first},
		Threshold: threshold,
	}
}

// SExp はノードのS式を返す。
func (n *RRollList) SExp() string {
	var out bytes.Buffer

	rRollSExps := make([]string, 0, len(n.RRolls))
	for _, rRoll := range n.RRolls {
		rRollSExps = append(rRollSExps, rRoll.SExp())
	}

	out.WriteString("(RRollList ")
	out.WriteString(n.Threshold.SExp())
	out.WriteString(" ")
	out.WriteString(strings.Join(rRollSExps, " "))
	out.WriteString(")")

	return out.String()
}

// Append はリストにRRollを追加する。
func (n *RRollList) Append(r *VariableInfixExpression) {
	n.RRolls = append(n.RRolls, r)
}
