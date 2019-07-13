package repl

import (
	"bufio"
	"fmt"
	"github.com/raa0121/GoBCDice/internal/ast"
	"github.com/raa0121/GoBCDice/internal/command"
	"github.com/raa0121/GoBCDice/internal/dicebot/testcase"
	"github.com/raa0121/GoBCDice/internal/die"
	"github.com/raa0121/GoBCDice/internal/die/feeder"
	"github.com/raa0121/GoBCDice/internal/die/roller"
	"github.com/raa0121/GoBCDice/internal/evaluator"
	"github.com/raa0121/GoBCDice/internal/lexer"
	"github.com/raa0121/GoBCDice/internal/parser"
	"github.com/raa0121/GoBCDice/internal/token"
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
	// 結果の2行目以降でのインデント用の文字列
	RESULT_INDENT = "   "

	COMMAND_TOKEN          = "token"
	COMMAND_AST            = "ast"
	COMMAND_EVAL           = "eval"
	COMMAND_ROLL           = "roll"
	COMMAND_SET_DIE_FEEDER = "set-die-feeder"
	COMMAND_SET_DICE_QUEUE = "set-dice-queue"

	COMMAND_HELP = "help"
)

// コマンドハンドラの型
// 返り値は、REPLを終了するならばtrue、続けるならばfalse
type CommandHandler func(r *REPL, c *Command, input string)

// REPLコマンドを表す構造体
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

type REPL struct {
	in         io.Reader
	out        io.Writer
	dieFeeder  feeder.DieFeeder
	diceRoller *roller.DiceRoller
}

// initはパッケージを初期化する
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

// Newは新しいREPLを構築し、返す。
// REPLは、inから入力された文字列をコマンドとして実行し、
// outにその結果を出力する。
func New(in io.Reader, out io.Writer) *REPL {
	f := feeder.NewMT19937WithSeedFromTime()

	return &REPL{
		in:         in,
		out:        out,
		dieFeeder:  f,
		diceRoller: roller.New(f),
	}
}

// StartはREPLを開始する
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

// printCommandUsageは、コマンドcの使用法を出力する
func (r *REPL) printCommandUsage(c *Command) {
	fmt.Fprintf(r.out, "使用法: %s\n", c.Usage)
}

// printTokensは、inputを字句解析し、得られたトークン列を出力する
func printTokens(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	l := lexer.New(input)

	first := true
	for tok := l.NextToken(); tok.Type != token.EOT; tok = l.NextToken() {
		header := RESULT_INDENT
		if first {
			header = RESULT_HEADER
		}

		fmt.Fprintf(r.out, "%s%s\n", header, tok)

		first = false
	}
}

// printSExpは、inputを構文解析し、得られたASTをS式の形で出力する
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

// evalは、inputを構文解析して評価し、その結果を出力する。
// ダイスロールが行われた場合、その結果も出力する。
func eval(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	node, parseErr := parser.Parse(input)
	if parseErr != nil {
		fmt.Fprintf(r.out, "構文エラー: %s\n", parseErr)
		return
	}

	commandNode, nodeIsCommandNode := node.(ast.Command)
	if !nodeIsCommandNode {
		fmt.Fprintln(r.out, "実行可能なコマンドではありません")
		return
	}

	// ダイスロール結果を指定していた場合は、評価後にそれを復元する
	if r.dieFeeder.CanSpecifyDie() {
		dice := r.dieFeeder.(*feeder.Queue).Dice()
		defer func() {
			f := r.dieFeeder.(*feeder.Queue)
			f.Set(dice)
		}()
	}

	evaluator := evaluator.NewEvaluator(
		r.diceRoller, evaluator.NewEnvironment())

	result, execErr := command.Execute(commandNode, "DiceBot", evaluator)
	if execErr != nil {
		fmt.Fprintln(r.out, execErr)
		return
	}

	fmt.Fprintln(r.out, RESULT_HEADER+result.Message())
}

var rollDiceRe = regexp.MustCompile(`\A(\d+)\s+(\d+)\z`)

// rollDiceは、inputで指定されたダイスロールを行う。
// inputは、「振る数 面数」の形の文字列とする。
func rollDice(r *REPL, c *Command, input string) {
	matches := rollDiceRe.FindStringSubmatch(input)
	if matches == nil {
		r.printCommandUsage(c)
		return
	}

	// ダイスロール結果を指定していた場合は、評価後にそれを復元する
	if r.dieFeeder.CanSpecifyDie() {
		dice := r.dieFeeder.(*feeder.Queue).Dice()
		defer func() {
			f := r.dieFeeder.(*feeder.Queue)
			f.Set(dice)
		}()
	}

	num, _ := strconv.Atoi(matches[1])
	sides, _ := strconv.Atoi(matches[2])

	dice, err := r.diceRoller.RollDice(num, sides)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	fmt.Fprintf(r.out, "%s%s\n", RESULT_HEADER, die.FormatDice(dice))
}

// setDieFeederは、ダイス供給機を設定する。
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

	fmt.Fprintln(r.out, "OK")
}

// setDiceQueueは、ダイスロールで取り出されるダイスの列を設定する。
// inputは "値/面数, 値/面数, ..." という形にする。
func setDiceQueue(r *REPL, c *Command, input string) {
	if !r.dieFeeder.CanSpecifyDie() {
		fmt.Fprintln(r.out, "現在のダイス供給方法では、取り出されるダイスの列を設定できません")
		return
	}

	dice, err := testcase.ParseDice(input)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	f := r.dieFeeder.(*feeder.Queue)
	f.Set(dice)

	fmt.Fprintln(r.out, "OK")
}

// printHelpは、利用できるコマンドの使用法と説明を出力する
func printHelp(r *REPL, _ *Command, _ string) {
	for _, c := range commands {
		fmt.Fprintf(r.out, "%s\n    %s\n", c.Usage, c.Description)
	}
}
