package ast

import (
	"bytes"
)

// Command はトップレベルにあるコマンドを表す。
type Command struct {
	NodeImpl
	NonNilNode

	// Expression はコマンドの引数である式のノード
	Expression Node
}

// Command がNodeを実装していることの確認。
var _ Node = (*Command)(nil)

// IsVariable は可変ノードかどうかを返す。
//
// コマンドでは、引数の式が可変ノードならばtrueを返す。
// 引数の式が可変ノードでない場合はfalseを返す。
func (n *Command) IsVariable() bool {
	return n.Expression.IsVariable()
}

// SExp はノードのS式を返す。
func (n *Command) SExp() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(n.Type().String())
	buf.WriteString(" ")
	buf.WriteString(n.Expression.SExp())
	buf.WriteString(")")

	return buf.String()
}

// newCommand は新しいコマンドのノードを返す。
//
// nodeType: ノードの種類,
// expression: 式。
func newCommand(nodeType NodeType, expression Node) *Command {
	return &Command{
		NodeImpl: NodeImpl{
			nodeType:            nodeType,
			isPrimaryExpression: false,
		},

		Expression: expression,
	}
}

// NewCalc は新しい計算コマンドを返す。
//
// expression: 式。
func NewCalc(expression Node) *Command {
	return newCommand(CALC_NODE, expression)
}

// NewDRollExpr は新しい加算ロール式のノードを返す。
//
// expression: 式。
func NewDRollExpr(expression Node) *Command {
	return newCommand(D_ROLL_EXPR_NODE, expression)
}

// NewDRollComp は新しい加算ロール式の成功判定のノードを返す。
//
// expression: 式。
func NewDRollComp(expression Node) *Command {
	return newCommand(D_ROLL_COMP_NODE, expression)
}

// NewBRollComp は新しいバラバラロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewBRollComp(expression Node) *Command {
	return newCommand(B_ROLL_COMP_NODE, expression)
}

// NewRRollComp は新しい個数振り足しロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewRRollComp(expression Node) *Command {
	return newCommand(R_ROLL_COMP_NODE, expression)
}

// NewURollComp は新しい上方無限ロールの成功数カウントのノードを返す。
//
// expression: 式。
func NewURollComp(expression Node) *Command {
	return newCommand(U_ROLL_COMP_NODE, expression)
}
