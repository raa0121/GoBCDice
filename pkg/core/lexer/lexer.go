/*
BCDiceコマンドの字句解析処理のパッケージ。
*/
package lexer

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
	"strings"
)

// 字句解析器を表す構造体。
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

// New は新しいLexerを構築して返す。
// inputには入力する文字列を指定する。
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()

	return l
}

// 1文字のトークンに対応するトークンの種類
var oneCharTokenType = map[rune]token.TokenType{
	'+': token.PLUS,
	'-': token.MINUS,
	'*': token.ASTERISK,
	'/': token.SLASH,
	'=': token.EQ,
	'(': token.L_PAREN,
	')': token.R_PAREN,
	'[': token.L_BRACKET,
	']': token.R_BRACKET,
}

// NextToken は次のトークンを返す。
func (l *Lexer) NextToken() token.Token {
	// トークンが何文字目で発見されたか
	// 利用者に示すものなので、1-indexed
	column := l.position + 1

	var tok token.Token

	if tokenType, ok := oneCharTokenType[l.ch]; ok {
		tok = newToken(tokenType, l.ch, column)
	} else {
		switch l.ch {
		case '.':
			if literal, ok := l.tryRead(".."); ok {
				tok.Type = token.DOTS
				tok.Literal = literal
				tok.Column = column

				return tok
			}

			tok = newToken(token.ILLEGAL, l.ch, column)
		case '<':
			if literal, ok := l.tryRead("="); ok {
				tok.Type = token.LTEQ
				tok.Literal = literal
				tok.Column = column

				return tok
			}

			if literal, ok := l.tryRead(">"); ok {
				tok.Type = token.DIAMOND
				tok.Literal = literal
				tok.Column = column

				return tok
			}

			tok = newToken(token.LT, l.ch, column)
		case '>':
			if literal, ok := l.tryRead("="); ok {
				tok.Type = token.GTEQ
				tok.Literal = literal
				tok.Column = column

				return tok
			}

			tok = newToken(token.GT, l.ch, column)
		case 0:
			tok.Type = token.EOT
			tok.Literal = ""
			tok.Column = column
		default:
			if isDigit(l.ch) {
				// tok.Column は必ず readNumber() の前に設定する
				tok.Column = column

				tok.Type = token.INT
				tok.Literal = l.readNumber()

				return tok
			}

			if isLetter(l.ch) {
				// tok.Column は必ず readIdentifier() の前に設定する
				tok.Column = column

				tok.Literal = l.readIdentifier()
				tok.Type = token.LookUpIdent(strings.ToUpper(tok.Literal))

				return tok
			}

			tok = newToken(token.ILLEGAL, l.ch, column)
		}
	}

	l.readChar()

	return tok
}

// inputStr はinputを文字列に変換して返す。
func (l *Lexer) inputStr() string {
	return string(l.input)
}

// readChar は文字を読み込む。
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

// peekChar は読み込む位置にある文字を返す。
// ただし、読み込み位置は次に進めない。
func (l *Lexer) peekChar(step int) rune {
	posToPeek := l.readPosition + step - 1

	if posToPeek >= len(l.input) {
		return 0
	}

	return l.input[posToPeek]
}

// setPosition は、読み込み位置をposに設定する
func (l *Lexer) setPosition(pos int) {
	l.position = pos
	l.readPosition = pos + 1
}

// newToken は新しいトークンを返す
//
// tokenTypeにはトークンの種類を指定する。
// chには文字を指定する。
// 返り値のLiteralではstringに変換される。
// columnにはトークンが何文字目で発見されたかを指定する。
func newToken(tokenType token.TokenType, ch rune, column int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Column:  column,
	}
}

// charTestは、文字の種類が条件を満たしているか調べる関数の型。
type charTest func(rune) bool

// readCharsWhile は、条件を満たしている間文字を読み込み続ける。
//
// testには、文字の種類が条件を満たしているか調べる関数を指定する。
func (l *Lexer) readCharsWhile(test charTest) string {
	position := l.position

	for test(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

// readIdentifier は識別子を読み込んで返す。
func (l *Lexer) readIdentifier() string {
	return l.readCharsWhile(isLetter)
}

// readNumber は数値を読み込んで返す。
func (l *Lexer) readNumber() string {
	return l.readCharsWhile(isDigit)
}

// tryRead は文字列expectedの読み込みを試す。
// expected を読み込めた場合、literalでリテラルを、okでtrueを返す。
// 読み込めなかった場合、okでfalseを返す。
func (l *Lexer) tryRead(expected string) (literal string, ok bool) {
	expectedChars := []rune(expected)
	n := len(expectedChars)

	for i := 0; i < n; i++ {
		if l.peekChar(i+1) != expectedChars[i] {
			return string(l.ch), false
		}
	}

	chars := make([]rune, n+1)
	for i := 0; i < n+1; i++ {
		chars[i] = l.ch
		l.readChar()
	}

	return string(chars), true
}

// isLetter はchがアルファベットかどうかを返す。
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z')
}

// isDigit はchが数字かどうかを返す。
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
