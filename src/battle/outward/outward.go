package outward

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/fontlib"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/load"
)

const (
	CmFight = iota
	CmAct
	CmItem
	CmMercy
	CmAll
)

var (
	meHPMax    int
	playerName string
	love       int
)

func UTui(screen *ebiten.Image, cmmode int, selmd bool, meHP int) {

	chcm := [CmAll]int{0, 0, 0, 0}

	if selmd {
		chcm[cmmode] = 1
	}

	c := CmFight
	showCommand(screen, c, chcm[c], placeFx, cmmy)
	c = CmAct
	showCommand(screen, c, chcm[c], placeAx, cmmy)
	c = CmItem
	showCommand(screen, c, chcm[c], placeIx, cmmy)
	c = CmMercy
	showCommand(screen, c, chcm[c], placeMx, cmmy)

	showDatas(screen, meHP)

}

const darken = 0xff - 0x88

func UtBox(screen *ebiten.Image, bxy *[4]float64) {

	BattleBox(screen, float32((*bxy)[0]), float32((*bxy)[1]), float32((*bxy)[2]), float32((*bxy)[3]))

}

const soulchipnum = 6

var deadcounter int
var scatterRad [soulchipnum]float64

func UtSouls(screen *ebiten.Image, selmode bool, damedmode bool, mex float64, mey float64, dead bool) {
	if dead {
		if deadcounter >= 120 {

			if deadcounter == 120 {
				sound.PlaySE(soul_nano, "wav")
				for i := 0; i < soulchipnum; i++ {
					rand.Seed(time.Now().UnixNano() * int64(1+i))
					wide := (math.Pi)
					scatterRad[i] = (float64(i) * math.Pi / 3) + wide - rand.Float64()*(wide/2)
				}
			}

			downt := deadcounter - 120
			for i := 0; i < soulchipnum; i++ {

				UtSoulsChip(screen, mex, mey, float64(downt), scatterRad[i])
			}

		} else if deadcounter >= 60 {
			if deadcounter == 60 {
				sound.PlaySE(soul_half, "wav")
			}
			// Image.DepictStd(screen, SoulDead, 1, mex-1, mey-1, 0, false)
			show.ShowImgStd(screen, SoulDead, 1, mex-1, mey-1, 0, false, 1, 1, 1)
		} else {
			// Image.DepictStd(screen, Soul, 1, mex, mey, 0, false)
			show.ShowImgStd(screen, Soul, 1, mex, mey, 0, false, 1, 1, 1)
		}

		deadcounter++
	} else if selmode {
		// Image.DepictFill(screen, Soul, Image.StdOpCol(), 1, mex, mey, 0, 0x4e0000 /*0x000000*/, 1, false, false, false)
		show.ShowImgFillStd(screen, Soul, 1, mex, mey, 0, false, 1, 1, 0x4e0000 /*0x000000*/, 1)

	} else {
		// Image.DepictStd(screen, Soul, 1, mex, mey, 0, false)
		show.ShowImgStd(screen, Soul, 1, mex, mey, 0, false, 1, 1, 1)
		if damedmode {
			// Image.DepictFill(screen, Soul, Image.StdOpCol(), 1, mex, mey, 0, 0x880000, 1, false, false, false)
			show.ShowImgFillStd(screen, Soul, 1, mex, mey, 0, false, 1, 1, 0x880000, 1)
		}
	}
}

func UtSoulsChip(screen *ebiten.Image, mex float64, mey float64, t float64, rad float64) {
	g := 9.8 / 50
	v0 := float64(5)
	show.ShowImgStd(screen, SoulMini, 1, mex+v0*t*math.Cos(rad), mey+(0.5*g*(t*t))+(v0*t*math.Sin(rad)), math.Pi*t/100, false, 1, 1, 1)
}

func InvTimeFlashing(InvT *int, spead int) bool {
	if *InvT > 0 {
		*InvT--
		return (*InvT/spead)%2 == 0
	}
	return false
}

// コマンドを選択するタマシイの位置(cmmodeでselecting判定も)
func CmmView(cmmode int, mex *float64, mey *float64) (selecting bool) {
	switch cmmode {
	case CmFight:
		*mex, *mey = placeFx-38, cmsouldef

	case CmAct:
		*mex, *mey = placeAx-38, cmsouldef

	case CmItem:
		*mex, *mey = placeIx-38, cmsouldef

	case CmMercy:
		*mex, *mey = placeMx-38, cmsouldef

	default:
		return false
	}
	return true
}

// 縦選択(foract==trueなら縦横選択)
func Selecting(mex *float64, mey *float64, num int, foract bool) {
	if foract {
		if num%2 == 0 {
			*mex, *mey = 30+35, 250+33+float64(26*1.5*int(num/2))
		} else {
			*mex, *mey = 335, 250+33+float64(26*1.5*int(num/2))
		}
	} else {

		*mex, *mey = 30+35, 250+33+float64(26*1.5*num)
	}
}

var InBarrageZone *ebiten.Image

// 戦闘箱描画
func BattleBox(screen *ebiten.Image, x1 float32, x2 float32, y1 float32, y2 float32) {
	w := float32(5)

	// Image.BoxDepict(screen, Image.StdOpTri(), x1, x2, y1, y2, true, 1, 0xffffff, 1, false)
	show.ShowPathStd(screen, show.BoxPath(x1, x2, y1, y2), true, 1, 0xffffff, 1, false)

	// Image.BoxDepict(InBarrageZone, Image.StdOpTri(), x1+w, x2-w, y1+w, y2-w, true, 1, 0x000000, 1, false)
	show.ShowPathStd(InBarrageZone, show.BoxPath(x1+w, x2-w, y1+w, y2-w), true, 1, 0x000000, 1, false)
	// Image.BoxDepict(screen, Image.StdOpTri(), x1+(w/2), x2-(w/2), y1+(w/2), y2-(w/2), false, w, 0xff0000, 1, false)
}

func InputIBZ(gX, gY int) {
	InBarrageZone = ebiten.NewImage(gX, gY)
}

// コマンド表示
func showCommand(screen *ebiten.Image, cmname int, mode int, x float64, y float64) {
	// Image.DepictStd(screen, Command[cmname][mode%2], 1, x, y, 0, false)
	show.ShowImgStd(screen, Command[cmname][mode%2], 1, x, y, 0, false, 1, 1, 1)
}

var HpLong = 0 //体力の長さ(orgin:1000 - 99)
// プレイヤーデータ表示
func showDatas(screen *ebiten.Image, meHP int) {

	// Image.BoxDepictStd(screen, 255, 255+(float32(meHPMax*111)/float32(meHPMax-HpLong)), 250+140+10, 250+140+10+22, true, 1, 0xb50006)
	// Image.BoxDepictStd(screen, 255, 255+(float32(meHP)*111)/float32(meHPMax-HpLong), 250+140+10, 250+140+10+22, true, 1, 0xffff00)

	show.ShowPathStd(screen, show.BoxPath(255, 255+(float32(meHPMax*111)/float32(meHPMax-HpLong)), 250+140+10, 250+140+10+22), true, 1, 0xb50006, 1, false)
	show.ShowPathStd(screen, show.BoxPath(255, 255+(float32(meHP)*111)/float32(meHPMax-HpLong), 250+140+10, 250+140+10+22), true, 1, 0xffff00, 1, false)

	// Image.PrintDefF(screen, fmt.Sprintf("%s", playerName), Font.Namefont, 35, 250+140+(56/2))
	show.PrintShowOrg(screen, playerName, fontlib.NameFontBT, 35, 250+140+(56/4), 10)

	// Image.PrintDefF(screen, fmt.Sprintf("LV %d", love), Font.NumsL, 120, 250+140+(56/2))
	show.PrintShowOrg(screen, fmt.Sprintf("LV %d", love), fontlib.NumsLove, 120, 250+140+(56/4), 10)

	// Image.PrintDefF(screen, "HP", Font.Hpkr, 230-10, 250+140+(50/2))
	show.PrintShowOrg(screen, "HP", fontlib.KRfont, 230-10, 250+140+(50/3), 10)

	// Image.PrintDefF(screen, fmt.Sprintf("%d ／ %d", meHP, meHPMax), Font.NumsL, 230+145+40, 250+140+(50/2))
	show.PrintShowOrg(screen, fmt.Sprintf("%d ／ %d", meHP, meHPMax), fontlib.NumsLove, 230+145+40, 250+140+(50/4), 10)
}

const (
	placeFx = 55 + 30            //Fightコマンド通常x座標
	placeAx = placeFx + 110 + 45 //Actコマンド通常x座標
	placeIx = placeAx + 110 + 50 //Itemコマンド通常x座標
	placeMx = placeIx + 110 + 45 //Mercyコマンド通常x座標

	cmmy = 250 + 140 + 43 + 20 //コマンド通常y座標

	cmsouldef = 250 + 140 + 43 + 20 //コマンド選択タマシイ通常y座標
)

var Command [CmAll][2]*ebiten.Image //コマンド[種類][off/on]
var Soul *ebiten.Image              //タマシイ
var SoulDead *ebiten.Image          //タマ死イ
var SoulMini *ebiten.Image          //タマシイだったもの

var (
	soul_half sound.SoundF
	soul_nano sound.SoundF
)

func init() {
	load.InitalPro[load.InOutward] = func(gx, gy float64) {
		InputIBZ(int(gx), int(gy))
		//Font.Initing()
		st := [CmAll]string{"fight", "act", "item", "mercy"}
		for i := 0; i < CmAll; i++ {
			// fmt.Printf("material/img/Nae_command_set/%s.png\n", st[i])
			show.LoadImage(&Command[i][0], fmt.Sprintf("material/img/Nae_command_set/%s.png", st[i]))
			show.LoadImage(&Command[i][1], fmt.Sprintf("material/img/Nae_command_set/%s2.png", st[i]))

		}
		show.LoadImage(&Soul, "material/img/soul.png")
		show.LoadImage(&SoulDead, "material/img/soul_dead.png")
		show.LoadImage(&SoulMini, "material/img/it_was_soul.png")
		sound.LoadSound(&soul_half, "material/sounds/se/snd_break1_c.wav")
		sound.LoadSound(&soul_nano, "material/sounds/se/snd_break2_c.wav")
	}

}

func InputMainMydata(hp int, name string, lv int) {
	meHPMax = hp
	playerName = name
	love = lv
	deadcounter = 0
}
