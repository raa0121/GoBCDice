package lexer

import (
	"github.com/raa0121/GoBCDice/internal/token"
	"strings"
)

// 字句解析器を表す構造体
type Lexer struct {
	// 入力文字列
	input []rune
	// 現在の文字の位置
	position int
	// 次に読み込む文字の位置
	readPosition int
	// 現在検査中の文字
	ch rune
}

// Newは新しいLexerを構築して返す
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()

	return l
}

// NextTokenは次のトークンを返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '=':
		tok = newToken(token.EQ, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.L_PAREN, l.ch)
	case ')':
		tok = newToken(token.R_PAREN, l.ch)
	case '[':
		tok = newToken(token.L_BRACKET, l.ch)
	case ']':
		tok = newToken(token.R_BRACKET, l.ch)
	case 0:
		tok.Type = token.EOT
		tok.Literal = ""
	default:
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()

			return tok
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(strings.ToUpper(tok.Literal))

			return tok
		}
	}

	l.readChar()

	return tok
}

// inputStrはinputを文字列に変換して返す
func (l *Lexer) inputStr() string {
	return string(l.input)
}

// readCharは文字を読み込む
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 入力文字列の終端に達した
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// 状態更新：次の文字へ進む
	l.position = l.readPosition
	l.readPosition++
}

// newTokenは新しいトークンを返す
//
// tokenTypeにはトークンの種類を指定する
// chには文字を指定する。返り値のLiteralではstringに変換される。
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// charTestは、文字の種類が条件を満たしているか調べる関数の型
type charTest func(rune) bool

// readCharsWhileは、条件を満たしている間文字を読み込み続ける
//
// testには、文字の種類が条件を満たしているか調べる関数を指定する
func (l *Lexer) readCharsWhile(test charTest) string {
	position := l.position

	for test(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

// readIdentifierは識別子を読み込んで返す
func (l *Lexer) readIdentifier() string {
	return l.readCharsWhile(isLetter)
}

// readNumberは数値を読み込んで返す
func (l *Lexer) readNumber() string {
	return l.readCharsWhile(isDigit)
}

// isLetterはchがアルファベットかどうかを返す
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z')
}

// isDigitはchが数字かどうかを返す
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
