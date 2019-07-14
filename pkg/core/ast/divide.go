package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 除算の端数処理の方法を表す型。
type RoundingMethodType int

const (
	// 小数点以下切り捨て
	ROUNDING_METHOD_ROUND_DOWN RoundingMethodType = iota
	// 小数点以下四捨五入
	ROUNDING_METHOD_ROUND
	// 小数点以下切り上げ
	ROUNDING_METHOD_ROUND_UP
)

// 端数処理方法に対応する文字列
var roundingMethodString = map[RoundingMethodType]string{
	ROUNDING_METHOD_ROUND_UP:   "U",
	ROUNDING_METHOD_ROUND:      "R",
	ROUNDING_METHOD_ROUND_DOWN: "",
}

// String は端数処理方法に対応する文字列を返す。
func (t RoundingMethodType) String() string {
	if s, ok := roundingMethodString[t]; ok {
		return s
	}

	return "UNKNOWN"
}

// 除算のインターフェース。
type Divide interface {
	InfixExpression

	// IsDivide は除算かどうかを返す（ダミーメソッド）。
	IsDivide() bool
	// RoundingMethod は端数処理の方法を返す。
	RoundingMethod() RoundingMethodType
}

// 除算のノードに共通する要素。
type DivideImpl struct{}

// 小数点以下を切り上げる除算のノード。
// 中置式、除算。
type DivideWithRoundingUp struct {
	InfixExpressionImpl
	DivideImpl
}

// 小数点以下を四捨五入する除算のノード。
// 中置式、除算。
type DivideWithRounding struct {
	InfixExpressionImpl
	DivideImpl
}

// 小数点以下を切り捨てる除算のノード。
// 中置式、除算。
type DivideWithRoundingDown struct {
	InfixExpressionImpl
	DivideImpl
}

// DivideWithRoundingUp がNodeを実装していることの確認。
var _ Node = (*DivideWithRoundingUp)(nil)

// DivideWithRounding がNodeを実装していることの確認。
var _ Node = (*DivideWithRounding)(nil)

// DivideWithRoundingDown がNodeを実装していることの確認。
var _ Node = (*DivideWithRoundingDown)(nil)

// DivideWithRoundingUp がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*DivideWithRoundingUp)(nil)

// DivideWithRounding がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*DivideWithRounding)(nil)

// DivideWithRoundingDown がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*DivideWithRoundingDown)(nil)

// DivideWithRoundingUp がDivideを実装していることの確認。
var _ Divide = (*DivideWithRoundingUp)(nil)

// DivideWithRounding がDivideを実装していることの確認。
var _ Divide = (*DivideWithRounding)(nil)

// DivideWithRoundingDown がDivideを実装していることの確認。
var _ Divide = (*DivideWithRoundingDown)(nil)

// IsDivide は除算かどうかを返す（ダミーメソッド）。
// trueを返す。
func (n *DivideImpl) IsDivide() bool {
	return true
}

// Precedence は演算子の優先順位を返す。
func (n *DivideImpl) Precedence() OperatorPrecedenceType {
	return PREC_MULTITIVE
}

// IsLeftAssociative は左結合性かどうかを返す。
// 除算ではtrueを返す。
func (n *DivideImpl) IsLeftAssociative() bool {
	return true
}

// IsRightAssociative は右結合性かどうかを返す。
// 除算ではfalseを返す。
func (n *DivideImpl) IsRightAssociative() bool {
	return false
}

// Type はノードの種類を返す。
func (n *DivideWithRoundingUp) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_UP_NODE
}

// Type はノードの種類を返す。
func (n *DivideWithRounding) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_NODE
}

// Type はノードの種類を返す。
func (n *DivideWithRoundingDown) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_DOWN_NODE
}

// RoundingMethod は端数処理の方法を返す。
func (n *DivideWithRoundingUp) RoundingMethod() RoundingMethodType {
	return ROUNDING_METHOD_ROUND_UP
}

// RoundingMethod は端数処理の方法を返す。
func (n *DivideWithRounding) RoundingMethod() RoundingMethodType {
	return ROUNDING_METHOD_ROUND
}

// RoundingMethod は端数処理の方法を返す。
func (n *DivideWithRoundingDown) RoundingMethod() RoundingMethodType {
	return ROUNDING_METHOD_ROUND_DOWN
}

// NewDivideWithRoundingUp は小数点以下を切り上げる除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// tok: 対応するトークン,
// divisor: 除数のノード。
func NewDivideWithRoundingUp(dividend Node, tok token.Token, divisor Node) *DivideWithRoundingUp {
	return &DivideWithRoundingUp{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            dividend,
			operator:        "/U",
			operatorForSExp: "/U",
			right:           divisor,
		},
	}
}

// NewDivideWithRounding は小数点以下を四捨五入する除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// tok: 対応するトークン,
// divisor: 除数のノード。
func NewDivideWithRounding(dividend Node, tok token.Token, divisor Node) *DivideWithRounding {
	return &DivideWithRounding{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            dividend,
			operator:        "/R",
			operatorForSExp: "/R",
			right:           divisor,
		},
	}
}

// NewDivideWithRoundingDown は小数点以下を切り捨てる除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// tok: 対応するトークン,
// divisor: 除数のノード。
func NewDivideWithRoundingDown(dividend Node, tok token.Token, divisor Node) *DivideWithRoundingDown {
	return &DivideWithRoundingDown{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				tok: tok,
			},
			left:            dividend,
			operator:        "/",
			operatorForSExp: "/",
			right:           divisor,
		},
	}
}
