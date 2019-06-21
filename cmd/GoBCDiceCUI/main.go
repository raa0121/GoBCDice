package main

import (
	"fmt"
	"github.com/raa0121/GoBCDice/diceBot"
)

func main() {
	sw := diceBot.SwordWorld{}
	fmt.Printf("%v\n", sw.HelpMessage())
}
