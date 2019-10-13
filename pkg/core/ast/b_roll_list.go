package ast

import (
	"bytes"
	"strings"
)

// バラバラロール列のノード。
type BRollList struct {
	NodeImpl
	NonNilNode
	VariableNode

	// バラバラロールのスライス。
	// 2b6+4d10のように連続してダイスロールを行えるように、複数のバラバラロールを格納する。
	BRolls []*BRoll
}

// BRollList がNodeを実装していることの確認。
var _ Node = (*BRollList)(nil)

// NewBRollList は新しいバラバラロール列のノードを返す。
//
// first: 最初のバラバラロール
func NewBRollList(first *BRoll) *BRollList {
	return &BRollList{
		NodeImpl: NodeImpl{
			nodeType:            B_ROLL_LIST_NODE,
			isPrimaryExpression: false,
		},

		BRolls: []*BRoll{first},
	}
}

// SExp はノードのS式を返す。
func (n *BRollList) SExp() string {
	var out bytes.Buffer

	bRollSExps := make([]string, 0, len(n.BRolls))
	for _, bRoll := range n.BRolls {
		bRollSExps = append(bRollSExps, bRoll.SExp())
	}

	out.WriteString("(BRollList ")
	out.WriteString(strings.Join(bRollSExps, " "))
	out.WriteString(")")

	return out.String()
}

// Append はリストにBRollを追加する。
func (n *BRollList) Append(b *BRoll) {
	n.BRolls = append(n.BRolls, b)
}
