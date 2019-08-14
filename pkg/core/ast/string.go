package ast

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/token"
)

// 文字列のノード。
type String struct {
	NodeImpl

	// 文字列
	Value string
}

// String がNodeを実装していることの確認。
var _Node = (*String)(nil)

// Type はノードの種類を返す。
func (n *String) Type() NodeType {
	return STRING_NODE
}

// SExp はノードのS式を返す。
func (n *String) SExp() string {
	return fmt.Sprintf("%q", n.Value)
}

// IsPrimaryExpression は一次式かどうかを返す。
// Stringではtrueを返す。
func (n *String) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// Stringではfalseを返す。
func (n *String) IsVariable() bool {
	return false
}

// NewString は新しい文字列のノードを返す。
//
// value: 文字列,
// tok: 対応するトークン。
func NewString(value string, tok token.Token) *String {
	return &String{
		NodeImpl: NodeImpl{
			tok: tok,
		},
		Value: value,
	}
}

// NewString2 は新しい文字列のノードを返す。
//
// value: 文字列,
func NewString2(value string) *String {
	return &String{
		Value: value,
	}
}
