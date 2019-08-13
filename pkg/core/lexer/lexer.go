/*
BCDiceコマンドの字句解析処理のパッケージ。
*/
package lexer

import (
	"github.com/raa0121/GoBCDice/pkg/core/token"
	"strings"
	"unicode"
)

// LexerState は字句解析器の状態の型。
type LexerStateType int

const (
	// 初期状態
	STATE_INIT LexerStateType = iota
	// 数式を読み取る状態
	STATE_EXPRESSION
	// ランダム選択の選択肢を読み取る状態
	STATE_CHOICE
	// 読み取りを終了した状態
	STATE_END
)

// 字句解析器を表す構造体。
type Lexer struct {
	// 現在の字句解析器の状態
	state LexerStateType
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
	if l.state == STATE_INIT {
		// 初期状態のときのみ、ランダム選択かどうかを判断する
		if tok, foundChoiceBegin := l.nextTokenOnInit(); foundChoiceBegin {
			l.state = STATE_CHOICE
			return tok
		} else {
			// ランダム選択ではないので、数式読み取り状態に遷移する
			l.state = STATE_EXPRESSION
		}
	}

	switch l.state {
	case STATE_EXPRESSION:
		return l.nextTokenOnExpresion()
	case STATE_CHOICE:
		return l.nextTokenOnChoice()
	default:
		{
			// 読み取り終了

			// トークンが何文字目で発見されたか
			// 利用者に示すものなので、1-indexed
			column := l.position + 1

			return newEOT(column)
		}
	}
}

// nextTokenOnInit は次のトークンを返す（初期状態）。
func (l *Lexer) nextTokenOnInit() (tok token.Token, foundChoiceBegin bool) {
	// トークンが何文字目で発見されたか
	// 利用者に示すものなので、1-indexed
	column := l.position + 1

	if l.ch == 'C' || l.ch == 'c' {
		if literal, ok := l.tryRead("HOICE["); ok {
			return token.Token{
				Type:    token.CHOICE_BEGIN,
				Literal: literal,
				Column:  column,
			}, true
		}
	}

	return token.Token{}, false
}

// nextTokenOnInit は次のトークンを返す（数式読み取り状態）。
func (l *Lexer) nextTokenOnExpresion() token.Token {
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
			tok = newEOT(column)
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

// nextTokenOnInit は次のトークンを返す（ランダム選択の選択肢読み取り状態）。
func (l *Lexer) nextTokenOnChoice() token.Token {
	l.skipSpaces()

	// トークンが何文字目で発見されたか
	// 利用者に示すものなので、1-indexed
	column := l.position + 1

	var tok token.Token
	switch l.ch {
	case ']':
		tok = newToken(token.CHOICE_END, l.ch, column)
		l.state = STATE_END
	case ',':
		tok = newToken(token.COMMA, l.ch, column)
	case 0:
		tok = newEOT(column)
	default:
		// tok.Column は必ず readItemInList() の前に設定する
		tok.Column = column

		tok.Type = token.STRING
		tok.Literal = l.readItemInList()

		return tok
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

// skipSpaces は空白を読み飛ばす。
func (l *Lexer) skipSpaces() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
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

// newEOT は新しいEOT（入力終端）トークンを返す。
func newEOT(column int) token.Token {
	return token.Token{
		Type:    token.EOT,
		Literal: "",
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

	for l.ch != 0 && test(l.ch) {
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

// readItemInList はリスト [...] の項目を読み込んで返す。
func (l *Lexer) readItemInList() string {
	return strings.TrimSpace(l.readCharsWhile(isNotSpecialCharInList))
}

// tryRead は文字列expectedの読み込みを試す。
// expected を読み込めた場合、literalでリテラルを、okでtrueを返す。
// 読み込めなかった場合、okでfalseを返す。
//
// 読み込み時は大文字小文字を無視する。
func (l *Lexer) tryRead(expected string) (literal string, ok bool) {
	expectedChars := []rune(expected)
	n := len(expectedChars)

	for i := 0; i < n; i++ {
		if unicode.ToUpper(l.peekChar(i+1)) != unicode.ToUpper(expectedChars[i]) {
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

// isNotSpecialCharInList はリスト [...] における特別な意味の文字でないかを返す。
func isNotSpecialCharInList(ch rune) bool {
	return ch != ',' && ch != ']'
}
