package diceBot

import (
	"testing"
)

func TestSwordWorld_GameType(t *testing.T) {
	bot := NewSwordWorld()
	got := bot.GameName()
	expected := "SwordWorld"

	if got != expected {
		t.Errorf("bot.GameName = %s; want %s", got, expected)
	}
}
