package ast

import (
	"fmt"
)

// 加算ロール式の成功判定のノード。
// コマンド。
type DRollComp struct {
	CommandImpl
}

// DRollComp がNodeを実装していることの確認。
var _ Node = (*DRollComp)(nil)

// DRollComp がCommandを実装していることの確認。
var _ Command = (*DRollComp)(nil)

// NewDRollComp は新しい加算ロール式の成功判定のノードを返す。
//
// expression: 式。
func NewDRollComp(expression Node) *DRollComp {
	return &DRollComp{
		CommandImpl: CommandImpl{
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *DRollComp) Type() NodeType {
	return D_ROLL_COMP_NODE
}

// SExp はノードのS式を返す。
func (n *DRollComp) SExp() string {
	return fmt.Sprintf("(DRollComp %s)", n.Expression().SExp())
}
