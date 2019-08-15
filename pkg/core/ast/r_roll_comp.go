package ast

import (
	"fmt"
)

// 個数振り足しロールの成功数カウントのノード。
// コマンド。
type RRollComp struct {
	CommandImpl
}

// RRollComp がNodeを実装していることの確認。
var _ Node = (*RRollComp)(nil)

// RRollComp がCommandを実装していることの確認。
var _ Command = (*RRollComp)(nil)

// NewRRollComp は新しい個数振り足しロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewRRollComp(expression Node) *RRollComp {
	return &RRollComp{
		CommandImpl: CommandImpl{
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *RRollComp) Type() NodeType {
	return R_ROLL_COMP_NODE
}

// SExp はノードのS式を返す。
func (n *RRollComp) SExp() string {
	return fmt.Sprintf("(RRollComp %s)", n.Expression().SExp())
}
