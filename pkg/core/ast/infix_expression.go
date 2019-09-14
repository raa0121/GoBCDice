package ast

import (
	"fmt"
)

// 中置式のインターフェース。
type InfixExpression interface {
	Node

	// IsInfixExpression は中置式であるかを返す（ダミー関数）。
	IsInfixExpression() bool
	// Left は左のノードを返す。
	Left() Node
	// SetLeft は左のノードを設定する。
	SetLeft(l Node)
	// Operator は演算子を返す。
	Operator() string
	// OperatorForSExp はS式で表示する演算子を返す。
	OperatorForSExp() string
	// Right は右のノードを返す。
	Right() Node
	// SetRight は右のノードを設定する。
	SetRight(r Node)
	// Precedence は演算子の優先順位を返す。
	Precedence() OperatorPrecedenceType
	// IsLeftAssociative は左結合性かどうかを返す。
	IsLeftAssociative() bool
	// IsRightAssociative は右結合性かどうかを返す。
	IsRightAssociative() bool
}

// 中置式のノードが共通して持つ要素。
type InfixExpressionImpl struct {
	NonNilNode

	// 左のノード
	left Node
	// 演算子
	operator string
	// S式で表示する演算子
	operatorForSExp string
	// 右のノード
	right Node
}

// InfixExpressionImpl がNodeを実装していることの確認。
var _ Node = (*InfixExpressionImpl)(nil)

// Type はノードの種類を返す。
func (n *InfixExpressionImpl) Type() NodeType {
	return INFIX_EXPRESSION_NODE
}

// IsInfixExpression は中置式であるかを返す（ダミー関数）。
// 中置式ではtrueを返す。
func (n *InfixExpressionImpl) IsInfixExpression() bool {
	return true
}

// Left は左のノードを返す。
func (n *InfixExpressionImpl) Left() Node {
	return n.left
}

// SetLeft は左のノードを設定する。
func (n *InfixExpressionImpl) SetLeft(l Node) {
	n.left = l
}

// Operator は演算子を返す。
func (n *InfixExpressionImpl) Operator() string {
	return n.operator
}

// OperatorForSExp はS式で表示する演算子を返す。
func (n *InfixExpressionImpl) OperatorForSExp() string {
	return n.operatorForSExp
}

// Right は右のノードを返す。
func (n *InfixExpressionImpl) Right() Node {
	return n.right
}

// SetRight は右のノードを設定する。
func (n *InfixExpressionImpl) SetRight(r Node) {
	n.right = r
}

// SExp はノードのS式を返す。
func (n *InfixExpressionImpl) SExp() string {
	var leftSExp string
	var rightSExp string

	if n.Left() == nil {
		leftSExp = "nil"
	} else {
		leftSExp = n.Left().SExp()
	}

	if n.Right() == nil {
		rightSExp = "nil"
	} else {
		rightSExp = n.Right().SExp()
	}

	return fmt.Sprintf("(%s %s %s)",
		n.OperatorForSExp(), leftSExp, rightSExp)
}

// IsPrimaryExpression は一次式かどうかを返す。
// 中置式ではfalseを返す。
func (n *InfixExpressionImpl) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// 中置式では、左または右のノードが可変ノードならばtrueを返す。
// 左右の両方のノードが可変ノードでない場合はfalseを返す。
func (n *InfixExpressionImpl) IsVariable() bool {
	return n.Left().IsVariable() || n.Right().IsVariable()
}
