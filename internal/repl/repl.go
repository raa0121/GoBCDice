package repl

import (
	"bufio"
	"fmt"
	"github.com/raa0121/GoBCDice/internal/dicebot/testcase"
	"github.com/raa0121/GoBCDice/internal/die/feeder"
	"github.com/raa0121/GoBCDice/internal/die/roller"
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
	COMMAND_ROLL           = "roll"
	COMMAND_SET_DIE_FEEDER = "set-die-feeder"
	COMMAND_SET_DICE_QUEUE = "set-dice-queue"

	COMMAND_HELP = "help"
)

// コマンドハンドラの型
// 返り値は、REPLを終了するならばtrue、続けるならばfalse
type commandHandler func(r *REPL, c *Command, input string)

// REPLコマンドを表す構造体
type Command struct {
	// コマンド名
	Name string
	// 使用法
	Usage string
	// 解説
	Description string
	// コマンドハンドラ
	Handler commandHandler
}

var (
	commands   []Command
	commandMap = map[string]*Command{}

	commandRe    = regexp.MustCompile(`\A\.([-a-z]+)(?:\s+(.+))*`)
	tailSpacesRe = regexp.MustCompile(`\s+\z`)
)

type REPL struct {
	in         io.Reader
	out        io.Writer
	dieFeeder  feeder.DieFeeder
	diceRoller *roller.DiceRoller
}

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

func New(in io.Reader, out io.Writer) *REPL {
	f := feeder.NewMT19937WithSeedFromTime()

	return &REPL{
		in:         in,
		out:        out,
		dieFeeder:  f,
		diceRoller: roller.New(f),
	}
}

func (r *REPL) Start() {
	scanner := bufio.NewScanner(r.in)

	for {
		fmt.Print(PROMPT)

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
			printSExp(r, commandMap[COMMAND_AST], line)
			continue
		}

		commandName := matches[1]
		cmd, ok := commandMap[commandName]
		if !ok {
			printSExp(r, commandMap[COMMAND_AST], line)
			continue
		}

		commandArgs := tailSpacesRe.ReplaceAllString(matches[2], "")
		cmd.Handler(r, cmd, commandArgs)
	}
}

func (r *REPL) printCommandUsage(c *Command) {
	fmt.Fprintf(r.out, "使用法: %s\n", c.Usage)
}

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

var rollDiceRe = regexp.MustCompile(`\A(\d+)\s+(\d+)\z`)

func rollDice(r *REPL, c *Command, input string) {
	matches := rollDiceRe.FindStringSubmatch(input)
	if matches == nil {
		r.printCommandUsage(c)
		return
	}

	num, _ := strconv.Atoi(matches[1])
	sides, _ := strconv.Atoi(matches[2])

	dice, err := r.diceRoller.RollDice(num, sides)
	if err != nil {
		fmt.Fprintln(r.out, err)
		return
	}

	diceStrs := []string{}
	for _, d := range dice {
		diceStrs = append(diceStrs, fmt.Sprintf("%d/%d", d.Value, d.Sides))
	}

	fmt.Fprintf(r.out, "%s%s\n", RESULT_HEADER, strings.Join(diceStrs, ", "))
}

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
	f.Clear()
	f.Append(dice)

	fmt.Fprintln(r.out, "OK")
}

func printHelp(r *REPL, _ *Command, _ string) {
	for _, c := range commands {
		fmt.Fprintf(r.out, "%s\n    %s\n", c.Usage, c.Description)
	}
}
