package ast

// 上方無限ロールのノード。
// 一次式、可変ノード、中置式。
type URoll struct {
	InfixExpressionImpl
}

// URoll がNodeを実装していることの確認。
var _ Node = (*URoll)(nil)

// URoll がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*URoll)(nil)

// Type はノードの種類を返す。
func (n *URoll) Type() NodeType {
	return U_ROLL_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *URoll) Precedence() OperatorPrecedenceType {
	return PREC_ROLL
}

// IsLeftAssociative は左結合性かどうかを返す。
// URollではfalseを返す。
func (n *URoll) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// URollではfalseを返す。
func (n *URoll) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpression は一次式かどうかを返す。
// URollではtrueを返す。
func (n *URoll) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// URollではtrueを返す。
func (n *URoll) IsVariable() bool {
	return true
}

// NewURoll はバラバラロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewURoll(num Node, sides Node) *URoll {
	return &URoll{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            num,
			operator:        "U",
			operatorForSExp: "URoll",
			right:           sides,
		},
	}
}
