package bcdice

import (
	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
	dicebotlist "github.com/raa0121/GoBCDice/pkg/dicebot/list"
	"regexp"
)

// BCDiceの全体動作を統括する構造体。
type BCDice struct {
	// 設定されているダイスボット
	DiceBot    dicebot.DiceBot
	dieFeeder  feeder.DieFeeder
	diceRoller *roller.DiceRoller
}

// New は新しいBCDiceを構築する。
func New(f feeder.DieFeeder) *BCDice {
	b := &BCDice{}

	b.SetDieFeeder(f)
	b.SetDiceBotByGameID("DiceBot")

	return b
}

// SetDiceBotByGameID は、指定された識別子を持つゲームシステムのダイスボットを使用するよう設定する。
func (b *BCDice) SetDiceBotByGameID(gameID string) error {
	diceBotConstructor, err := dicebotlist.Find(gameID)
	if err != nil {
		return err
	}

	b.DiceBot = diceBotConstructor()

	return nil
}

// DieFeeder は設定されているダイス供給機を返す。
func (b *BCDice) DieFeeder() feeder.DieFeeder {
	return b.dieFeeder
}

// SetDieFeeder はダイス供給機を指定されたものに設定する。
// 合わせてダイスローラーも設定される。
func (b *BCDice) SetDieFeeder(f feeder.DieFeeder) {
	b.dieFeeder = f
	b.diceRoller = roller.New(f)
}

// 空白で区切られた入力文字列から最初の部分を取り出すための正規表現
var commandFirstPartRe = regexp.MustCompile(`\A([^\s]*)(\s.*)?`)

// ExecuteCommand は指定されたコマンドを実行する。
func (b *BCDice) ExecuteCommand(c string) (*command.Result, error) {
	separated := commandFirstPartRe.FindStringSubmatch(c)
	firstPart := separated[1]

	{
		result, err := b.ExecuteDiceBotCommand(firstPart)
		if err == nil {
			return result, nil
		}
	}

	// FIXME: CHOICEコマンドの実装の都合上、空白ありの状態で一度実行してみて、
	// 失敗したら最初の部分のみで実行するという、無駄の多い処理になっている。
	// これはPEG構文解析器を導入した場合に不要となる予定。
	{
		result, err := b.ExecuteBasicCommand(c)
		if err == nil {
			return result, nil
		}
	}
	{
		result, err := b.ExecuteBasicCommand(firstPart)
		if err == nil {
			return result, nil
		}

		return nil, err
	}
}

// ExecuteDiceBotCommand は設定されているダイスボットを使用して指定されたコマンドを実行する。
func (b *BCDice) ExecuteDiceBotCommand(c string) (*command.Result, error) {
	env := evaluator.NewEnvironment()
	ev := evaluator.NewEvaluator(b.diceRoller, env)

	result, err := b.DiceBot.ExecuteCommand(c, ev)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ExecuteBasicCommand はBCDiceの基本コマンドを実行する。
func (b *BCDice) ExecuteBasicCommand(c string) (*command.Result, error) {
	node, parseErr := parser.Parse(c)
	if parseErr != nil {
		return nil, parseErr
	}

	env := evaluator.NewEnvironment()
	ev := evaluator.NewEvaluator(b.diceRoller, env)

	return command.Execute(node, b.DiceBot.GameID(), ev)
}
