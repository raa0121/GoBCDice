/*
BCDiceコマンドのトークンを定義するパッケージ。
主にlexerパッケージから使われる。
*/
package token

import (
	"fmt"
)

// トークンの種類の型
type TokenType int

func (t TokenType) String() string {
	if str, ok := tokenTypeString[t]; ok {
		return str
	}

	return "UNKNOWN"
}

// トークンを表す構造体
type Token struct {
	// トークンの種類
	Type TokenType
	// リテラル
	Literal string
	// トークンが何文字目で発見されたか
	Column int
}

func (t Token) String() string {
	return fmt.Sprintf("<Token %s Literal:%q Column:%d>",
		t.Type, t.Literal, t.Column)
}

const (
	// テキストの終端（End Of Text）
	EOT TokenType = iota

	// 無効
	ILLEGAL

	// 識別子
	IDENT
	// 整数
	INT
	// 文字列
	STRING

	// プラス "+"
	PLUS
	// マイナス "-"
	MINUS
	// アスタリスク "*"
	ASTERISK
	// スラッシュ "/"
	SLASH

	// 等号 "="
	EQ
	// 小なり "<"
	LT
	// 大なり ">"
	GT
	// 以下 "<="
	LTEQ
	// 以上 ">="
	GTEQ
	// ダイアモンド "<>"
	DIAMOND

	// 開き括弧 "("
	L_PAREN
	// 閉じ括弧 ")"
	R_PAREN
	// 開き角括弧 "["
	L_BRACKET
	// 閉じ角括弧 "]"
	R_BRACKET
	// カンマ ","
	COMMA

	// ダイスロール：加算ロールとD66ロール "D"
	D
	// バラバラロール "B"
	B
	// 個数振り足しロール、切り上げ "R"
	R
	// 上方無限ロール、四捨五入 "U"
	U
	// シークレットロール "S"
	SECRET
	// ランダム数値の埋め込み "..."
	DOTS

	// 計算 "C"
	CALC

	// ランダム選択の開始 "CHOICE["
	CHOICE_BEGIN
	// ランダム選択の終了 "]"
	CHOICE_END
)

var tokenTypeString = map[TokenType]string{
	EOT:     "EOT",
	ILLEGAL: "ILLEGAL",

	IDENT:  "IDENT",
	INT:    "INT",
	STRING: "STRING",

	PLUS:     "+",
	MINUS:    "-",
	ASTERISK: "*",
	SLASH:    "/",

	EQ:      "=",
	LT:      "<",
	GT:      ">",
	LTEQ:    "<=",
	GTEQ:    ">=",
	DIAMOND: "<>",

	L_PAREN:   "(",
	R_PAREN:   ")",
	L_BRACKET: "[",
	R_BRACKET: "]",
	COMMA:     ",",

	D:      "D",
	B:      "B",
	R:      "R",
	U:      "U",
	SECRET: "SECRET",
	DOTS:   "...",

	CALC: "CALC",

	CHOICE_BEGIN: "CHOICE[",
	CHOICE_END:   "]",
}

// 識別子 -> キーワードの対応表
var keywords = map[string]TokenType{
	"D": D,
	"B": B,
	"R": R,
	"U": U,
	"S": SECRET,
	"C": CALC,
}

// LookUpIdentは、identがキーワードかどうかを調べ、トークンの種類を返す
// キーワードならば特別なトークンの種類を、そうでなければIDENTを返す
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
