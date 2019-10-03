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

// NewRRoll は新しい個数振り足しロールのノードを返す。
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

// NewURoll は新しい上方無限ロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
//
// 機能はRRollと同じなので、同じ構造体を使用している。
// 演算子の表記のみRRollと異なる。
func NewURoll(num Node, sides Node) *RRoll {
	return &RRoll{
		InfixExpressionImpl: InfixExpressionImpl{
			left:            num,
			operator:        "U",
			operatorForSExp: "URoll",
			right:           sides,
		},
	}
}
