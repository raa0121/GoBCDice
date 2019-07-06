package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/token"
)

// 中置演算子のノードを表す構造体
type InfixExpression struct {
	// トークン
	Tok token.Token
	// 左のノード
	Left Node
	// 演算子
	Operator string
	// S式で表示する演算子
	OperatorForSExp string
	// 右のノード
	Right Node
}

// Tokenは対応するトークンを返す
func (n *InfixExpression) Token() token.Token {
	return n.Tok
}

// Typeはノードの種類を返す
func (n *InfixExpression) Type() NodeType {
	return INFIX_EXPRESSION_NODE
}

// SExpはノードのS式を返す
func (n *InfixExpression) SExp() string {
	return fmt.Sprintf("(%s %s %s)",
		n.OperatorForSExp, n.Left.SExp(), n.Right.SExp())
}

// NewDivideWithRoundingUpは、小数点以下を切り上げる除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
func NewDivideWithRoundingUp(dividend Node, tok token.Token, divisor Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            dividend,
		Operator:        "/U",
		OperatorForSExp: "/U",
		Right:           divisor,
	}
}

// NewDivideWithRoundingは、小数点以下を四捨五入する除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
func NewDivideWithRounding(dividend Node, tok token.Token, divisor Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            dividend,
		Operator:        "/R",
		OperatorForSExp: "/R",
		Right:           divisor,
	}
}

// NewDivideWithRoundingDownは、小数点以下を切り捨てる除算のノードを返す
//
// * dividend: 被除数のノード
// * tok: 対応するトークン
// * divisor: 除数のノード
func NewDivideWithRoundingDown(dividend Node, tok token.Token, divisor Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            dividend,
		Operator:        "/",
		OperatorForSExp: "/",
		Right:           divisor,
	}
}

// NewDRollは加算ロールのノードを返す
//
// * num: 振るダイスの数のノード
// * tok: 対応するトークン
// * sides: ダイスの面数のノード
func NewDRoll(num Node, tok token.Token, sides Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            num,
		Operator:        "D",
		OperatorForSExp: "DRoll",
		Right:           sides,
	}
}

// NewRandはランダム数値取り出しのノードを返す
//
// * min: 最小値のノード
// * tok: 対応するトークン
// * max: 最大値のノード
func NewRand(min Node, tok token.Token, max Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            min,
		Operator:        "...",
		OperatorForSExp: "Rand",
		Right:           max,
	}
}

// NewInfixExpressionは中置演算子のノードを返す。
// 評価時とS式とで演算子の表示を変更しなくてもよい場合に使う。
//
// * left: 左のノード
// * tok: 対応するトークン
// * right: 右のノード
func NewInfixExpression(left Node, tok token.Token, right Node) *InfixExpression {
	return &InfixExpression{
		Tok:             tok,
		Left:            left,
		Operator:        tok.Literal,
		OperatorForSExp: tok.Literal,
		Right:           right,
	}
}
