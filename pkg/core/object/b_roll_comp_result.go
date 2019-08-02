package object

import (
	"bytes"
)

// バラバラロールの成功数カウントの結果を表すオブジェクト。
type BRollCompResult struct {
	// 出目
	Values *Array
	// 成功数
	NumOfSuccesses *Integer
}

// NewBRollCompResult は、新しいバラバラロールの成功数カウントの結果を表すオブジェクトを返す。
func NewBRollCompResult(values *Array, numOfSuccesses *Integer) *BRollCompResult {
	return &BRollCompResult{
		Values:         values,
		NumOfSuccesses: numOfSuccesses,
	}
}

// Type はオブジェクトの種類を返す。
func (r *BRollCompResult) Type() ObjectType {
	return B_ROLL_COMP_RESULT_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (r *BRollCompResult) Inspect() string {
	var out bytes.Buffer

	out.WriteString("<BRollCompResult Values=")
	out.WriteString(r.Values.Inspect())
	out.WriteString(", NumOfSuccesses=")
	out.WriteString(r.NumOfSuccesses.Inspect())
	out.WriteString(">")

	return out.String()
}
