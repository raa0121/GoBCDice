package ast

import (
	"bytes"
	"strings"
)

// URollList は上方無限ロール列のノード。
type URollList struct {
	NonNilNode

	// 上方無限ロールのスライス。
	URolls []*RRoll
	// 上方無限の閾値。
	Threshold Node
}

// URollList がNodeを実装していることの確認。
var _ Node = (*URollList)(nil)

// NewURollList は新しい上方無限ロール列のノードを返す。
//
// first: 最初の上方無限ロール
// threshold: 上方無限の閾値。
func NewURollList(first *RRoll, threshold Node) *URollList {
	return &URollList{
		URolls:    []*RRoll{first},
		Threshold: threshold,
	}
}

// Type はノードの種類を返す。
func (n *URollList) Type() NodeType {
	return U_ROLL_LIST_NODE
}

// SExp はノードのS式を返す。
func (n *URollList) SExp() string {
	var out bytes.Buffer

	uRollSExps := make([]string, 0, len(n.URolls))
	for _, rRoll := range n.URolls {
		uRollSExps = append(uRollSExps, rRoll.SExp())
	}

	out.WriteString("(URollList ")
	out.WriteString(n.Threshold.SExp())
	out.WriteString(" ")
	out.WriteString(strings.Join(uRollSExps, " "))
	out.WriteString(")")

	return out.String()
}

// IsPrimaryExpression は一次式かどうかを返す。
// URollListではfalseを返す。
func (n *URollList) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// URollListではtrueを返す。
func (n *URollList) IsVariable() bool {
	return true
}

// Append はリストにURollを追加する。
func (n *URollList) Append(r *RRoll) {
	n.URolls = append(n.URolls, r)
}
