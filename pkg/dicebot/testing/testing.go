/*
ダイスボットのテストの共通処理のパッケージ。
*/
package testing

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/bcdice"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"testing"
)

// Run はダイスボットのテストを実行する。
//
// gameID: ゲーム識別子,
// testDataPath: テストデータファイルのパス,
// t: テストの状態管理。
func Run(gameID string, testDataPath string, t *testing.T) {
	testcases, loadErr := ParseFile(testDataPath, gameID)
	if loadErr != nil {
		t.Fatalf("%s を読み込めません: %s", testDataPath, loadErr)
		return
	}

	for _, test := range testcases {
		name := fmt.Sprintf(
			"%s-%d:%q[%s]",
			test.GameID,
			test.Index,
			test.Input[0],
			dice.FormatDiceWithoutSpaces(test.Dice),
		)
		t.Run(name, func(t *testing.T) {
			f := feeder.NewQueue(test.Dice)
			b := bcdice.New(f)

			// TODO: エラーが発生することの予想を明示できるようにする
			expectErr := (test.Output == "")

			// TODO: 入力文字列のすべての行のコマンドを実行する
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
