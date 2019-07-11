package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 除算のインターフェース
type Divide interface {
	// IsDivideは除算かどうかを返す（ダミーメソッド）
	IsDivide() bool
	// Precedenceは演算子の優先順位を返す
	Precedence() OperatorPrecedenceType
	// IsLeftAssociativeは左結合性かどうかを返す
	IsLeftAssociative() bool
	// IsRightAssociativeは右結合性かどうかを返す
	IsRightAssociative() bool
}

// 除算の共通要素を定義する構造体
type DivideImpl struct{}

// 小数点以下を切り上げる除算のノード
type DivideWithRoundingUp struct {
	InfixExpressionImpl
	DivideImpl
}

// 小数点以下を四捨五入する除算のノード
type DivideWithRounding struct {
	InfixExpressionImpl
	DivideImpl
}

// 小数点以下を切り捨てる除算のノード
type DivideWithRoundingDown struct {
	InfixExpressionImpl
	DivideImpl
}

// DivideWithRoundingUpがNodeを実装していることの確認
var _ Node = (*DivideWithRoundingUp)(nil)

// DivideWithRoundingがNodeを実装していることの確認
var _ Node = (*DivideWithRounding)(nil)

// DivideWithRoundingDownがNodeを実装していることの確認
var _ Node = (*DivideWithRoundingDown)(nil)

// DivideWithRoundingUpがInfixExpressionを実装していることの確認
var _ InfixExpression = (*DivideWithRoundingUp)(nil)

// DivideWithRoundingがInfixExpressionを実装していることの確認
var _ InfixExpression = (*DivideWithRounding)(nil)

// DivideWithRoundingDownがInfixExpressionを実装していることの確認
var _ InfixExpression = (*DivideWithRoundingDown)(nil)

// DivideWithRoundingUpがDivideを実装していることの確認
var _ Divide = (*DivideWithRoundingUp)(nil)

// DivideWithRoundingがDivideを実装していることの確認
var _ Divide = (*DivideWithRounding)(nil)

// DivideWithRoundingDownがDivideを実装していることの確認
var _ Divide = (*DivideWithRoundingDown)(nil)

// IsDivideは除算かどうかを返す（ダミーメソッド）
func (n *DivideImpl) IsDivide() bool {
	return true
}

// Precedenceは演算子の優先順位を返す
func (n *DivideImpl) Precedence() OperatorPrecedenceType {
	return PREC_MULTITIVE
}

// IsLeftAssociativeは左結合性かどうかを返す
func (n *DivideImpl) IsLeftAssociative() bool {
	return true
}

// IsRightAssociativeは右結合性かどうかを返す
func (n *DivideImpl) IsRightAssociative() bool {
	return false
}

// Typeはノードの種類を返す
func (n *DivideWithRoundingUp) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_UP_NODE
}

// Typeはノードの種類を返す
func (n *DivideWithRounding) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_NODE
}

// Typeはノードの種類を返す
func (n *DivideWithRoundingDown) Type() NodeType {
	return DIVIDE_WITH_ROUNDING_DOWN_NODE
}

// IsCommutativeは可換演算子かどうかを返す
func (n *DivideWithRoundingUp) IsCommutative() bool {
	return false
}

// IsCommutativeは可換演算子かどうかを返す
func (n *DivideWithRounding) IsCommutative() bool {
	return false
}

// IsCommutativeは可換演算子かどうかを返す
func (n *DivideWithRoundingDown) IsCommutative() bool {
	return false
}

// NewDivideWithRoundingUpは、小数点以下を切り上げる除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
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

// NewDivideWithRoundingは、小数点以下を四捨五入する除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
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

// NewDivideWithRoundingDownは、小数点以下を切り捨てる除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
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
