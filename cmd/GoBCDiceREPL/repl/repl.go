package repl

import (
	"bufio"
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/bcdice"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/lexer"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	"github.com/raa0121/GoBCDice/pkg/core/token"
	dicebotlist "github.com/raa0121/GoBCDice/pkg/dicebot/list"
	"github.com/raa0121/GoBCDice/pkg/dicebot/testcase"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	// REPLのプロンプト
	PROMPT = ">> "
	// 結果の初めに出力する文字列
	RESULT_HEADER = "=> "

	COMMAND_TOKEN          = "token"
	COMMAND_AST            = "ast"
	COMMAND_EVAL           = "eval"
	COMMAND_ROLL           = "roll"
	COMMAND_SET_DIE_FEEDER = "set-die-feeder"
	COMMAND_SET_DICE_QUEUE = "set-dice-queue"
	COMMAND_SET_GAME       = "set-game"
	COMMAND_LIST_GAMES     = "list-games"

	COMMAND_HELP = "help"
)

// コマンドハンドラの型。
// 返り値は、REPLを終了するならばtrue、続けるならばfalse。
type CommandHandler func(r *REPL, c *Command, input string)

// REPLコマンドを表す構造体。
type Command struct {
	// コマンド名
	Name string
	// 使用法
	Usage string
	// 解説
	Description string
	// コマンドハンドラ
	Handler CommandHandler
}

var (
	// 利用できるコマンド
	commands []Command
	// コマンド名とコマンドとの対応
	commandMap = map[string]*Command{}

	// コマンド実行を表す正規表現
	commandRe = regexp.MustCompile(`\A\.([-a-z]+)(?:\s+(.+))*`)
	// 末尾の空白を表す正規表現
	tailSpacesRe = regexp.MustCompile(`\s+\z`)
)

// REPLで使用するデータを格納する構造体。
type REPL struct {
	in         io.Reader
	out        io.Writer
	dieFeeder  feeder.DieFeeder
	diceRoller *roller.DiceRoller
	bcDice     *bcdice.BCDice
}

// init はパッケージを初期化する。
func init() {
	commands = []Command{
		{
			Name:        COMMAND_TOKEN,
			Usage:       "." + COMMAND_TOKEN + " BCDiceコマンド",
			Description: "入力されたBCDiceコマンドのトークンを出力します",
			Handler:     printTokens,
		},
		{
			Name:        COMMAND_AST,
			Usage:       "." + COMMAND_AST + " BCDiceコマンド",
			Description: "入力されたBCDiceコマンドのASTをS式の形で出力します",
			Handler:     printSExp,
		},
		{
			Name:        COMMAND_EVAL,
			Usage:       "." + COMMAND_EVAL + " BCDiceコマンド",
			Description: "入力されたBCDiceコマンドを評価します",
			Handler:     eval,
		},
		{
			Name:        COMMAND_ROLL,
			Usage:       "." + COMMAND_ROLL + " 振る数 面の数",
			Description: "ダイスロールを行い、出目を出力します",
			Handler:     rollDice,
		},
		{
			Name:        COMMAND_SET_GAME,
			Usage:       "." + COMMAND_SET_GAME + " ゲーム識別子",
			Description: "指定されたゲームシステムのダイスボットを使用するように設定します",
			Handler:     setGame,
		},
		{
			Name:        COMMAND_LIST_GAMES,
			Usage:       "." + COMMAND_LIST_GAMES,
			Description: "利用可能なゲームシステムの識別子の一覧を出力します",
			Handler:     listGames,
		},
		{
			Name:        COMMAND_SET_DIE_FEEDER,
			Usage:       "." + COMMAND_SET_DIE_FEEDER + " queue/mt",
			Description: "ダイスの供給方法を設定します - queue: 手動指定、mt: ランダム",
			Handler:     setDieFeeder,
		},
		{
			Name:        COMMAND_SET_DICE_QUEUE,
			Usage:       "." + COMMAND_SET_DICE_QUEUE + " [値/面数[, 値/面数]...]",
			Description: "ダイスロール時に取り出されるダイスの列を設定します",
			Handler:     setDiceQueue,
		},
		{
			Name:        COMMAND_HELP,
			Usage:       "." + COMMAND_HELP,
			Description: "利用できるコマンドの使用法と説明を出力します",
			Handler:     printHelp,
		},
	}

	for i, _ := range commands {
		c := &commands[i]
		commandMap[c.Name] = c
	}
}

// New は新しいREPLを構築し、返す。
//
// REPLは、inから入力された文字列をコマンドとして実行し、
// outにその結果を出力する。
func New(in io.Reader, out io.Writer) *REPL {
	f := feeder.NewMT19937WithSeedFromTime()
	r := roller.New(f)

	return &REPL{
		in:         in,
		out:        out,
		dieFeeder:  f,
		diceRoller: r,
		bcDice:     bcdice.New(f),
	}
}

// Start はREPLを開始する。
func (r *REPL) Start() {
	scanner := bufio.NewScanner(r.in)

	for {
		fmt.Printf("\n%s", PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == ".q" || line == ".quit" {
			break
		}

		matches := commandRe.FindStringSubmatch(line)
		if matches == nil {
			r.executeDefaultCommand(line)
			continue
		}

		commandName := matches[1]
		cmd, ok := commandMap[commandName]
		if !ok {
			r.executeDefaultCommand(line)
			continue
		}

		commandArgs := tailSpacesRe.ReplaceAllString(matches[2], "")
		cmd.Handler(r, cmd, commandArgs)
	}
}

// executeDefaultCommand は、inputを引数として既定のコマンドを実行する。
// コマンドが指定されていなかったとき、マッチしなかったときに使う。
func (r *REPL) executeDefaultCommand(input string) {
	eval(r, commandMap[COMMAND_EVAL], input)
}

// printCommandUsage は、コマンドcの使用法を出力する。
func (r *REPL) printCommandUsage(c *Command) {
	fmt.Fprintf(r.out, "使用法: %s\n", c.Usage)
}

// printTokens は、inputを字句解析し、得られたトークン列を出力する。
func printTokens(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	l := lexer.New(input)
	for tok := l.NextToken(); tok.Type != token.EOT; tok = l.NextToken() {
		fmt.Fprintln(r.out, tok)
	}
}

// printSExp は、inputを構文解析し、得られたASTをS式の形で出力する。
func printSExp(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	ast, err := parser.Parse(input)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	fmt.Fprintf(r.out, "%s%s\n", RESULT_HEADER, ast.SExp())
}

// eval はinputを構文解析して評価し、その結果を出力する。
// ダイスロールが行われた場合、その結果も出力する。
func eval(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	// ダイスロール結果を指定していた場合は、評価後にそれを復元する
	if r.dieFeeder.CanSpecifyDie() {
		ds := r.dieFeeder.(*feeder.Queue).Dice()
		defer func() {
			f := r.dieFeeder.(*feeder.Queue)
			f.Set(ds)
		}()
	}

	result, err := r.bcDice.ExecuteCommand(input)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	fmt.Fprintln(r.out, RESULT_HEADER+result.Message())
}

var rollDiceRe = regexp.MustCompile(`\A(\d+)\s+(\d+)\z`)

// rollDice はinputで指定されたダイスロールを行う。
// inputは、「振る数 面数」の形の文字列とする。
func rollDice(r *REPL, c *Command, input string) {
	matches := rollDiceRe.FindStringSubmatch(input)
	if matches == nil {
		r.printCommandUsage(c)
		return
	}

	// ダイスロール結果を指定していた場合は、評価後にそれを復元する
	if r.dieFeeder.CanSpecifyDie() {
		ds := r.dieFeeder.(*feeder.Queue).Dice()
		defer func() {
			f := r.dieFeeder.(*feeder.Queue)
			f.Set(ds)
		}()
	}

	num, _ := strconv.Atoi(matches[1])
	sides, _ := strconv.Atoi(matches[2])

	rolledDice, err := r.diceRoller.RollDice(num, sides)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	fmt.Fprintf(r.out, "%s%s\n", RESULT_HEADER, dice.FormatDice(rolledDice))
}

// setGame は指定されたゲームシステムのダイスボットを使用するように設定する。
func setGame(r *REPL, c *Command, input string) {
	if input == "" {
		// ゲームシステムが指定されていなかった場合、現在のゲーム識別子を出力する
		fmt.Fprintln(r.out, r.bcDice.DiceBot.GameID())
		return
	}

	err := r.bcDice.SetDiceBotByGameID(input)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	fmt.Fprintln(r.out, "OK")
}

// listGames は利用可能なゲームシステムの識別子の一覧を出力する。
func listGames(r *REPL, c *Command, input string) {
	gameIDs := dicebotlist.AvailableGameIDs(true)
	for _, gameID := range gameIDs {
		fmt.Fprintln(r.out, gameID)
	}
}

// setDieFeeder は、ダイス供給機を設定する。
// inputには以下を指定できる。
//
// * "queue": 出目を指定する。
// * "mt"   : ランダムな出目とする。
func setDieFeeder(r *REPL, c *Command, input string) {
	var f feeder.DieFeeder

	feederType := strings.ToLower(input)
	switch feederType {
	case "queue":
		f = feeder.NewEmptyQueue()
	case "mt":
		f = feeder.NewMT19937WithSeedFromTime()
	default:
		r.printCommandUsage(c)
		return
	}

	r.dieFeeder = f
	r.diceRoller = roller.New(f)
	r.bcDice.SetDieFeeder(f)

	fmt.Fprintln(r.out, "OK")
}

// setDiceQueue は、ダイスロールで取り出されるダイスの列を設定する。
// inputは "値/面数, 値/面数, ..." という形にする。
func setDiceQueue(r *REPL, c *Command, input string) {
	if !r.dieFeeder.CanSpecifyDie() {
		fmt.Fprintln(r.out, "現在のダイス供給方法では、取り出されるダイスの列を設定できません")
		return
	}

	ds, err := testcase.ParseDice(input)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	f := r.dieFeeder.(*feeder.Queue)
	f.Set(ds)

	fmt.Fprintln(r.out, "OK")
}

// printHelp は、利用できるコマンドの使用法と説明を出力する
func printHelp(r *REPL, _ *Command, _ string) {
	for _, c := range commands {
		fmt.Fprintf(r.out, "%s\n    %s\n", c.Usage, c.Description)
	}
}
