package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// 中置演算子のインターフェース
type InfixExpression interface {
	Node
	IsInfixExpression() bool
	// Leftは左のノードを返す
	Left() Node
	// Operatorは演算子を返す
	Operator() string
	// OperatorForSExpはS式で表示する演算子を返す
	OperatorForSExp() string
	// Rightは右のノードを返す
	Right() Node
}

// 中置演算子のノードを表す構造体
type InfixExpressionImpl struct {
	NodeImpl
	// 左のノード
	left Node
	// 演算子
	operator string
	// S式で表示する演算子
	operatorForSExp string
	// 右のノード
	right Node
}

// InfixExpressionがNodeを実装していることの確認
var _ Node = (*InfixExpressionImpl)(nil)

// Typeはノードの種類を返す
func (n *InfixExpressionImpl) Type() NodeType {
	return INFIX_EXPRESSION_NODE
}

func (n *InfixExpressionImpl) IsInfixExpression() bool {
	return true
}

// Leftは左のノードを返す
func (n *InfixExpressionImpl) Left() Node {
	return n.left
}

// Operatorは演算子を返す
func (n *InfixExpressionImpl) Operator() string {
	return n.operator
}

// OperatorForSExpはS式で表示する演算子を返す
func (n *InfixExpressionImpl) OperatorForSExp() string {
	return n.operatorForSExp
}

// Rightは右のノードを返す
func (n *InfixExpressionImpl) Right() Node {
	return n.right
}

// SExpはノードのS式を返す
func (n *InfixExpressionImpl) SExp() string {
	return fmt.Sprintf("(%s %s %s)",
		n.OperatorForSExp(), n.Left().SExp(), n.Right().SExp())
}

// InfixNotationは中置記法表記を返す
func (n *InfixExpressionImpl) InfixNotation() string {
	return n.left.InfixNotation() + n.Operator() + n.right.InfixNotation()
}

// NewInfixExpressionは中置演算子のノードを返す。
// 評価時とS式とで演算子の表示を変更しなくてもよい場合に使う。
//
// * left: 左のノード
// * tok: 対応するトークン
// * right: 右のノード
func NewInfixExpression(left Node, tok token.Token, right Node) *InfixExpressionImpl {
	return &InfixExpressionImpl{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		left:            left,
		operator:        tok.Literal,
		operatorForSExp: tok.Literal,
		right:           right,
	}
}
