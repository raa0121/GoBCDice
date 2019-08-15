package ast

// 単項マイナスのノード。
// 前置式。
type UnaryMinus struct {
	PrefixExpressionImpl
}

// UnaryMinus がNodeを実装していることの確認。
var _ Node = (*UnaryMinus)(nil)

// UnaryMinus がPrefixExpressionを実装していることの確認。
var _ PrefixExpression = (*UnaryMinus)(nil)

// Type はノードの種類を返す。
func (n *UnaryMinus) Type() NodeType {
	return UNARY_MINUS_NODE
}

// NewUnaryMinus は新しい単項マイナスのノードを返す。
//
// right: 右のノード。
func NewUnaryMinus(right Node) *UnaryMinus {
	return &UnaryMinus{
		PrefixExpressionImpl: PrefixExpressionImpl{
			operator:        "-",
			operatorForSExp: "-",
			right:           right,
		},
	}
}
