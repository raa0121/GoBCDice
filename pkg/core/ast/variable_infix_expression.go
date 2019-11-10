package ast

// VariableInfixExpression は可変の中置式ノード。
type VariableInfixExpression struct {
	InfixExpressionImpl
	VariableNode
}

// variableInfixExpressionOperator はノードの種類と演算子との対応。
var variableInfixExpressionOperator = map[NodeType]string{
	D_ROLL_NODE:        "D",
	B_ROLL_NODE:        "B",
	R_ROLL_NODE:        "R",
	U_ROLL_NODE:        "U",
	RANDOM_NUMBER_NODE: "...",
}

// variableInfixExpressionPrecedence はノードの種類と演算子の優先順位との対応。
var variableInfixExpressionPrecedence = map[NodeType]OperatorPrecedenceType{
	D_ROLL_NODE:        PREC_ROLL,
	B_ROLL_NODE:        PREC_ROLL,
	R_ROLL_NODE:        PREC_ROLL,
	U_ROLL_NODE:        PREC_ROLL,
	RANDOM_NUMBER_NODE: PREC_DOTS,
}

func newVariableInfixExpression(
	left Node,
	right Node,
	nodeType NodeType,
) *VariableInfixExpression {
	return &VariableInfixExpression{
		InfixExpressionImpl: InfixExpressionImpl{
			NodeImpl: NodeImpl{
				nodeType:            nodeType,
				isPrimaryExpression: true,
			},

			left:               left,
			operator:           variableInfixExpressionOperator[nodeType],
			operatorForSExp:    nodeType.String(),
			right:              right,
			precedence:         variableInfixExpressionPrecedence[nodeType],
			isLeftAssociative:  false,
			isRightAssociative: false,
		},
	}
}

// NewDRoll は新しい加算ロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewDRoll(num Node, sides Node) *VariableInfixExpression {
	return newVariableInfixExpression(num, sides, D_ROLL_NODE)
}

// NewBRoll はバラバラロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewBRoll(num Node, sides Node) *VariableInfixExpression {
	return newVariableInfixExpression(num, sides, B_ROLL_NODE)
}

// NewRRoll は新しい個数振り足しロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
func NewRRoll(num Node, sides Node) *VariableInfixExpression {
	return newVariableInfixExpression(num, sides, R_ROLL_NODE)
}

// NewURoll は新しい上方無限ロールのノードを返す。
//
// num: 振るダイスの数のノード,
// sides: ダイスの面数のノード。
//
// 機能はRRollと同じなので、同じ構造体を使用している。
// 演算子の表記のみRRollと異なる。
func NewURoll(num Node, sides Node) *VariableInfixExpression {
	return newVariableInfixExpression(num, sides, U_ROLL_NODE)
}

// NewRandomNumber はランダム数値取り出しのノードを返す。
//
// min: 最小値のノード,
// max: 最大値のノード。
func NewRandomNumber(min Node, max Node) *VariableInfixExpression {
	return newVariableInfixExpression(min, max, RANDOM_NUMBER_NODE)
}
