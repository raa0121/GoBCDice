package ast

// 演算子の優先順位の型。
type OperatorPrecedenceType int

const (
	PREC_MIN OperatorPrecedenceType = iota

	// Compare (=, >, <, >=, <=, <>)
	PREC_COMPARE
	// Add, Subtract
	PREC_ADDITIVE
	// Multiply, Divide
	PREC_MULTIPLICATIVE
	// ダイスロール
	PREC_ROLL
	// RandomNumber
	PREC_DOTS
	// 単項マイナス、単項プラス
	PREC_U_PLUS_MINUS
)
