package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

type NodeType int

func (t NodeType) String() string {
	if str, ok := nodeTypeString[t]; ok {
		return str
	}

	return nodeTypeString[UNKNOWN_NODE]
}

const (
	UNKNOWN_NODE NodeType = iota

	COMMAND_NODE
	PREFIX_EXPRESSION_NODE
	INFIX_EXPRESSION_NODE
	INT_NODE
)

var nodeTypeString = map[NodeType]string{
	UNKNOWN_NODE: "UNKNOWN",

	COMMAND_NODE:           "Command",
	PREFIX_EXPRESSION_NODE: "PrefixExpression",
	INFIX_EXPRESSION_NODE:  "InfixExpression",
	INT_NODE:               "Int",
}

// 抽象構文木のノードのインターフェース
type Node interface {
	// Tokenは対応するトークンを返す
	Token() token.Token
	// Typeはノードの種類を返す
	Type() NodeType
	// SExpはノードのS式を返す
	SExp() string
}
