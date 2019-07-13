package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// 前置演算子のインターフェース
type PrefixExpression interface {
	Node
	IsPrefixExpression() bool
	// Operatorは演算子を返す
	Operator() string
	// OperatorForSExpはS式で表示する演算子を返す
	OperatorForSExp() string
	// Rightは右のノードを返す
	Right() Node
	// SetRightは右のノードを設定する
	SetRight(r Node)
}

// 前置演算子のノードを表す構造体
type PrefixExpressionImpl struct {
	NodeImpl
	// 演算子
	operator string
	// S式で表示する演算子
	operatorForSExp string
	// 右のノード
	right Node
}

// PrefixExpressionがNodeを実装していることの確認
var _ Node = (*PrefixExpressionImpl)(nil)

// Typeはノードの種類を返す
func (n *PrefixExpressionImpl) Type() NodeType {
	return PREFIX_EXPRESSION_NODE
}

func (n *PrefixExpressionImpl) IsPrefixExpression() bool {
	return true
}

// Operatorは演算子を返す
func (n *PrefixExpressionImpl) Operator() string {
	return n.operator
}

// OperatorForSExpはS式で表示する演算子を返す
func (n *PrefixExpressionImpl) OperatorForSExp() string {
	return n.operatorForSExp
}

// Rightは右のノードを返す
func (n *PrefixExpressionImpl) Right() Node {
	return n.right
}

// SetRightは右のノードを設定する
func (n *PrefixExpressionImpl) SetRight(r Node) {
	n.right = r
}

// SExpはノードのS式を返す
func (n *PrefixExpressionImpl) SExp() string {
	return fmt.Sprintf("(%s %s)", n.OperatorForSExp(), n.Right().SExp())
}

// IsPrimaryExpressionは一次式かどうかを返す
func (n *PrefixExpressionImpl) IsPrimaryExpression() bool {
	return false
}

// IsVariableは可変ノードかどうかを返す。
func (n *PrefixExpressionImpl) IsVariable() bool {
	return n.Right().IsVariable()
}

// NewPrefixExpressionは前置演算子のノードを返す。
// 評価時とS式とで演算子の表示を変更しなくてもよい場合に使う。
//
// * tok: 対応するトークン
// * right: 右のノード
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
