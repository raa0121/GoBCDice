package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// 計算コマンドのノード
type Calc struct {
	CommandImpl
}

// CalcがNodeを実装していることの確認
var _ Node = (*Calc)(nil)

// CalcがCommandを実装していることの確認
var _ Command = (*Calc)(nil)

// NewCalcは新しい計算コマンドを返す
//
// * tok: トークン
// * expression: 式
func NewCalc(tok token.Token, expression Node) *Calc {
	return &Calc{
		CommandImpl: CommandImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			expr: expression,
		},
	}
}

// Typeはノードの種類を返す
func (n *Calc) Type() NodeType {
	return CALC_NODE
}

// SExpはノードのS式を返す
func (n *Calc) SExp() string {
	return fmt.Sprintf("(Calc %s)", n.Expression().SExp())
}
