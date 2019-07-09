package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 小数点以下を切り上げる除算のノード
type DivideWithRoundingUp struct {
	InfixExpressionImpl
}

// 小数点以下を四捨五入する除算のノード
type DivideWithRounding struct {
	InfixExpressionImpl
}

// 小数点以下を切り捨てる除算のノード
type DivideWithRoundingDown struct {
	InfixExpressionImpl
}

// DivideWithRoundingUpがNodeを実装していることの確認
var _ Node = (*DivideWithRoundingUp)(nil)

// DivideWithRoundingがNodeを実装していることの確認
var _ Node = (*DivideWithRounding)(nil)

// DivideWithRoundingDownがNodeを実装していることの確認
var _ Node = (*DivideWithRoundingDown)(nil)

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
