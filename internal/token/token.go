package token

// トークンの種類の型
type TokenType int

// トークンを表す構造体
type Token struct {
	// トークンの種類
	Type TokenType
	// リテラル
	Literal string
	// トークンが何文字目で発見されたか
	Column int
}

const (
	// 無効
	ILLEGAL TokenType = iota
	// テキストの終端（End Of Text）
	EOT

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
	// 個数振り足しロール
	R_ROLL
	// 上方無限ロール
	U_ROLL
	// シークレットロール
	SECRET

	// 計算
	CALC
	// ランダム選択
	CHOICE
)

var tokenTypeString = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOT:     "EOT",

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
	R_ROLL: "R_ROLL",
	U_ROLL: "U_ROLL",
	SECRET: "SECRET_ROLL",

	CALC:   "CALC",
	CHOICE: "CHOICE",
}

// 識別子 -> キーワードの対応表
var keywords = map[string]TokenType{
	"D":      D_ROLL,
	"B":      B_ROLL,
	"R":      R_ROLL,
	"U":      U_ROLL,
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

func (tt TokenType) String() string {
	if ttStr, ok := tokenTypeString[tt]; ok {
		return ttStr
	}

	return "UNKNOWN"
}
