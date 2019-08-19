package ast

import (
	"fmt"
)

// 計算コマンドのノード。
// コマンド。
type Calc struct {
	CommandImpl
}

// Calc がNodeを実装していることの確認。
var _ Node = (*Calc)(nil)

// Calc がCommandを実装していることの確認。
var _ Command = (*Calc)(nil)

// NewCalc は新しい計算コマンドを返す。
//
// expression: 式。
func NewCalc(expression Node) *Calc {
	return &Calc{
		CommandImpl: CommandImpl{
			expr: expression,
		},
	}
}

// Type はノードの種類を返す。
func (n *Calc) Type() NodeType {
	return CALC_NODE
}

// SExp はノードのS式を返す。
func (n *Calc) SExp() string {
	return fmt.Sprintf("(Calc %s)", n.Expression().SExp())
}
