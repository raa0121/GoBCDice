package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// NodeTypeはノードの種類を表す型
type NodeType int

// Stringはノードの種類を文字列として返す
func (t NodeType) String() string {
	if str, ok := nodeTypeString[t]; ok {
		return str
	}

	return nodeTypeString[UNKNOWN_NODE]
}

const (
	UNKNOWN_NODE NodeType = iota

	D_ROLL_EXPR_NODE
	CALC_NODE

	PREFIX_EXPRESSION_NODE
	UNARY_MINUS_NODE

	INFIX_EXPRESSION_NODE
	ADD_NODE
	SUBTRACT_NODE
	MULTIPLY_NODE
	DIVIDE_WITH_ROUNDING_UP_NODE
	DIVIDE_WITH_ROUNDING_NODE
	DIVIDE_WITH_ROUNDING_DOWN_NODE
	D_ROLL_NODE
	RANDOM_NUMBER_NODE

	INT_NODE
	SUM_ROLL_RESULT_NODE
)

// ノードの種類とそれを表す文字列との対応
var nodeTypeString = map[NodeType]string{
	UNKNOWN_NODE: "UNKNOWN",

	D_ROLL_EXPR_NODE:               "DRollExpr",
	CALC_NODE:                      "Calc",
	PREFIX_EXPRESSION_NODE:         "PrefixExpression",
	UNARY_MINUS_NODE:               "UnaryMinus",
	INFIX_EXPRESSION_NODE:          "InfixExpression",
	ADD_NODE:                       "Add",
	SUBTRACT_NODE:                  "Subtract",
	MULTIPLY_NODE:                  "Multiply",
	DIVIDE_WITH_ROUNDING_UP_NODE:   "DivideWithRoundingUp",
	DIVIDE_WITH_ROUNDING_NODE:      "DivideWithRounding",
	DIVIDE_WITH_ROUNDING_DOWN_NODE: "DivideWithRoundingDown",
	D_ROLL_NODE:                    "DRoll",
	RANDOM_NUMBER_NODE:             "RandomNumber",
	INT_NODE:                       "Int",
	SUM_ROLL_RESULT_NODE:           "SumRollResult",
}

// 抽象構文木のノードのインターフェース
type Node interface {
	// Tokenは対応するトークンを返す
	Token() token.Token
	// Typeはノードの種類を返す
	Type() NodeType
	// SExpはノードのS式を返す
	SExp() string
	// IsPrimaryExpressionは一次式かどうかを返す
	IsPrimaryExpression() bool
	// IsVariableは可変ノードかどうかを返す。
	// 可変ノードとは、ダイスロールやランダム数値の取り出しなど、
	// 実行のたびに値が変わり得るノードのこと。
	IsVariable() bool
}

// Nodeが共通で持つ要素
type NodeImpl struct {
	// トークン
	tok token.Token
}

// Tokenは対応するトークンを返す
func (n *NodeImpl) Token() token.Token {
	return n.tok
}
