/*
BCDiceコマンドの評価結果として生成される数値などのオブジェクトの内部表現のパッケージ。
*/
package object

// オブジェクトの種類を表す型。
type ObjectType int

// String はオブジェクトの種類を文字列として返す。
func (t ObjectType) String() string {
	if str, ok := objectTypeString[t]; ok {
		return str
	}

	return objectTypeString[ILLEGAL_OBJ]
}

const (
	ILLEGAL_OBJ ObjectType = iota

	INTEGER_OBJ
	BOOLEAN_OBJ
	STRING_OBJ
	ARRAY_OBJ
	B_ROLL_COMP_RESULT_OBJ
	R_ROLL_COMP_RESULT_OBJ
)

// オブジェクトの種類とそれを表す文字列との対応
var objectTypeString = map[ObjectType]string{
	ILLEGAL_OBJ: "ILLEGAL",

	INTEGER_OBJ:            "INTEGER",
	BOOLEAN_OBJ:            "BOOLEAN",
	STRING_OBJ:             "STRING",
	ARRAY_OBJ:              "ARRAY",
	B_ROLL_COMP_RESULT_OBJ: "B_ROLL_COMP_RESULT",
	R_ROLL_COMP_RESULT_OBJ: "R_ROLL_COMP_RESULT",
}

// オブジェクトが持つインターフェース。
type Object interface {
	// Type はオブジェクトの種類を返す。
	Type() ObjectType
	// Inspect はオブジェクトの内容を文字列として返す。
	Inspect() string
}
