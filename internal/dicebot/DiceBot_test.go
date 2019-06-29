package dicebot

import (
	"testing"
)

func TestDiceBot_GameType(t *testing.T) {
	bot := NewDiceBot()
	got := bot.GameName()
	expected := "DiceBot"

	if got != expected {
		t.Errorf("bot.GameName = %s; want %s", got, expected)
	}
}
