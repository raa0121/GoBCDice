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

	// プラス
	PLUS
	// マイナス
	MINUS
	// アスタリスク
	ASTERISK
	// スラッシュ
	SLASH

	// 等号
	EQ
	// 小なり
	LT
	// 大なり
	GT

	// 開き括弧
	L_PAREN
	// 閉じ括弧
	R_PAREN
	// 開き角括弧
	L_BRACKET
	// 閉じ角括弧
	R_BRACKET

	// ダイスロール：加算ロールとD66ロール
	D_ROLL
	// バラバラロール
	B_ROLL
	// 個数振り足しロール、切り上げ
	R
	// 上方無限ロール、四捨五入
	U
	// シークレットロール
	SECRET
	// ランダム数値の埋め込み "..."
	DOTS

	// 計算
	CALC
	// ランダム選択
	CHOICE
)

var tokenTypeString = map[TokenType]string{
	EOT:     "EOT",
	ILLEGAL: "ILLEGAL",

	IDENT: "IDENT",
	INT:   "INT",

	PLUS:     "+",
	MINUS:    "-",
	ASTERISK: "*",
	SLASH:    "/",

	EQ: "=",
	LT: "<",
	GT: ">",

	L_PAREN:   "(",
	R_PAREN:   ")",
	L_BRACKET: "[",
	R_BRACKET: "]",

	D_ROLL: "D_ROLL",
	B_ROLL: "B_ROLL",
	R:      "R",
	U:      "U",
	SECRET: "SECRET",
	DOTS:   "...",

	CALC:   "CALC",
	CHOICE: "CHOICE",
}

// 識別子 -> キーワードの対応表
var keywords = map[string]TokenType{
	"D":      D_ROLL,
	"B":      B_ROLL,
	"R":      R,
	"U":      U,
	"S":      SECRET,
	"C":      CALC,
	"CHOICE": CHOICE,
}

// LookUpIdentは、identがキーワードかどうかを調べ、トークンの種類を返す
// キーワードならば特別なトークンの種類を、そうでなければIDENTを返す
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
