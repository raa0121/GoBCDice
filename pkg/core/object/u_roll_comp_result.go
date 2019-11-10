package object

import (
	"bytes"
)

// 上方無限ロールの成功数カウントの結果を表すオブジェクト。
type URollCompResult struct {
	// RollResult は上方無限ロールの結果。
	RollResult *URollExprResult
	// NumOfSuccesses は成功数を表す整数オブジェクト。
	NumOfSuccesses *Integer
}

// NewRRollCompResult は、新しい個数振り足しロールの成功数カウントの結果を表すオブジェクトを返す。
func NewURollCompResult(rollResult *URollExprResult, numOfSuccesses *Integer) *URollCompResult {
	return &URollCompResult{
		RollResult:     rollResult,
		NumOfSuccesses: numOfSuccesses,
	}
}

// Type はオブジェクトの種類を返す。
func (r *URollCompResult) Type() ObjectType {
	return U_ROLL_COMP_RESULT_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (r *URollCompResult) Inspect() string {
	var out bytes.Buffer

	out.WriteString("<URollCompResult RollResult=")
	out.WriteString(r.RollResult.Inspect())
	out.WriteString(", NumOfSuccesses=")
	out.WriteString(r.NumOfSuccesses.Inspect())
	out.WriteString(">")

	return out.String()
}
