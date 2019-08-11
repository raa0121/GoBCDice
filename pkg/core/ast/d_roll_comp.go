package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算ロール式の成功判定のノード。
// コマンド。
type DRollComp struct {
	NodeImpl
	CommandImpl
}

// DRollComp がNodeを実装していることの確認。
var _ Node = (*DRollComp)(nil)

// DRollComp がCommandを実装していることの確認。
var _ Command = (*DRollComp)(nil)

// NewDRollComp は新しい加算ロール式の成功判定のノードを返す。
//
// tok: 対応するトークン,
// expression: 式。
func NewDRollComp(tok token.Token, expression Node) *DRollComp {
	return &DRollComp{
		CommandImpl: CommandImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			expr: expression,
		},
	}
}

// NewDRollComp2 は新しい加算ロール式の成功判定のノードを返す。
//
// expression: 式。
func NewDRollComp2(expression Node) *DRollComp {
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
