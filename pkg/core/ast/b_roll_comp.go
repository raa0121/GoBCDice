package ast

import (
	"fmt"
)

// バラバラロールの成功数カウントのノード。
// コマンド。
type BRollComp struct {
	CommandImpl
}

// BRollComp がNodeを実装していることの確認。
var _ Node = (*BRollComp)(nil)

// BRollComp がCommandを実装していることの確認。
var _ Command = (*BRollComp)(nil)

// NewBRollComp は新しいバラバラロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewBRollComp(expression Node) *BRollComp {
	return &BRollComp{
		CommandImpl: CommandImpl{
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *BRollComp) Type() NodeType {
	return B_ROLL_COMP_NODE
}

// SExp はノードのS式を返す。
func (n *BRollComp) SExp() string {
	return fmt.Sprintf("(BRollComp %s)", n.Expression().SExp())
}
