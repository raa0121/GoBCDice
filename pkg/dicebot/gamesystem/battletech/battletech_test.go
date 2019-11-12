package battletech_test

import (
	"testing"

	"github.com/raa0121/GoBCDice/pkg/dicebot/gamesystem/battletech"
	dicebottesting "github.com/raa0121/GoBCDice/pkg/dicebot/testing"
)

func TestDiceBot(t *testing.T) {
	testDataFileBaseNames := []string{
		"CT.txt",
	}

	testDataFiles := dicebottesting.JoinWithTestData(testDataFileBaseNames)
	dicebottesting.Run("BattleTech", t, testDataFiles...)
}

func TestBattleTech_GameID(t *testing.T) {
	expected := "BattleTech"
	actual := battletech.New().GameID()

	if actual != expected {
		t.Fatalf("got: %q, want: %q", actual, expected)
	}
}

func TestBattleTech_GameName(t *testing.T) {
	expected := "バトルテック"
	actual := battletech.New().GameName()

	if actual != expected {
		t.Fatalf("got: %q, want: %q", actual, expected)
	}
}

func TestBattleTech_Usage(t *testing.T) {
	if len(battletech.New().Usage()) <= 0 {
		t.Fatal("Usage() が空文字列")
	}
}
