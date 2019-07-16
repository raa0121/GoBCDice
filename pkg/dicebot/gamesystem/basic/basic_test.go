package basic_test

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/bcdice"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/dicebot/testcase"
	"path/filepath"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	testDataPath := filepath.Join("testdata", "DiceBot.txt")
	testcases, loadErr := testcase.ParseFile(testDataPath, "DiceBot")
	if loadErr != nil {
		t.Fatalf("%s を読み込めません: %s", testDataPath, loadErr)
		return
	}

	for _, test := range testcases {
		t.Run(fmt.Sprintf("%d:%q", test.Index, test.Input[0]), func(t *testing.T) {
			f := feeder.NewQueue(test.Dice)
			b := bcdice.New(f)

			expectErr := (test.Output == "")

			result, commandErr := b.ExecuteCommand(test.Input[0])
			if commandErr != nil {
				if !expectErr {
					// 予期せぬエラー
					t.Fatalf("コマンド実行エラー: %s", commandErr)
				}

				// 予想通りエラーが発生した
				return
			}

			if expectErr {
				// エラーが発生するはずなのに発生しなかった
				t.Fatal("エラーが発生しませんでした")
				return
			}

			expected := test.Output
			actual := result.Message()

			if actual != expected {
				t.Errorf("got: %q, want: %q", actual, expected)
			}

			if !f.IsEmpty() {
				t.Error("ダイス残り: " + dice.FormatDice(f.Dice()))
			}
		})
	}
}
