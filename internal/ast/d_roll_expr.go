package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// 加算ロール式のノード
type DRollExpr struct {
	NodeImpl
	CommandImpl
}

// DRollExprがNodeを実装していることの確認
var _ Node = (*DRollExpr)(nil)

// DRollExprがCommandを実装していることの確認
var _ Command = (*DRollExpr)(nil)

// NewDRollExprは新しい計算コマンドを返す
//
// * tok: トークン
// * expression: 式
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

func (n *DRollExpr) Type() NodeType {
	return D_ROLL_EXPR_NODE
}

func (n *DRollExpr) SExp() string {
	return fmt.Sprintf("(DRollExpr %s)", n.Expression().SExp())
}

func (n *DRollExpr) InfixNotation() string {
	return n.Expression().InfixNotation()
}
