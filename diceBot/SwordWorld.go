package diceBot

type SwordWorld struct {
	DiceBot
	rating_table int
}

func NewSwordWorld() SwordWorld {
	sw := SwordWorld{
		DiceBot:      NewDiceBot(),
		rating_table: 0,
	}
	return sw
}

func (sw *SwordWorld) Dicebot() DiceBot {
	return sw.DiceBot
}

func (sw *SwordWorld) GameName() string {
	return "SwordWorld"
}

func (sw *SwordWorld) HelpMessage() string {
	return "・SW　レーティング表　　　　　(Kx[c]+m$f) (x:キー, c:クリティカル値, m:ボーナス, f:出目修正)"
}
