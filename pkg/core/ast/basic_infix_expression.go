package ast

// BasicInfixExpression は通常の中置式ノード。
type BasicInfixExpression struct {
	InfixExpressionImpl
}

// BasicInfixExpression がNodeを実装していることの確認。
var _ Node = (*BasicInfixExpression)(nil)

// IsVariable は可変ノードかどうかを返す。
//
// 通常の中置式では、左または右のノードが可変ノードならばtrueを返す。
// 左右の両方のノードが可変ノードでない場合はfalseを返す。
func (n *BasicInfixExpression) IsVariable() bool {
	return n.Left().IsVariable() || n.Right().IsVariable()
}

// NewAdd は新しい加算のノードを返す。
//
// left: 加えられる数のノード,
// right: 加える数のノード。
func NewAdd(left Node, right Node) *BasicInfixExpression {
	return &BasicInfixExpression{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				nodeType:            ADD_NODE,
				isPrimaryExpression: false,
			},

			left:               left,
			operator:           "+",
			operatorForSExp:    "+",
			right:              right,
			precedence:         PREC_ADDITIVE,
			isLeftAssociative:  true,
			isRightAssociative: true,
		},
	}
}

// NewSubtract は新しい減算のノードを返す。
//
// left: 引かれる数のノード,
// right: 引く数のノード。
func NewSubtract(left Node, right Node) *BasicInfixExpression {
	return &BasicInfixExpression{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				nodeType:            SUBTRACT_NODE,
				isPrimaryExpression: false,
			},

			left:               left,
			operator:           "-",
			operatorForSExp:    "-",
			right:              right,
			precedence:         PREC_ADDITIVE,
			isLeftAssociative:  true,
			isRightAssociative: false,
		},
	}
}

// NewMultiply は新しい乗算のノードを返す。
//
// multiplicand: 被乗数のノード,
// multiplier: 乗数のノード。
func NewMultiply(multiplicand Node, multiplier Node) *BasicInfixExpression {
	return &BasicInfixExpression{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				nodeType:            MULTIPLY_NODE,
				isPrimaryExpression: false,
			},

			left:               multiplicand,
			operator:           "*",
			operatorForSExp:    "*",
			right:              multiplier,
			precedence:         PREC_MULTIPLICATIVE,
			isLeftAssociative:  true,
			isRightAssociative: true,
		},
	}
}

// NewCompare は新しい比較のノードを返す。
//
// left: 左辺のノード,
// op: 比較演算子,
// right: 右辺のノード。
func NewCompare(left Node, op string, right Node) *BasicInfixExpression {
	return &BasicInfixExpression{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				nodeType:            COMPARE_NODE,
				isPrimaryExpression: false,
			},

			left:               left,
			operator:           op,
			operatorForSExp:    op,
			right:              right,
			precedence:         PREC_COMPARE,
			isLeftAssociative:  false,
			isRightAssociative: false,
		},
	}
}
