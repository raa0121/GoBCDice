package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/repl"
	"os"
)

func main() {
	fmt.Println("GoBCDice REPL")
	fmt.Println("\n* コマンドを入力すると AST (S 式) を出力します")
	fmt.Println("* \".token コマンド\" と入力するとトークンを出力します")
	fmt.Println("* \".q\" または \".quit\" と入力すると終了します")
	fmt.Println("")

	repl.Start(os.Stdin, os.Stdout)
}
