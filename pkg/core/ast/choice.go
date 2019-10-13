package ast

import (
	"bytes"
)

// ランダム選択のノード。
type Choice struct {
	NodeImpl
	NonNilNode
	VariableNode

	// 選択肢のスライス。
	Items []*String
}

// Choice がNodeを実装していることの確認。
var _ Node = (*Choice)(nil)

// NewChoice は新しいランダム選択ノードを返す。
//
// first: 最初の選択肢。
func NewChoice(first *String) *Choice {
	return &Choice{
		NodeImpl: NodeImpl{
			nodeType:            CHOICE_NODE,
			isPrimaryExpression: false,
		},

		Items: []*String{first},
	}
}

// SExp はノードのS式を返す。
func (n *Choice) SExp() string {
	var out bytes.Buffer

	out.WriteString("(Choice")

	for _, i := range n.Items {
		out.WriteString(" ")
		out.WriteString(i.SExp())
	}

	out.WriteString(")")

	return out.String()
}

// Append はリストに文字列ノードを追加する。
func (n *Choice) Append(s *String) {
	n.Items = append(n.Items, s)
}
