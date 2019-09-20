package ast

import (
	"fmt"
)

// 上方無限ロールの成功数カウントのノード。
// コマンド。
type URollComp struct {
	CommandImpl
}

// URollComp がNodeを実装していることの確認。
var _ Node = (*URollComp)(nil)

// URollComp がCommandを実装していることの確認。
var _ Command = (*URollComp)(nil)

// NewURollComp は新しい上方無限ロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewURollComp(expression Node) *URollComp {
	return &URollComp{
		CommandImpl: CommandImpl{
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *URollComp) Type() NodeType {
	return U_ROLL_COMP_NODE
}

// SExp はノードのS式を返す。
func (n *URollComp) SExp() string {
	return fmt.Sprintf("(URollComp %s)", n.Expression().SExp())
}
