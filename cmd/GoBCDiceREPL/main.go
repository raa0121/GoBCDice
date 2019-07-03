package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/repl"
	"os"
)

func main() {
	fmt.Println("GoBCDice REPL")

	repl.Start(os.Stdin, os.Stdout)
}
