package object

import (
	"bytes"
	"strings"
)

// 配列オブジェクトの構造体
type Array struct {
	// 要素
	Elements []Object
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect はオブジェクトの内容を文字列として返す。
func (a *Array) Inspect() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(a.JoinedElements(", "))
	out.WriteString("]")

	return out.String()
}

// InspectWithoutSpaces は、Inspectの区切り文字に空白が含まれないもの。
func (a *Array) InspectWithoutSpaces() string {
	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(a.JoinedElements(","))
	out.WriteString("]")

	return out.String()
}

// JoinedElements は要素の内容を区切り文字sepを使って結合した文字列を返す。
func (a *Array) JoinedElements(sep string) string {
	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	return strings.Join(elements, sep)
}
