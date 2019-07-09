package ast

type PrecType int

const (
	PREC_MIN PrecType = iota
	PREC_ADDITIVE
	PREC_MULTITIVE
	PREC_D_ROLL
	PREC_DOTS
	PREC_U_PLUS_MINUS
)
