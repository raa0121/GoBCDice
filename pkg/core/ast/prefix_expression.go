package ast

import (
	"fmt"
)

// 前置式のインターフェース。
type PrefixExpression interface {
	Node

	// IsPrefixExpression は前置式であるかを返す（ダミー関数）。
	IsPrefixExpression() bool
	// Operator は演算子を返す。
	Operator() string
	// OperatorForSExp はS式で表示する演算子を返す。
	OperatorForSExp() string
	// Right は右のノードを返す。
	Right() Node
	// SetRight は右のノードを設定する。
	SetRight(r Node)
}

// 前置式のノードが共通して持つ要素。
type PrefixExpressionImpl struct {
	NonNilNode
	NodeImpl

	// 演算子
	operator string
	// S式で表示する演算子
	operatorForSExp string
	// 右のノード
	right Node
}

// PrefixExpressionImpl がNodeを実装していることの確認。
var _ Node = (*PrefixExpressionImpl)(nil)

// PrefixExpressionImpl がPrefixExpressionを実装していることの確認。
var _ PrefixExpression = (*PrefixExpressionImpl)(nil)

// IsPrefixExpression は前置式であるかを返す（ダミー関数）。
// 前置式ではtrueを返す。
func (n *PrefixExpressionImpl) IsPrefixExpression() bool {
	return true
}

// Operator は演算子を返す。
func (n *PrefixExpressionImpl) Operator() string {
	return n.operator
}

// OperatorForSExp はS式で表示する演算子を返す。
func (n *PrefixExpressionImpl) OperatorForSExp() string {
	return n.operatorForSExp
}

// Right は右のノードを返す。
func (n *PrefixExpressionImpl) Right() Node {
	return n.right
}

// SetRight は右のノードを設定する。
func (n *PrefixExpressionImpl) SetRight(r Node) {
	n.right = r
}

// SExp はノードのS式を返す。
func (n *PrefixExpressionImpl) SExp() string {
	var rightSExp string

	if n.Right() == nil {
		rightSExp = "nil"
	} else {
		rightSExp = n.Right().SExp()
	}

	return fmt.Sprintf("(%s %s)", n.OperatorForSExp(), rightSExp)
}

// IsVariable は可変ノードかどうかを返す。
//
// 前置式では、右のノードが可変ノードならばtrueを返す。
// 右のノードが可変ノードでない場合はfalseを返す。
func (n *PrefixExpressionImpl) IsVariable() bool {
	return n.Right().IsVariable()
}

// NewUnaryMinus は新しい単項マイナスのノードを返す。
//
// right: 右のノード。
func NewUnaryMinus(right Node) *PrefixExpressionImpl {
	return &PrefixExpressionImpl{
		NodeImpl: NodeImpl{
			nodeType:            UNARY_MINUS_NODE,
			isPrimaryExpression: false,
		},

		operator:        "-",
		operatorForSExp: "-",
		right:           right,
	}
}
