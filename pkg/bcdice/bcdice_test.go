package bcdice

import (
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"testing"
)

func TestDefaultDiceBot(t *testing.T) {
	f := feeder.NewEmptyQueue()
	b := New(f)

	expected := "DiceBot"
	actual := b.DiceBot.GameID()
	if actual != expected {
		t.Fatalf("got: %q, want: %q", actual, expected)
	}
}

func TestSetUnknownDiceBot(t *testing.T) {
	f := feeder.NewEmptyQueue()
	b := New(f)

	err := b.SetDiceBotByGameID("Unknown")
	if err == nil {
		t.Fatal("未知のダイスボットを設定できてしまった")
	}
}
