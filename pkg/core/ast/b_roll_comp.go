package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// バラバラロールの成功数カウントのノード。
// コマンド。
type BRollComp struct {
	NodeImpl
	CommandImpl
}

// BRollComp がNodeを実装していることの確認。
var _ Node = (*BRollComp)(nil)

// BRollComp がCommandを実装していることの確認。
var _ Command = (*BRollComp)(nil)

// NewDRollComp は新しいバラバラロールの成功数カウントのノードを返す。
//
// tok: 対応するトークン,
// expression: 式。
func NewBRollComp(tok token.Token, expression Node) *BRollComp {
	return &BRollComp{
		CommandImpl: CommandImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
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
