package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/dicebot"
)

func main() {
	sw := dicebot.SwordWorld{}
	fmt.Printf("%v\n", sw.HelpMessage())
}
