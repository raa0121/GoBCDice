package repl

import (
	"bufio"
	"fmt"
	"github.com/raa0121/GoBCDice/internal/lexer"
	"github.com/raa0121/GoBCDice/internal/token"
	"io"
)

// REPLのプロンプト
const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOT; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
