package basic_test

import (
	"testing"

	"github.com/raa0121/GoBCDice/pkg/dicebot/gamesystem/basic"
	dicebottesting "github.com/raa0121/GoBCDice/pkg/dicebot/testing"
)

func TestDiceBot(t *testing.T) {
	testDataFileBaseNames := []string{
		"calc.txt",
		"d_roll_expr.txt",
		"d_roll_comp.txt",
		"b_roll_list.txt",
		"b_roll_comp.txt",
		"r_roll_list.txt",
		"r_roll_comp.txt",
		"u_roll_expr.txt",
		"u_roll_comp.txt",
		"choice.txt",
		"secret_roll.txt",
	}

	testDataFiles := dicebottesting.JoinWithTestData(testDataFileBaseNames)
	dicebottesting.Run("DiceBot", t, testDataFiles...)
}

func TestBasic_GameID(t *testing.T) {
	expected := "DiceBot"
	actual := basic.New().GameID()

	if actual != expected {
		t.Fatalf("got: %q, want: %q", actual, expected)
	}
}

func TestBasic_GameName(t *testing.T) {
	expected := "ダイスボット (指定無し)"
	actual := basic.New().GameName()

	if actual != expected {
		t.Fatalf("got: %q, want: %q", actual, expected)
	}
}

func TestBasic_Usage(t *testing.T) {
	if len(basic.New().Usage()) <= 0 {
		t.Fatal("Usage() が空文字列")
	}
}
