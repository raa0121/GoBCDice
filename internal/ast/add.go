package ast

import (
	"github.com/raa0121/GoBCDice/internal/token"
)

// 加算のノード
type Add struct {
	InfixExpressionImpl
}

// AddがNodeを実装していることの確認
var _ Node = (*Add)(nil)

// Typeはノードの種類を返す
func (n *Add) Type() NodeType {
	return ADD_NODE
}

// NewAddは、加算のノードを返す
//
// * left: 加えられる数のノード
// * tok: 対応するトークン
// * right: 加える数のノード
func NewAdd(left Node, tok token.Token, right Node) *Add {
	return &Add{
		InfixExpressionImpl: *NewInfixExpression(left, tok, right),
	}
}
