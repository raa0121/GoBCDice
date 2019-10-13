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

// InfixExpressionImpl は中置式のノードが共通して持つ要素。
type InfixExpressionImpl struct {
	NodeImpl
	NonNilNode

	// left は左のノード。
	left Node
	// operator は演算子。
	operator string
	// operatorForSExp はS式で表示する演算子。
	operatorForSExp string
	// right は右のノード。
	right Node
	// precedence は演算子の優先順位。
	precedence OperatorPrecedenceType
	// isLeftAssociative は左結合性かどうか。
	isLeftAssociative bool
	// isRightAssociative は右結合性かどうか。
	isRightAssociative bool
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

// Precedence は演算子の優先順位を返す。
func (n *InfixExpressionImpl) Precedence() OperatorPrecedenceType {
	return n.precedence
}

// IsLeftAssociative は左結合性かどうかを返す。
func (n *InfixExpressionImpl) IsLeftAssociative() bool {
	return n.isLeftAssociative
}

// IsRightAssociative は右結合性かどうかを返す。
func (n *InfixExpressionImpl) IsRightAssociative() bool {
	return n.isRightAssociative
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
