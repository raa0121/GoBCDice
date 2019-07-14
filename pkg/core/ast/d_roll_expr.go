package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 加算ロール式のノード。
// コマンド。
type DRollExpr struct {
	NodeImpl
	CommandImpl
}

// DRollExpr がNodeを実装していることの確認。
var _ Node = (*DRollExpr)(nil)

// DRollExpr がCommandを実装していることの確認。
var _ Command = (*DRollExpr)(nil)

// NewDRollExpr は新しい加算ロール式のノードを返す。
//
// tok: 対応するトークン,
// expression: 式。
func NewDRollExpr(tok token.Token, expression Node) *DRollExpr {
	return &DRollExpr{
		CommandImpl: CommandImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *DRollExpr) Type() NodeType {
	return D_ROLL_EXPR_NODE
}

// SExp はノードのS式を返す。
func (n *DRollExpr) SExp() string {
	return fmt.Sprintf("(DRollExpr %s)", n.Expression().SExp())
}
