package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/repl"
	"os"
)

func main() {
	fmt.Println("GoBCDice REPL")
	fmt.Println("\n* コマンドを入力すると、そのASTをS式の形で出力します")
	fmt.Println("* \".help\" と入力すると、利用できるコマンドの使用法と説明を出力します")
	fmt.Println("* \".q\" または \".quit\" と入力すると終了します")
	fmt.Println("")

	r := repl.New(os.Stdin, os.Stdout)
	r.Start()
}
