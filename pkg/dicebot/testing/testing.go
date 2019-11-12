/*
ダイスボットのテストの共通処理のパッケージ。
*/
package testing

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/raa0121/GoBCDice/pkg/bcdice"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
)

// Run はダイスボットのテストを実行する。
//
// gameID: ゲーム識別子,
// t: テストの状態管理。
// testDataFiles: テストデータファイルのパス,
func Run(gameID string, t *testing.T, testDataFiles ...string) {
	testcases, loadErr := ParseFiles(testDataFiles, gameID)
	if loadErr != nil {
		t.Fatalf("テストデータファイルの読み込み失敗: %s", loadErr)
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

			{
				err := b.SetDiceBotByGameID(test.GameID)
				if err != nil {
					t.Fatalf("ダイスボット設定エラー: %s", err)
					return
				}
			}

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

			var actual string
			if result.IsSecret {
				actual = fmt.Sprintf("%s###secret dice###", result.Message())
			} else {
				actual = result.Message()
			}

			if actual != expected {
				t.Errorf("got: %q, want: %q", actual, expected)
			}

			if !f.IsEmpty() {
				t.Error("ダイス残り: " + dice.FormatDice(f.Dice()))
			}
		})
	}
}

// JoinWithTestData はbasenamesの各要素の先頭に "testdata/" を追加したスライスを返す。
func JoinWithTestData(basenames []string) []string {
	files := make([]string, 0, len(basenames))

	for _, b := range basenames {
		files = append(files, filepath.Join("testdata", b))
	}

	return files
}
