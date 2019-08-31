package ast

// 個数振り足しロールのノード。
// 一次式、可変ノード、中置式。
type RRoll struct {
	InfixExpressionImpl
}

// RRoll がNodeを実装していることの確認。
var _ Node = (*RRoll)(nil)

// RRoll がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*RRoll)(nil)

// Type はノードの種類を返す。
func (n *RRoll) Type() NodeType {
	return R_ROLL_NODE
}

// Precedence は演算子の優先順位を返す。
func (n *RRoll) Precedence() OperatorPrecedenceType {
	return PREC_ROLL
}

// IsLeftAssociative は左結合性かどうかを返す。
// RRollではfalseを返す。
func (n *RRoll) IsLeftAssociative() bool {
	return false
}

// IsRightAssociative は右結合性かどうかを返す。
// RRollではfalseを返す。
func (n *RRoll) IsRightAssociative() bool {
	return false
}

// IsPrimaryExpression は一次式かどうかを返す。
// RRollではtrueを返す。
func (n *RRoll) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// RRollではtrueを返す。
func (n *RRoll) IsVariable() bool {
	return true
}

// NewRRoll はバラバラロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewRRoll(num Node, sides Node) *RRoll {
	return &RRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            num,
			operator:        "R",
			operatorForSExp: "RRoll",
			right:           sides,
		},
	}
}
