package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/cmd/GoBCDiceREPL/repl"
	"os"
)

func main() {
	fmt.Println("\033[1mGoBCDice REPL\033[0m")
	fmt.Println("\n* BCDiceコマンドを入力すると、その評価結果を出力します")
	fmt.Println("* \".help\" と入力すると、利用できるコマンドの使用法と説明を出力します")
	fmt.Println("* \".q\" または \".quit\" と入力すると終了します")
	fmt.Println("")

	r := repl.New(os.Stdin, os.Stdout)
	r.Start()
}
