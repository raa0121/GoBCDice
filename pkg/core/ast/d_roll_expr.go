package ast

import (
	"fmt"
)

// 加算ロール式のノード。
// コマンド。
type DRollExpr struct {
	CommandImpl
}

// DRollExpr がNodeを実装していることの確認。
var _ Node = (*DRollExpr)(nil)

// DRollExpr がCommandを実装していることの確認。
var _ Command = (*DRollExpr)(nil)

// NewDRollExpr は新しい加算ロール式のノードを返す。
//
// expression: 式。
func NewDRollExpr(expression Node) *DRollExpr {
	return &DRollExpr{
		CommandImpl: CommandImpl{
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
