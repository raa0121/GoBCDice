package ast

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

// Divide は除算のノード。
type Divide struct {
	BasicInfixExpression

	// RoundingMethod は端数処理の方法を返す。
	RoundingMethod RoundingMethodType
}

// Divide がNodeを実装していることの確認。
var _ Node = (*Divide)(nil)

// Divide がInfixExpressionを実装していることの確認。
var _ InfixExpression = (*Divide)(nil)

// nodeTypeToRoundingMethod はノードの種類と端数処理の方法との対応。
var nodeTypeToRoundingMethod = map[NodeType]RoundingMethodType{
	DIVIDE_WITH_ROUNDING_UP_NODE:   ROUNDING_METHOD_ROUND_UP,
	DIVIDE_WITH_ROUNDING_NODE:      ROUNDING_METHOD_ROUND,
	DIVIDE_WITH_ROUNDING_DOWN_NODE: ROUNDING_METHOD_ROUND_DOWN,
}

// newDivide は除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// divisor: 除数のノード,
// roundingMethod: 端数処理の方法。
func newDivide(dividend Node, divisor Node, nodeType NodeType) *Divide {
	roundingMethod := nodeTypeToRoundingMethod[nodeType]
	operator := "/" + roundingMethod.String()

	return &Divide{
		BasicInfixExpression: BasicInfixExpression{
			InfixExpressionImpl: InfixExpressionImpl{
				NodeImpl: NodeImpl{
					nodeType:            nodeType,
					isPrimaryExpression: false,
				},

				left:               dividend,
				operator:           operator,
				operatorForSExp:    operator,
				right:              divisor,
				precedence:         PREC_MULTITIVE,
				isLeftAssociative:  true,
				isRightAssociative: false,
			},
		},

		RoundingMethod: roundingMethod,
	}
}

// NewDivideWithRoundingUp は小数点以下を切り上げる除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// divisor: 除数のノード。
func NewDivideWithRoundingUp(dividend Node, divisor Node) *Divide {
	return newDivide(dividend, divisor, DIVIDE_WITH_ROUNDING_UP_NODE)
}

// NewDivideWithRounding は小数点以下を四捨五入する除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// divisor: 除数のノード。
func NewDivideWithRounding(dividend Node, divisor Node) *Divide {
	return newDivide(dividend, divisor, DIVIDE_WITH_ROUNDING_NODE)
}

// NewDivideWithRoundingDown は小数点以下を切り捨てる除算の新しいノードを返す。
//
// dividend: 被除数のノード,
// divisor: 除数のノード。
func NewDivideWithRoundingDown(dividend Node, divisor Node) *Divide {
	return newDivide(dividend, divisor, DIVIDE_WITH_ROUNDING_DOWN_NODE)
}
