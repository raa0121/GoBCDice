package ast

// NodeType はノードの種類を表す型。
type NodeType int

// String はノードの種類を文字列として返す。
func (t NodeType) String() string {
	if str, ok := nodeTypeString[t]; ok {
		return str
	}

	return nodeTypeString[UNKNOWN_NODE]
}

const (
	UNKNOWN_NODE NodeType = iota

	D_ROLL_EXPR_NODE
	D_ROLL_COMP_NODE
	B_ROLL_LIST_NODE
	B_ROLL_COMP_NODE
	R_ROLL_LIST_NODE
	R_ROLL_COMP_NODE
	CALC_NODE
	CHOICE_NODE

	PREFIX_EXPRESSION_NODE
	UNARY_MINUS_NODE

	INFIX_EXPRESSION_NODE
	COMPARE_NODE
	ADD_NODE
	SUBTRACT_NODE
	MULTIPLY_NODE
	DIVIDE_WITH_ROUNDING_UP_NODE
	DIVIDE_WITH_ROUNDING_NODE
	DIVIDE_WITH_ROUNDING_DOWN_NODE
	D_ROLL_NODE
	B_ROLL_NODE
	R_ROLL_NODE
	RANDOM_NUMBER_NODE

	INT_NODE
	STRING_NODE
	NIL_NODE
	SUM_ROLL_RESULT_NODE
)

// ノードの種類とそれを表す文字列との対応。
var nodeTypeString = map[NodeType]string{
	UNKNOWN_NODE: "UNKNOWN",

	D_ROLL_EXPR_NODE: "DRollExpr",
	D_ROLL_COMP_NODE: "DRollComp",
	B_ROLL_LIST_NODE: "BRollList",
	B_ROLL_COMP_NODE: "BRollComp",
	R_ROLL_LIST_NODE: "RRollList",
	R_ROLL_COMP_NODE: "RRollComp",
	CALC_NODE:        "Calc",
	CHOICE_NODE:      "Choice",

	PREFIX_EXPRESSION_NODE: "PrefixExpression",
	UNARY_MINUS_NODE:       "UnaryMinus",

	INFIX_EXPRESSION_NODE:          "InfixExpression",
	COMPARE_NODE:                   "Compare",
	ADD_NODE:                       "Add",
	SUBTRACT_NODE:                  "Subtract",
	MULTIPLY_NODE:                  "Multiply",
	DIVIDE_WITH_ROUNDING_UP_NODE:   "DivideWithRoundingUp",
	DIVIDE_WITH_ROUNDING_NODE:      "DivideWithRounding",
	DIVIDE_WITH_ROUNDING_DOWN_NODE: "DivideWithRoundingDown",
	D_ROLL_NODE:                    "DRoll",
	B_ROLL_NODE:                    "BRoll",
	R_ROLL_NODE:                    "RRoll",
	RANDOM_NUMBER_NODE:             "RandomNumber",

	INT_NODE:             "Int",
	STRING_NODE:          "String",
	NIL_NODE:             "Nil",
	SUM_ROLL_RESULT_NODE: "SumRollResult",
}

// 抽象構文木のノードのインターフェース。
type Node interface {
	// Type はノードの種類を返す。
	Type() NodeType
	// SExp はノードのS式を返す。
	SExp() string
	// IsNil はnilかどうかを返す。
	IsNil() bool
	// IsPrimaryExpression は一次式かどうかを返す。
	IsPrimaryExpression() bool
	// IsVariable は可変ノードかどうかを返す。
	// 可変ノードとは、ダイスロールやランダム数値の取り出しなど、実行のたびに値が変わり得るノードのこと。
	IsVariable() bool
}

// NonNilNode はnilでないノードの型。
type NonNilNode struct{}

// IsNil はnilかどうかを返す。
func (n *NonNilNode) IsNil() bool {
	return false
}
