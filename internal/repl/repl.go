package repl

import (
	"bufio"
	"fmt"
	"github.com/raa0121/GoBCDice/internal/lexer"
	"github.com/raa0121/GoBCDice/internal/parser"
	"github.com/raa0121/GoBCDice/internal/token"
	"io"
	"regexp"
)

const (
	// REPLのプロンプト
	PROMPT = ">> "

	// トークン出力コマンド
	COMMAND_TOKEN = ".token "
)

// コマンドハンドラの型
// 返り値は、REPLを終了するならばtrue、続けるならばfalse
type commandHandler func(input string, out io.Writer)

var (
	commandHandlers = map[string]commandHandler{
		"token": printTokens,
		"ast":   printSExp,
	}

	commandRe = regexp.MustCompile("\\A\\.([a-z]+) (.+)")
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == ".q" || line == ".quit" {
			break
		}

		matches := commandRe.FindStringSubmatch(line)

		if matches == nil {
			printSExp(line, out)
			continue
		}

		command, ok := commandHandlers[matches[1]]
		if !ok {
			printSExp(line, out)
			continue
		}

		command(matches[2], out)
	}
}

func printTokens(input string, out io.Writer) {
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != token.EOT; tok = l.NextToken() {
		fmt.Fprintf(out, "%s\n", tok)
	}
}

func printSExp(input string, out io.Writer) {
	ast, err := parser.Parse(input)
	if err != nil {
		fmt.Fprintf(out, "%s\n", err)
		return
	}

	fmt.Fprintf(out, "%s\n", ast.SExp())
}
