package diceBot

const EMPTY_PREFIXES_PATTERN = `(?i)(^|\s)(S)?()(\s|$)`

type SendMode int

const (
	onlyResult SendMode = iota
	resultWithExpression
	resultWithExpressionEachDice
)

type SortType int

const (
	noneDiceDisableSort SortType = iota
	sumDiceEnableSort
	eachDiceEnableSort
	bothDiceEnableSort
)

type SameDiceRerollCount int

const (
	noneReroll SameDiceRerollCount = iota
	sameAllReroll
	sameOver2Reroll
)

type SameDiceRerollType int

const (
	onlyCheckReroll SameDiceRerollType = iota
	onlyDamageReroll
	bothCheckAndDamageReroll
)

type D66Type int

const (
	noneD66 D66Type = iota
	noneSortD66
	sortASCD66
)

type FractionType string

const (
	fractionOmmit   = "ommit"
	fractionRountUp = "roundUp"
)

type DiceBot struct {
	defaultSendMode       SendMode
	sendMode              SendMode
	sortType              SortType
	sameDiceRerollCount   SameDiceRerollCount
	sameDiceRerollType    SameDiceRerollType
	d66Type               D66Type
	isPrintMaxDice        bool
	upperRollThreshold    int
	unlimitedRollDiceType int
	rerollNumber          int
	defaultSuccessTarget  string
	rerollLimitCount      int
	fractionType          FractionType
	gameType              string
	gameName              string
	prefixes              []string
}

func NewDiceBot() DiceBot {
	return DiceBot{
		defaultSendMode:     resultWithExpressionEachDice,
		sendMode:            resultWithExpressionEachDice,
		sortType:            noneDiceDisableSort,
		sameDiceRerollCount: noneReroll,
		sameDiceRerollType:  onlyCheckReroll,
		d66Type:             noneSortD66,
		rerollLimitCount:    10000,
		fractionType:        fractionOmmit,
		gameType:            "DiceBot",
	}
}

func (d *DiceBot) Info() (name string, gameType string, prefixes []string, info string) {
	name = d.GameName()
	gameType = d.GameType()
	prefixes = d.Prefixes()
	info = d.HelpMessage()
	return name, gameType, prefixes, info
}

func (d DiceBot) GameName() string {
	return d.gameType
}

func (d DiceBot) GameType() string {
	return d.gameType
}

func (d *DiceBot) SetGameType(gameType string) {
	d.gameType = gameType
}

func (d DiceBot) Prefixes() []string {
	return d.prefixes
}

func (d DiceBot) HelpMessage() string {
	return ""
}

func (d DiceBot) DiceCommand(command string, nick_e string) (string, bool) {
	if true == d.isGetOriginalMessage() {
	}
	return "1", false
}

func (d DiceBot) isGetOriginalMessage() bool {
	return false
}
