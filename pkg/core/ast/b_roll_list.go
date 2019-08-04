package ast

import (
	"strings"
)

// バラバラロール列のノード。
type BRollList struct {
	NodeImpl

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
			tok: first.tok,
		},
		BRolls: []*BRoll{first},
	}
}

// Type はノードの種類を返す。
func (n *BRollList) Type() NodeType {
	return B_ROLL_LIST_NODE
}

// SExp はノードのS式を返す。
func (n *BRollList) SExp() string {
	bRollSExps := make([]string, 0, len(n.BRolls))
	for _, bRoll := range n.BRolls {
		bRollSExps = append(bRollSExps, bRoll.SExp())
	}

	return "(BRollList " + strings.Join(bRollSExps, " ") + ")"
}

// IsPrimaryExpression は一次式かどうかを返す。
// BRollListではfalseを返す。
func (n *BRollList) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// BRollListではtrueを返す。
func (n *BRollList) IsVariable() bool {
	return true
}

// Append はリストにBRollを追加する。
func (n *BRollList) Append(b *BRoll) {
	n.BRolls = append(n.BRolls, b)
}
