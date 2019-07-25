package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
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

// Type はノードの種類を返す。
func (n *PrefixExpressionImpl) Type() NodeType {
	return PREFIX_EXPRESSION_NODE
}

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

// IsPrimaryExpression は一次式かどうかを返す。
// 前置式ではfalseを返す。
func (n *PrefixExpressionImpl) IsPrimaryExpression() bool {
	return false
}

// IsVariable は可変ノードかどうかを返す。
//
// 前置式では、右のノードが可変ノードならばtrueを返す。
// 右のノードが可変ノードでない場合はfalseを返す。
func (n *PrefixExpressionImpl) IsVariable() bool {
	return n.Right().IsVariable()
}

// NewPrefixExpression は新しい前置式のノードを返す。
// 評価時とS式とで演算子を変更しなくてもよい場合に使う。
//
// tok: 対応するトークン,
// right: 右のノード。
func NewPrefixExpression(tok token.Token, right Node) *PrefixExpressionImpl {
	return &PrefixExpressionImpl{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		operator:        tok.Literal,
		operatorForSExp: tok.Literal,
		right:           right,
	}
}
