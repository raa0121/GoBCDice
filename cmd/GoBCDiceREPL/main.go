package main

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceREPL/repl"
	"os"
)

func main() {
	out := colorable.NewColorableStdout()

	fmt.Fprintln(out, "\033[1mGoBCDice REPL\033[0m")
	fmt.Fprintln(out, "\n* BCDiceコマンドを入力すると、その評価結果を出力します")
	fmt.Fprintln(out, "* \".help\" と入力すると、利用できるコマンドの使用法と説明を出力します")
	fmt.Fprintln(out, "* \".q\" または \".quit\" と入力すると終了します")
	fmt.Fprintln(out, "")

	r := repl.New(os.Stdin, out)
	r.Start()
}
