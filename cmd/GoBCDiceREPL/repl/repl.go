package repl

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/raa0121/GoBCDice/pkg/bcdice"
	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"github.com/raa0121/GoBCDice/pkg/core/dice/feeder"
	"github.com/raa0121/GoBCDice/pkg/core/dice/roller"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
	dicebotlist "github.com/raa0121/GoBCDice/pkg/dicebot/list"
	dicebottesting "github.com/raa0121/GoBCDice/pkg/dicebot/testing"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	// 書式設定をリセットするエスケープシーケンス
	ESC_RESET = "\033[0m"
	// 太字にするエスケープシーケンス
	ESC_BOLD = "\033[1m"
	// 文字色を赤色にするエスケープシーケンス
	ESC_RED = "\033[31m"
	// 文字色を黄色にするエスケープシーケンス
	ESC_YELLOW = "\033[33m"
	// 文字色をシアンにするエスケープシーケンス
	ESC_CYAN = "\033[36m"

	// REPLのプロンプト
	PROMPT = ESC_YELLOW + ">>" + ESC_RESET + " "
	// 結果の初めに出力する文字列
	RESULT_HEADER = ESC_CYAN + "=>" + ESC_RESET + " "

	COMMAND_AST            = "ast"
	COMMAND_EVAL           = "eval"
	COMMAND_ROLL           = "roll"
	COMMAND_SET_DIE_FEEDER = "set-die-feeder"
	COMMAND_SET_DICE_QUEUE = "set-dice-queue"
	COMMAND_SET_GAME       = "set-game"
	COMMAND_LIST_GAMES     = "list-games"

	COMMAND_HELP = "help"
	COMMAND_QUIT = "quit"
)

// コマンドハンドラの型。
// 返り値は、REPLを終了するならばtrue、続けるならばfalse。
type CommandHandler func(r *REPL, c *Command, input string)

// REPLコマンドを表す構造体。
type Command struct {
	// コマンド名
	Name string
	// 引数の説明
	ArgsDescription string
	// 解説
	Description string
	// コマンドハンドラ
	Handler CommandHandler
	// 自動補完の候補
	Completers []readline.PrefixCompleterInterface
}

// Usage はコマンドの使用方法の説明を返す。
func (c *Command) Usage() string {
	commandPart := "." + c.Name

	if c.ArgsDescription == "" {
		return commandPart
	}

	return commandPart + " " + c.ArgsDescription
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
	terminated bool
	dieFeeder  feeder.DieFeeder
	diceRoller *roller.DiceRoller
	bcDice     *bcdice.BCDice
	completer  *readline.PrefixCompleter
}

// init はパッケージを初期化する。
func init() {
	commands = []Command{
		{
			Name:            COMMAND_AST,
			ArgsDescription: "BCDiceコマンド",
			Description:     "入力されたBCDiceコマンドのASTをS式の形で出力します",
			Handler:         printSExp,
		},
		{
			Name:            COMMAND_EVAL,
			ArgsDescription: "BCDiceコマンド",
			Description:     "入力されたBCDiceコマンドを評価します",
			Handler:         eval,
		},
		{
			Name:            COMMAND_ROLL,
			ArgsDescription: "振る数 面の数",
			Description:     "ダイスロールを行い、出目を出力します",
			Handler:         rollDice,
		},
		{
			Name:            COMMAND_SET_GAME,
			ArgsDescription: "ゲーム識別子",
			Description:     "指定されたゲームシステムのダイスボットを使用するように設定します",
			Handler:         setGame,
		},
		{
			Name:        COMMAND_LIST_GAMES,
			Description: "利用可能なゲームシステムの識別子の一覧を出力します",
			Handler:     listGames,
		},
		{
			Name:            COMMAND_SET_DIE_FEEDER,
			ArgsDescription: "mt/queue",
			Description:     "ダイスの供給方法を設定します - mt: ランダム、queue: 手動指定",
			Handler:         setDieFeeder,
			Completers: []readline.PrefixCompleterInterface{
				readline.PcItem("mt"),
				readline.PcItem("queue"),
			},
		},
		{
			Name:            COMMAND_SET_DICE_QUEUE,
			ArgsDescription: "[値/面数[, 値/面数]...]",
			Description:     "ダイスロール時に取り出されるダイスの列を設定します",
			Handler:         setDiceQueue,
		},
		{
			Name:        COMMAND_HELP,
			Description: "利用できるコマンドの使用法と説明を出力します",
			Handler:     printHelp,
		},
		{
			Name:        COMMAND_QUIT,
			Description: "GoBCDiceREPLを終了します",
			Handler:     terminateREPL,
		},
	}

	for i, _ := range commands {
		c := &commands[i]
		commandMap[c.Name] = c
	}

	commandSetGame := commandMap[COMMAND_SET_GAME]
	for _, gameId := range dicebotlist.AvailableGameIDs(true) {
		commandSetGame.Completers = append(
			commandSetGame.Completers,
			readline.PcItem(gameId),
		)
	}
}

// New は新しいREPLを構築し、返す。
//
// REPLは、inから入力された文字列をコマンドとして実行し、
// outにその結果を出力する。
func New(in io.Reader, out io.Writer) *REPL {
	f := feeder.NewMT19937WithSeedFromTime()
	r := roller.New(f)

	completers := make([]readline.PrefixCompleterInterface, 0, len(commands))
	for _, c := range commands {
		completers = append(completers, readline.PcItem("."+c.Name, c.Completers...))
	}

	return &REPL{
		in:         in,
		out:        out,
		terminated: false,
		dieFeeder:  f,
		diceRoller: r,
		bcDice:     bcdice.New(f),
		completer:  readline.NewPrefixCompleter(completers...),
	}
}

// filterInput はreadlineでブロックする文字かどうかを判定する
func filterInput(r rune) (rune, bool) {
	switch r {
	// ^Z をブロックする
	// 現在は^Zを押すと動作がおかしくなるため
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

// Start はREPLを開始する。
func (r *REPL) Start() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:              PROMPT,
		HistoryFile:         "GoBCDiceREPL_history.txt",
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		FuncFilterInputRune: filterInput,
		AutoComplete:        r.completer,
	})
	if err != nil {
		r.printError(err)
		return
	}
	defer l.Close()

	r.printWelcomeMessage()

	for !r.terminated {
		line, readlineErr := l.Readline()

		switch readlineErr {
		case io.EOF:
			// ^D が押されたら修了する
			break
		case readline.ErrInterrupt:
			// ^C が押されたら次の読み込みに移る
			continue
		}

		line = strings.TrimSpace(line)

		// REPL終了の ".q" のみ特別扱い
		if line == ".q" {
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
	fmt.Fprintf(r.out, "使用法: %s\n", c.Usage())
}

// printWelcomeMessage は起動時の歓迎メッセージを出力する。
func (r *REPL) printWelcomeMessage() {
	fmt.Fprintln(r.out, ESC_BOLD+"GoBCDice REPL"+ESC_RESET)
	fmt.Fprintln(r.out, "\n* BCDiceコマンドを入力すると、その評価結果を出力します")
	fmt.Fprintln(r.out, "* \".help\" と入力すると、利用できるコマンドの使用法と説明を出力します")
	fmt.Fprintln(r.out, "* \".q\" または \".quit\" と入力すると終了します")
	fmt.Fprintln(r.out, "")
}

// printOK はコマンドの実行に成功した旨のメッセージを出力する。
func (r *REPL) printOK() {
	fmt.Fprintln(r.out, ESC_CYAN+"OK"+ESC_RESET)
}

// printError はエラーメッセージを強調して出力する。
func (r *REPL) printError(err error) {
	fmt.Fprintln(r.out, ESC_RED+err.Error()+ESC_RESET)
}

// printSExp は、inputを構文解析し、得られたASTをS式の形で出力する。
func printSExp(r *REPL, c *Command, input string) {
	if input == "" {
		r.printCommandUsage(c)
		return
	}

	parseResult, err := parser.Parse("REPL", []byte(input))
	if err != nil {
		r.printError(err)
		return
	}

	node := parseResult.(ast.Node)

	fmt.Fprintf(r.out, "%s%s\n", RESULT_HEADER, node.SExp())
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
		r.printError(err)
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
		r.printError(err)
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
		r.printError(err)
		return
	}

	r.printOK()
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

	r.printOK()
}

// setDiceQueue は、ダイスロールで取り出されるダイスの列を設定する。
// inputは "値/面数, 値/面数, ..." という形にする。
func setDiceQueue(r *REPL, c *Command, input string) {
	if !r.dieFeeder.CanSpecifyDie() {
		r.printError(fmt.Errorf("現在のダイス供給方法では、取り出されるダイスの列を設定できません"))
		return
	}

	ds, err := dicebottesting.ParseDice(input)
	if err != nil {
		r.printError(err)
		return
	}

	f := r.dieFeeder.(*feeder.Queue)
	f.Set(ds)

	r.printOK()
}

// printHelp は、利用できるコマンドの使用法と説明を出力する
func printHelp(r *REPL, _ *Command, _ string) {
	for _, c := range commands {
		fmt.Fprint(r.out, ESC_BOLD+"."+c.Name+ESC_RESET)
		if c.ArgsDescription != "" {
			fmt.Fprint(r.out, " "+c.ArgsDescription)
		}
		fmt.Fprintln(r.out, "")

		fmt.Fprintln(r.out, "    "+c.Description)
	}
}

// terminateREPL はREPLを終了させる
func terminateREPL(r *REPL, _ *Command, _ string) {
	r.terminated = true
}
