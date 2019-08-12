package ast

import (
	"bytes"
)

// ランダム選択のノード。
type Choice struct {
	NodeImpl

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
			tok: first.tok,
		},
		Items: []*String{first},
	}
}

// Type はノードの種類を返す。
func (n *Choice) Type() NodeType {
	return CHOICE_NODE
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

// IsPrimaryExpression は一次式かどうかを返す。
// Choiceではfalseを返す。
func (n *Choice) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// Choiceではtrueを返す。
func (n *Choice) IsVariable() bool {
	return true
}

// Append はリストに文字列ノードを追加する。
func (n *Choice) Append(s *String) {
	n.Items = append(n.Items, s)
}
