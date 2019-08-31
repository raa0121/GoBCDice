package object

import (
	"bytes"
)

// 個数振り足しロールの成功数カウントの結果を表すオブジェクト。
type RRollCompResult struct {
	// 出目のグループの配列
	ValueGroups *Array
	// 成功数
	NumOfSuccesses *Integer
}

// NewRRollCompResult は、新しい個数振り足しロールの成功数カウントの結果を表すオブジェクトを返す。
func NewRRollCompResult(valueGroups *Array, numOfSuccesses *Integer) *RRollCompResult {
	return &RRollCompResult{
		ValueGroups:    valueGroups,
		NumOfSuccesses: numOfSuccesses,
	}
}

// Type はオブジェクトの種類を返す。
func (r *RRollCompResult) Type() ObjectType {
	return R_ROLL_COMP_RESULT_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (r *RRollCompResult) Inspect() string {
	var out bytes.Buffer

	out.WriteString("<RRollCompResult Values=")
	out.WriteString(r.ValueGroups.Inspect())
	out.WriteString(", NumOfSuccesses=")
	out.WriteString(r.NumOfSuccesses.Inspect())
	out.WriteString(">")

	return out.String()
}
