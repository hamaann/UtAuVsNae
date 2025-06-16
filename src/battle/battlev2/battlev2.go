package battlev2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/bulletattack/attack"
	"github.com/hamaa/UtAuVsNae/src/battle/geno"
	"github.com/hamaa/UtAuVsNae/src/battle/outward"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/iconlib"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/opbox"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/utrolling"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/item"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/load"
	"github.com/hamaa/UtAuVsNae/src/stduttexts"
)

var (
	myHP    int
	myHPmax int
	InvTime int //無敵時間
)

func RunPhase() {
	tymingStatus = ""
	if p := nowPhase(nowFirst, &tymingStatus); p != nil {
		ShiftPhase(p)
	} else {
		nowFirst = false
	}

	// fmt.Println(tymingStatus)
}

func ShiftPhase(p Phase) {
	nowPhase = p
	nowFirst = true
}

type Phase func(first bool, Status *string) (next Phase)

var (
	nowPhase Phase
	nowFirst bool
	nowStame string
)

var MainStatusMessages map[string][]*textdata.Uttxt

func init() {

	textdata.GetTextData(&MainStatusMessages, "material/data/mainMess.csv", -1)
}

var SoulX, SoulY float64

var (
	tymingStatus       string
	mainChosingCommand int
	defBox             = [4]float64{30, 30 + 575, 250, 250 + 140} //文字入れ箱  [x1,x2,y1,y2]
	Box                [4]float64                                 //箱  [x1,x2,y1,y2]
)

var StdChose int

func BTinit(name string) {
	InvTime = 0
	outward.InputMainMydata(myHPmax, name, 21)
	nowStame = "test1"
	ShiftPhase(MainSel)
	Box = opbox.PattBox[0]
}

func BattleView(sc *ebiten.Image, gx, gy float64) {
	// fmt.Println(tymingStatus)
	outward.UTui(sc, mainChosingCommand, tymingStatus == "main_S", myHP)
	outward.UtBox(sc, &Box)

	attack.ViewAttack(sc, true)
	show.ShowImgStd(sc, outward.InBarrageZone, 1, float64(outward.InBarrageZone.Bounds().Dx())/2,
		float64(outward.InBarrageZone.Bounds().Dy())/2, 0, false, 1, 1, 1)
	// ViewCharacter(sc, gx, gy) //Nae表示
	stduttexts.PrintUt(sc)

	if tymingStatus == "item_S" {
		inm := item.HowItems()
		// fmt.Println(inm, item.AllItemSet)
		list := [3]string{} //item.AllItemSet[item.ItemKey(reachTopItem+0)].Name, item.AllItemSet[item.ItemKey(reachTopItem+1)].Name, item.AllItemSet[item.ItemKey(reachTopItem+2)].Name}
		for i := range list {
			if int(item.ItemKey(reachTopItem+i)) < len(item.AllItemSet) {
				if item.AllItemSet[item.ItemKey(reachTopItem+i)] != nil {
					list[i] = item.AllItemSet[item.ItemKey(reachTopItem+i)].Name
				}
			}
		}

		ktui.UTitems(sc, reachTopItem, list, ktui.Uniform3(inm))
		ItemDot(sc, StdChose)
	}
	attack.ViewAttack(sc, false)
	geno.ViewSlashDam(sc, float32(gx), 50, 100, 150)

	outward.UtSouls(sc, tymingStatus == "main_S", outward.InvTimeFlashing(&InvTime, 10), SoulX, SoulY, false)
	outward.InBarrageZone.Clear()

}

func MainSel(first bool, Status *string) (next Phase) {
	*Status = "main_S"
	if first {
		// t := (*MainStatusMessages)[statusMessage]
		stduttexts.StartTextdata((MainStatusMessages)[nowStame],
			stduttexts.Nomal, 0, 0, true, true)
	}
	stduttexts.Control(false, false)
	p, t := tymingshift(&mainChosingCommand)
	if t {
		// fmt.Println("selected>>", mainChosingCommand)
		utrolling.RollingShowingEnd()
		return p
	}
	outward.CmmView(mainChosingCommand, &SoulX, &SoulY)
	return nil

}

var toMainJst int

func ItemSel(first bool, Status *string) (next Phase) {
	*Status = "item_S"
	if first {
		reachTopItem = 0
		StdChose = 0
	}
	ItemSelecter(&StdChose, item.HowItems())
	if keylib.JstInpKey(keylib.Cancel, &toMainJst) {
		return MainSel // mainSelに戻る(-際-に-必-要-な-こ-と)
	}
	if keylib.JstInpKey(keylib.Decide, &mainSelectJstSystem) {
		sound.PlaySE(SeDeci, "wav")
		// var nm int
		sayAns, nm := item.ItemEffeHp(StdChose, &myHP, 100)
		switch nm {
		case item.Hotcake:
			// sayAns = append(sayAns, &turnInfo.NaeCake[fontlib.Lang])
			// EnemyTelling = actlib.Hotcake //"hotcake"
		}

		// stduttexts.StartTextdata(sayAns,
		// 	stduttexts.Nomal, 0, 0, true, true)
		// nowStame = "test1"
		Anser = sayAns
		return AnserPhase //(sayAns[0].Text)
	}

	return nil

}

var Anser []*textdata.Uttxt

func AnserPhase(first bool, Status *string) (next Phase) {
	*Status = "ans_S"
	if first {
		stduttexts.StartTextdata(Anser,
			stduttexts.Nomal, 0, 0, true, true)
		SoulHide()
	}

	if !stduttexts.TextRollingNow {
		return MainSel
	}

	return nil
}

// 選択切り替え
func tymingshift(selectCommand *int) (next Phase, decided bool) {
	switch MainDecide(selectCommand) {
	case outward.CmFight:
		// ShiftPhase(fightSel)
	case outward.CmAct:
		// ShiftPhase(actSel)
	case outward.CmItem:
		if item.HowItems() > 0 {
			return ItemSel, true
		}
	case outward.CmMercy:
		// ShiftPhase(mercySel)
	}

	return nil, false
}

var mainSelectJstSystem int

func MainDecide(com *int) (decide int) {
	// dw := keylib.JustDir(hmutil.DOWN)
	// 			up := keylib.JustDir(hmutil.UP)
	if hmutil.MenuSelecter(keylib.JustDir(hmutil.RIGHT), keylib.JustDir(hmutil.LEFT), com, 0, 4, 1) {
		sound.PlaySE(SeSel, "wav")
	}
	if keylib.JstInpKey(keylib.Decide, &mainSelectJstSystem) {
		sound.PlaySE(SeDeci, "wav")
		return *com
	}
	return outward.CmAll
}

func setTymingStatus(tyming string) {
	if tymingStatus != tyming {
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

		myHP = 50
		myHPmax = 100
		sound.LoadSound(&SeSel, "material/sounds/se/snd_squeak.wav")
		sound.LoadSound(&SeDeci, "material/sounds/se/snd_select.wav")
		sound.LoadSound(&SeDam, "material/sounds/se/snd_hurt1_c.wav")
		sound.LoadSound(&defVoice, "material/sounds/se/SND_TXT1.wav")
		// show.LoadImage(&imgCharacterNae[0], "material/img/Nae/Nae0.png")
		// show.LoadImage(&imgCharacterNae[1], "material/img/Nae/Nae1.png")
		// show.LoadImage(&imgCharacterNae[2], "material/img/Nae/Nae2.png")
		// show.LoadImage(&imgCharacterNae[3], "material/img/Nae/Nae3.png")
		// show.LoadImage(&imgCharacterNae[4], "material/img/Nae/Nae4.png")

		iconlib.LoadUtAuSounds()

		item.InitItem()
		item.AddItem(item.BSPai)
		item.AddItem(item.Candy)
		item.AddItem(item.Hotcake)
		item.AddItem(item.IceBall)
		item.AddItem(item.BadMemory)
		item.AddItem(item.Candy)
		item.AddItem(item.BadMemory)
		item.AddItem(item.BadMemory)

	}
}

var reachTopItem int //一番上に表示されるitem欄番号
func ItemSelecter(sel *int, actNum int) {
	inm := item.HowItems()
	//ボタンが押された時
	if hmutil.MenuSelecter(keylib.JustDir(hmutil.DOWN), keylib.JustDir(hmutil.UP), sel, 0, inm, 1) {
		sound.PlaySE(SeSel, "wav")
		for *sel-reachTopItem > 2 { //一番上から2つ以上先を選択したなら
			reachTopItem++

		}
		for *sel-reachTopItem < 0 { //一番上より後ろを選択したなら
			reachTopItem--

		}

	}
	tp := *sel - reachTopItem
	outward.Selecting(&SoulX, &SoulY, tp, false)
}

const (
	idotstx    = float32(30 + 575 - 25.0)
	idotsty    = float32(250 + 30.0)
	idotmidy   = float32(250 + 75.0)
	idotdeltay = float32(10.0)
)

var itemdotMv = 0

func ItemDot(screen *ebiten.Image, Chose int) {
	mx := float32(item.HowItems())
	itemdotMv++

	if mx > 3 {
		addY := float32(((itemdotMv / 20) % 3) * 2)
		mxh := (mx / 2)
		if reachTopItem != 0 {
			// Image.TringleDepict(screen, Image.StdOpTri(), idotstx-7, idotstx+7, idotstx, -addY+4+idotmidy+(-1-mxh)*idotdeltay, -addY+float32(4+idotmidy+(-1-mxh)*idotdeltay), -addY+float32(-4+idotmidy+(-1-mxh)*idotdeltay), true, 1, 0xffffff, 1, false)
			show.ShowPathStd(screen, show.TringlePath(idotstx-7, idotstx+7, idotstx, -addY+4+idotmidy+(-1-mxh)*idotdeltay, -addY+float32(4+idotmidy+(-1-mxh)*idotdeltay), -addY+float32(-4+idotmidy+(-1-mxh)*idotdeltay)), true, 1, 0xffffff, 1, false)
		}
		for i := float32(0); i < mx; i++ {
			dtx, dty := float32(idotstx), float32(idotmidy+(i-mxh)*idotdeltay)

			sz := float32(2)
			if int(i) == Chose {
				sz = 5
			}

			// Image.BoxDepictStd(screen, dtx-sz, dtx+sz, dty-sz, dty+sz, true, 1, 0xffffff)
			show.ShowPathStd(screen, show.BoxPath(dtx-sz, dtx+sz, dty-sz, dty+sz), true, 1, 0xffffff, 1, false)

		}
		if reachTopItem+2 != int(mx)-1 {
			// Image.TringleDepict(screen, Image.StdOpTri(), idotstx-7, idotstx+7, idotstx, addY+float32(-4+idotmidy+mxh*idotdeltay), addY+float32(-4+idotmidy+mxh*idotdeltay), addY+float32(4+idotmidy+mxh*idotdeltay), true, 1, 0xffffff, 1, false)
			show.ShowPathStd(screen, show.TringlePath(idotstx-7, idotstx+7, idotstx, addY+float32(-4+idotmidy+mxh*idotdeltay), addY+float32(-4+idotmidy+mxh*idotdeltay), addY+float32(4+idotmidy+mxh*idotdeltay)), true, 1, 0xffffff, 1, false)

		}
	}
}

func SoulHide() {
	SoulX = -1000
	SoulY = -1000
}

// func mainSelBack(tyminger Phase) {
// 	if keylib.JstInpKey(keylib.Cancel, &toMainJst) {
// 		// startTyming[tyminger] = true
// 		// tyming = tyminger
// 		// ShiftPhase(tyminger)
// 		// turnInfo.CallMess(0)
// 	}
// }

// func decided(nextPhase Phase) bool {
// 	if  {
// 		ShiftPhase(nextPhase)
// 		return true
// 	} else {
// 		return false
// 	}
// }
