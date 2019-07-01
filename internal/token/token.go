package token

// トークンの種類の型
type TokenType string

// トークンを表す構造体
type Token struct {
	// トークンの種類
	Type TokenType
	// リテラル
	Literal string
}

const (
	// 無効
	ILLEGAL = "ILLEGAL"
	// テキストの終端（End Of Text）
	EOT = "EOT"

	// 識別子
	IDENT = "IDENT"
	// 整数
	INT = "INT"

	// プラス
	PLUS = "+"
	// マイナス
	MINUS = "-"
	// アスタリスク
	ASTERISK = "*"
	// スラッシュ
	SLASH = "/"

	// 等号
	EQ = "="
	// 小なり
	LT = "<"
	// 大なり
	GT = ">"

	// 開き括弧
	L_PAREN = "("
	// 閉じ括弧
	R_PAREN = ")"
	// 開き角括弧
	L_BRACKET = "["
	// 閉じ角括弧
	R_BRACKET = "]"
)

// キーワード
const (
	// ダイスロール：加算ロールとD66ロール
	D_ROLL = "D_ROLL"
	// バラバラロール
	B_ROLL = "B_ROLL"
	// 個数振り足しロール
	R_ROLL = "R_ROLL"
	// 上方無限ロール
	U_ROLL = "U_ROLL"
	// シークレットロール
	SECRET = "SECRET_ROLL"
	// 計算
	CALC = "CALC"
	// ランダム選択
	CHOICE = "CHOICE"
)

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
