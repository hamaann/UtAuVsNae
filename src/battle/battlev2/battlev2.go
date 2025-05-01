package battlev2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/outward"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/load"
	"github.com/hamaa/UtAuVsNae/src/stduttexts"
)

type Phase func(first bool)

var (
	nowPhase Phase
	nowFirst bool
)

var MainStatusMessages = textdata.GetTextData("material/data/mainMess.csv")

var SoulX, SoulY float64

var (
	tymingStatus       string
	mainChosingCommand int
	defBox             = [4]float64{30, 30 + 575, 250, 250 + 140} //文字入れ箱  [x1,x2,y1,y2]
	Box                [4]float64                                 //箱  [x1,x2,y1,y2]
)

func BattleView(sc *ebiten.Image) {
	outward.UTui(sc, mainChosingCommand, tymingStatus == "mainSel", 100)
	outward.UtBox(sc, &Box)
	// ViewCharacter(screen, gx, gy) //Nae表示
	stduttexts.PrintUt(sc)
}

func MakeMainSel(statusMessage string) Phase {
	return func(first bool) {
		setTymingStatus("mainSel")
		if first {
			// utrolling.RollingShowingStart()
			stduttexts.StartTextdata((*MainStatusMessages)[statusMessage],
				stduttexts.Nomal, 0, 0, true)
		}
		stduttexts.Control(false, false) //keylib.JstInpKey(keylib.Decide, &keylib.DecideJst), keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst))
		tymingshift(&mainChosingCommand)
		outward.CmmView(mainChosingCommand, &SoulX, &SoulY)
		// stduttexts.PrintUt(sc)
	}
}

// 選択切り替え
func tymingshift(selectCommand *int) {
	switch MainDecide(selectCommand) {
	case outward.CmFight:
		ShiftPhase(fightSel)
	case outward.CmAct:
		ShiftPhase(actSel)
	case outward.CmItem:
		if items.HowItems() > 0 {
			ShiftPhase(itemSel)
		}
	case outward.CmMercy:
		ShiftPhase(mercySel)
	}
}

func MainDecide(com *int) (decide int) {
	// dw := keylib.JustDir(hmutil.DOWN)
	// 			up := keylib.JustDir(hmutil.UP)
	if hmutil.MenuSelecter(keylib.JustDir(hmutil.RIGHT), keylib.JustDir(hmutil.LEFT), com, 0, 4, 1, 0) {
		sound.PlaySE(SeSel, "wav")
	}
	if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
		sound.PlaySE(SeDeci, "wav")
		return *com
	}
	return outward.CmAll
}
func RunPhase() {
	nowPhase(nowFirst)
	nowFirst = false
}

func ShiftPhase(p Phase) {
	nowPhase = p
	nowFirst = true
}

func setTymingStatus(tyming string) {
	if tymingStatus == tyming {
		tymingStatus = tyming
	}
}

var (
	SeSel  sound.SoundF //選択音
	SeDeci sound.SoundF //決定音

	SeDam sound.SoundF
)

var (
	defVoice sound.SoundF
)

func init() {
	load.InitalPro[load.InBattle] = func(gx, gy float64) {
		sound.LoadSound(&SeSel, "material/sounds/se/snd_squeak.wav")
		sound.LoadSound(&SeDeci, "material/sounds/se/snd_select.wav")
		sound.LoadSound(&SeDam, "material/sounds/se/snd_hurt1_c.wav")
		sound.LoadSound(&defVoice, "material/sounds/se/SND_TXT1.wav")
		// show.LoadImage(&imgCharacterNae[0], "material/img/Nae/Nae0.png")
		// show.LoadImage(&imgCharacterNae[1], "material/img/Nae/Nae1.png")
		// show.LoadImage(&imgCharacterNae[2], "material/img/Nae/Nae2.png")
		// show.LoadImage(&imgCharacterNae[3], "material/img/Nae/Nae3.png")
		// show.LoadImage(&imgCharacterNae[4], "material/img/Nae/Nae4.png")

		// items.InitItem()
		// items.AddItem(items.BSPai)
		// items.AddItem(items.Candy)
		// items.AddItem(items.Hotcake)
		// items.AddItem(items.IceBall)
		// items.AddItem(items.BadMemory)
		// items.AddItem(items.Candy)
		// items.AddItem(items.BadMemory)
		// items.AddItem(items.BadMemory)

	}
}
