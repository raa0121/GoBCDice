package ast

import (
	"bytes"
	"strings"
)

// 個数振り足しロール列のノード。
type RRollList struct {
	// 個数振り足しロールのスライス。
	RRolls []*RRoll
	// 個数振り足しの閾値。
	Threshold Node
}

// RRollList がNodeを実装していることの確認。
var _ Node = (*RRollList)(nil)

// NewRRollList は新しい個数振り足しロール列のノードを返す。
//
// first: 最初の個数振り足しロール
// threshold: 個数振り足しの閾値。
func NewRRollList(first *RRoll, threshold Node) *RRollList {
	return &RRollList{
		RRolls:    []*RRoll{first},
		Threshold: threshold,
	}
}

// Type はノードの種類を返す。
func (n *RRollList) Type() NodeType {
	return R_ROLL_LIST_NODE
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

// IsPrimaryExpression は一次式かどうかを返す。
// RRollListではfalseを返す。
func (n *RRollList) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// RRollListではtrueを返す。
func (n *RRollList) IsVariable() bool {
	return true
}

// Append はリストにRRollを追加する。
func (n *RRollList) Append(r *RRoll) {
	n.RRolls = append(n.RRolls, r)
}
