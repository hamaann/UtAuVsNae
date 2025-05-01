package battle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/bulletattack/attack"
	"github.com/hamaa/UtAuVsNae/src/battle/geno"
	"github.com/hamaa/UtAuVsNae/src/battle/outward"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/opbox"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/load"
)

const BGMvol = 1

var imgCharacterNae [5]*ebiten.Image

var defBox = [4]float64{30, 30 + 575, 250, 250 + 140} //文字入れ箱  [x1,x2,y1,y2]
var Box [4]float64                                    //箱  [x1,x2,y1,y2]
var sx, sy = 320.0, 320.0                             //タマシイ座標
var myHP = 100                                        //自分の体力
var enemyHPmax int
var enemyHP int
var ImAttaking bool

const (
	mainSel   = iota //コマンド選択画面
	fightSel         //攻撃対象選択
	actSel           //行動選択
	itemSel          //アイテム選択
	mercySel         //みのがす選択
	myAttack         //自分の攻撃
	ansDecide        //コマンド実行結果
	eneAttack        //敵攻撃
	turnEnd          //終わり→次のターンへ
	Youdead          //GAME OVER

	allTyming //すべての手順
)

var mainCNows int //選択中のコマンド
var tyming int    //現在のシーン
// var startTyming [allTyming]bool //true:最初
var Tyminged int        //前のシーン
var gameTurn int        //ターンカウント
var skipTxtMess = false //true:テキストメッセージスキップ
var MessEnd bool        //メッセージ終了チェック
var EneSpeaking bool    //敵が話し中か
var canGoRun bool       //にげあり？
var InvTime int         //無敵時間
var battleWait int

var (
	mainMess      turnInfo.BoxText //テキストボックスのメッセージ
	actlistNum    turnInfo.ActSet  //全行動リスト
	actlistLen    int              //全行動リスト数
	nowAttackPatt turnInfo.Hail    //攻撃パターン
	EnemyTelling  actlib.Topic     //敵セリフセット参照
)

var sayAns []*turnInfo.BoxText //決定に対する応え
var Chose = 0                  //選択一時保持
// var NNN = 0

func BttUpdate(gx, gy float64) {
	keylib.KeyOff = false
	// if hmutil.KboolMany(hmutil.Kok1, hmutil.Kok2) {
	// 	fmt.Println("jkdsc")
	// }
	// hmutil.PrintCanp(">>")
	if !sound.Check(pll) && tyming != Youdead {
		pll = sound.PlayBGM(NaeTheme, "wav")
	}

	if checkTYStart(tyming) { //初期動作
		turnInfo.CallMess(0)
		keylib.KeyOff = true
		switch tyming {
		case mainSel:
			// Image.MainScreenPermission(false)
			sound.SetVolPlayer(BGMvol, &pll)
			reachTopItem = 0
			Chose = 0
			turnInfo.CallMess(mainMess)
			geno.SetBar()
		// case myAttack:

		case ansDecide:

			messNum = 0
			serifNum = 0
		case eneAttack:
			battleWait = 30
			// Image.MainScreenPermission(true)
			sx, sy = 240+85, 250+70
			// sx = gx / 2
			opbox.ShiftBox(Box, opbox.PattBox[attack.AttackSetUp(string(nowAttackPatt))], 15)

		case Youdead:
			sound.Close(pll)
			// ktui.LetTerm(&EneSpeaking, 400, 100, "わかってるわ…\nあなたは　おうちが　こいしい\nのよね？　でも…", ktui.DefSp, 15, true, &defVoice, "wav")
		}
	}
	// hmutil.PrintCanp(">tymingsw<")

	// if NNN%8 == 0 {
	// 	Image.MakeGlich(20, 5, gy)
	// }
	//Image.AddScreen(0, 0, math.Pi/100, 0, 0)
	// Image.ChromaticaLeave(math.Abs(10*math.Sin(math.Pi*float64(NNN)/100)), [3]rune{'Y', 'C', 'M'})

	// NNN++

	switch tyming {
	case mainSel:

		tymingshift()
		outward.CmmView(mainCNows, &sx, &sy)
	case fightSel:

		StdSelecter(&Chose, 1)
		mainSelBack(mainSel)
		if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
			sound.PlaySE(SeDeci, "wav")
			tyming = myAttack
			ImAttaking = false

		}
	case actSel:

		ActSelecter(&Chose, actlistLen)
		mainSelBack(mainSel)
		// hmutil.PrintCanp(">act<")
		if decided() {
			sound.PlaySE(SeDeci, "wav")
			acom := turnInfo.ActList(actlistNum)[Chose]
			sayAns = turnInfo.ActAns(acom)
			switch acom {
			case actlib.Pet:
				EnemyTelling = actlib.GrossedOut //"ドン引き"
			}
		}
	case itemSel:
		ItemSelecter(&Chose, items.HowItems())
		mainSelBack(mainSel)
		if decided() {
			sound.PlaySE(SeDeci, "wav")
			var nm int
			sayAns, nm = items.ItemEffeHp(Chose, &myHP, 100)
			switch nm {
			case items.Hotcake:
				sayAns = append(sayAns, &turnInfo.NaeCake[fontlib.Lang])
				EnemyTelling = actlib.Hotcake //"hotcake"
			}
		}
	case mercySel:
		if canGoRun {
			StdSelecter(&Chose, 2)
		} else {
			StdSelecter(&Chose, 1)
		}
		mainSelBack(mainSel)

		switch Chose {
		case 0:
			if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
				tyming = ansDecide
				sayAns = []*turnInfo.BoxText{}
			}
		case 1:
		}
	case myAttack:
		hideSoul()
		// if {

		// tyming = ansDecide
		// sayAns = []*int{}

		// }
		if geno.StopBar((geno.GoingBar(6) || keylib.JstInpKey(keylib.Decide, &keylib.DecideJst)) && !ImAttaking, 600, 200, &enemyHP) {
			ImAttaking = true
		}
		if geno.RetDamaging() <= 0 && ImAttaking {
			tyming = ansDecide
			sayAns = []*turnInfo.BoxText{}
			break
		}
	case ansDecide:
		hideSoul()
		if speakingTM(sayAns) {
			tyming = eneAttack
			serif.CallSay(&EneSpeaking, 0, &naeNowFace)
			EneSpeaking = false

		}

	case eneAttack:

		if speakingSF(serif.SerifList(EnemyTelling)) {
			naeNowFace = 0
		}
		if battleWait > 0 {
			battleWait--
		}

		if !EneSpeaking && battleWait <= 0 {
			mv := attack.SoulMoving(&sx, &sy, 2, &Box)
			// fmt.Println(mv)
			if InvTime <= 0 {

				if attack.SoulAttaking(&sx, &sy, &InvTime, &myHP, mv) {
					sound.PlaySE(SeDam, "wav")
				}
			}

			if !attack.CallAttacking() {

				tyming = turnEnd
			}
		}

		if myHP <= 0 {
			tyming = Youdead
		}

	case turnEnd:
		opbox.ShiftBox(Box, opbox.PattBox[0], 15)
		gameTurn++
		tyming = mainSel
		mainMess, actlistNum, actlistLen, EnemyTelling, nowAttackPatt = turnInfo.AboutTurn(gameTurn)
		Chose = 0

	}
	// if time.Now().Second()%10 == 0 {
	// 	Image.MakeGlich(20, 10)
	// }

	opbox.ShiftBoxChenger(&Box)
}

var naeNowFace = serif.Face(0)

func ViewCharacter(screen *ebiten.Image, gx, gy float64) {
	show.ShowImgStd(screen, imgCharacterNae[naeNowFace], 1, gx/2, Box[2]-30, 0, true, 1, 1, 1)
}

func BttDraw(screen *ebiten.Image, gx, gy float64) {

	outward.UTui(screen, mainCNows, tyming == mainSel, myHP)
	outward.UtBox(screen, &Box)
	ViewCharacter(screen, gx, gy) //Nae表示
	// fmt.Println(gx/2, Box[2]-30-float64(imgCharacterNae.Bounds().Dy())/2)

	if eneAttack == tyming {
		attack.ViewAttack(outward.InBarrageZone, true)
	}
	// Image.DepictStd(screen, outward.InBarrageZone, 1, float64(outward.InBarrageZone.Bounds().Dx())/2,
	// 	float64(outward.InBarrageZone.Bounds().Dy())/2, 0, false)
	show.ShowImgStd(screen, outward.InBarrageZone, 1, float64(outward.InBarrageZone.Bounds().Dx())/2,
		float64(outward.InBarrageZone.Bounds().Dy())/2, 0, false, 1, 1, 1)
	// show.ShowImgStd(screen,outward.InBarrageZone, 1, float64(outward.InBarrageZone.Bounds().Dx())/2,
	// float64(outward.InBarrageZone.Bounds().Dy())/2, 0, false,1,1,1)

	// if gameTurn >= 1 {
	// 	Image.MainScreenPermission(true)
	// 	Image.SetGlitch(3, 100, float32(time.Now().Nanosecond())) //1000000
	// 	// Image.SetScreen(0, 0, 0, 1, -1)
	// 	// Image.ChromaticaLeave(3)
	// }

	//メッセージ////////////////////////////////////////////////////////////
	switch tyming {
	case mainSel:
		PutTextMess(screen)
	case fightSel:
		ktui.UTfight(screen, "Nae", enemyHP, enemyHPmax)
	case actSel:
		ktui.UTactions(screen, turnInfo.ActList(actlistNum))
	case itemSel:
		inm := items.HowItems()

		list := [3]string{items.ItemName(items.ItemList[reachTopItem+0]).Name, items.ItemName(items.ItemList[reachTopItem+1]).Name, items.ItemName(items.ItemList[reachTopItem+2]).Name}
		ktui.UTitems(screen, reachTopItem, list, ktui.Uniform3(inm))
		ItemDot(screen)
	case mercySel:
		ktui.UTmercy(screen, canGoRun)
	case myAttack:
		geno.AttackBar(screen, float32(gx))
		geno.ViewSlashDam(screen, float32(gx), enemyHP, enemyHPmax, 150)
	case ansDecide:
		PutTextMess(screen)

	case eneAttack:
		// Image.DepictCrush(screen, outward.Command[0][0], Image.StdOp(), 1, 240+85, 250+70, 0, 1, false, (sx-(240+85))/20, (sy-(250+70))/20)
		// fmt.Println((sx-(240+85))/20, (sy-(250+70))/20)
		attack.ViewAttack(screen, false)
	case Youdead:
		// Image.BoxDepictStd(screen, 0, float32(gx), 0, float32(gy), true, 1, 0x000000)
		show.ShowPathStd(screen, show.BoxPath(0, float32(gx), 0, float32(gy)), true, 1, 0x000000, 1, false)
	}

	if EneSpeaking {
		// fmt.Println("///")
		PutEneTerm(screen)
	}

	outward.UtSouls(screen, tyming == mainSel, outward.InvTimeFlashing(&InvTime, 10), sx, sy, tyming == Youdead)
	///////////////////////////////////////////////////////////////////////
	outward.InBarrageZone.Clear()
}

var NaeTheme sound.SoundF //BGM：The Hacker

var (
	SeSel  sound.SoundF //選択音
	SeDeci sound.SoundF //決定音

	SeDam sound.SoundF
)

var (
	defVoice sound.SoundF
)

var pll sound.Player //なんか音流してるやつ
// バトル初期設定
func BTinit(myName string) {

	gameTurn = 0
	mainMess, actlistNum, actlistLen, EnemyTelling, nowAttackPatt = turnInfo.AboutTurn(gameTurn)

	// for i := range startTyming {
	// 	startTyming[i] = true
	// }
	InvTime = 0
	Tyminged = turnEnd
	Box = opbox.PattBox[0]
	tyming = mainSel
	EneSpeaking = false

	// startTyming[mainSel] = true
	outward.InputMainMydata(100, myName, 21)
	// myHP = 1
	inputHPene(5000)

	sound.LoadSound(&NaeTheme, "material/sounds/bgm/The_hacker_short.wav")

	pll = sound.PlayBGM(NaeTheme, "wav")

	// items.UseItem(1)

}
func init() {
	load.InitalPro[load.InBattle] = func(gx, gy float64) {
		sound.LoadSound(&SeSel, "material/sounds/se/snd_squeak.wav")
		sound.LoadSound(&SeDeci, "material/sounds/se/snd_select.wav")
		sound.LoadSound(&SeDam, "material/sounds/se/snd_hurt1_c.wav")
		sound.LoadSound(&defVoice, "material/sounds/se/SND_TXT1.wav")
		show.LoadImage(&imgCharacterNae[0], "material/img/Nae/Nae0.png")
		show.LoadImage(&imgCharacterNae[1], "material/img/Nae/Nae1.png")
		show.LoadImage(&imgCharacterNae[2], "material/img/Nae/Nae2.png")
		show.LoadImage(&imgCharacterNae[3], "material/img/Nae/Nae3.png")
		show.LoadImage(&imgCharacterNae[4], "material/img/Nae/Nae4.png")

		// i := 0
		// switch i {
		// case 0:
		// ll
		// case 1:
		// 	// keyinp.Initing()

		// case 2:
		// 	//outward.Initing()
		// case 3:
		// 	// items.InitItem()
		// case 4:
		// 	// ktui.Initing()
		// case 5:
		// 	// serif.LoadSerifs()
		// case 6:
		// 	// turnInfo.LoadMesses()
		// case 7:
		// 	// attack.Initing(gx, gy)
		// case 8:
		// 	// outward.InputIBZ(int(gx), int(gy))
		// case 9:
		// 	// geno.Initing()
		// default:
		// 	//return -1
		// }

		//return i + 1
		items.InitItem()
		items.AddItem(items.BSPai)
		items.AddItem(items.Candy)
		items.AddItem(items.Hotcake)
		items.AddItem(items.IceBall)
		items.AddItem(items.BadMemory)
		items.AddItem(items.Candy)
		items.AddItem(items.BadMemory)
		items.AddItem(items.BadMemory)

	}
}

func decided() bool {
	if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
		tyming = ansDecide
		return true
	} else {
		return false
	}
}

// 選択切り替え
func tymingshift() {
	switch MainDecide(&mainCNows) {
	case commands.CmFight:
		tyming = fightSel
	case commands.CmAct:
		tyming = actSel
	case commands.CmItem:
		if items.HowItems() > 0 {
			tyming = itemSel
		}
	case commands.CmMercy:
		tyming = mercySel
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
	return commands.CmAll
}

// 行動の選択
func ActSelecter(pp *int, actNum int) {
	t1 := hmutil.MenuSelecter(keylib.JustDir(hmutil.RIGHT), keylib.JustDir(hmutil.LEFT), pp, 0, actNum, 1, 0)
	t2 := hmutil.MenuSelecter(keylib.JustDir(hmutil.DOWN), keylib.JustDir(hmutil.UP), pp, 0, actNum, 2, 1)
	if t1 || t2 {
		sound.PlaySE(SeSel, "wav")
	}
	outward.Selecting(&sx, &sy, *pp, true)
}

var reachTopItem int //一番上に表示されるitem欄番号
// アイテム選択
func ItemSelecter(pp *int, actNum int) {
	inm := items.HowItems()
	//ボタンが押された時
	if hmutil.MenuSelecter(keylib.JustDir(hmutil.DOWN), keylib.JustDir(hmutil.UP), pp, 0, inm, 1, 0) {
		sound.PlaySE(SeSel, "wav")
		for *pp-reachTopItem > 2 { //一番上から2つ以上先を選択したなら
			reachTopItem++

		}
		for *pp-reachTopItem < 0 { //一番上より後ろを選択したなら
			reachTopItem--

		}

	}
	tp := *pp - reachTopItem
	outward.Selecting(&sx, &sy, tp, false)
}

// 縦方向選択(上限項数３)
func StdSelecter(pp *int, num int) {
	if hmutil.MenuSelecter(keylib.JustDir(hmutil.DOWN), keylib.JustDir(hmutil.UP), pp, 0, num, 1, 0) {
		sound.PlaySE(SeSel, "wav")
	}

	outward.Selecting(&sx, &sy, *pp, false)
}

const (
	idotstx    = float32(30 + 575 - 25.0)
	idotsty    = float32(250 + 30.0)
	idotmidy   = float32(250 + 75.0)
	idotdeltay = float32(10.0)
)

var itemdotMv = 0

func ItemDot(screen *ebiten.Image) {
	mx := float32(items.HowItems())
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

// 敵セリフ表示
func PutEneTerm(screen *ebiten.Image) {
	Dialogue(ktui.PrintTerM, screen)

}

// メッセージ表示
func PutTextMess(screen *ebiten.Image) {
	Dialogue(ktui.PrintTxM, screen)

}

func Dialogue(diafunc func(screen *ebiten.Image, snd bool) (end bool), screen *ebiten.Image) {
	MessEnd = diafunc(screen, true)
	if skipTxtMess {
		for !diafunc(screen, false) {
			// fmt.Println(screen)
			// EneSpeaking = false
			MessEnd = true
			skipTxtMess = false
		}
	}
}

var putTextWait int

const ptwTime = 20

var messNum int

func speakingTM(messlist []*turnInfo.BoxText) (fin bool) {

	if len(messlist) <= 0 {
		return true
	}

	if messNum == 0 {
		turnInfo.CallMess(*messlist[messNum])
		messNum++
		putTextWait = ptwTime
	}
	if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) && (MessEnd && putTextWait <= 0) {
		if messNum < len(messlist) {
			turnInfo.CallMess(*messlist[messNum])
			messNum++
			putTextWait = ptwTime

		} else {
			return true
		}
	}

	if MessEnd && putTextWait > 0 {
		putTextWait--
	}

	if keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst) {
		if !MessEnd {
			skipTxtMess = true
		}
	}

	return false
}

var serifNum int

func speakingSF(seriflist []*int) (fin bool) {

	if len(seriflist) <= 0 {
		EneSpeaking = false
		return true
	}

	if serifNum == 0 {
		serif.CallSay(&EneSpeaking, *seriflist[serifNum], &naeNowFace)
		serifNum++
		putTextWait = ptwTime
	}
	if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) && (MessEnd && putTextWait <= 0) {

		if serifNum < len(seriflist) {
			serif.CallSay(&EneSpeaking, *seriflist[serifNum], &naeNowFace)
			serifNum++
			putTextWait = ptwTime

		} else {
			EneSpeaking = false
			return true
		}

	}

	if MessEnd && putTextWait > 0 {
		putTextWait--
		// fmt.Println(putTextWait, (MessEnd || putTextWait <= 0))
	}

	if keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst) {

		if !MessEnd {
			skipTxtMess = true
		}
	}

	return false
}

// tymingerの最初かを判定
func checkTYStart(tyminger int) bool {
	if tyminger != Tyminged {
		// startTyming[tyminger] = false
		Tyminged = tyminger
		return true
	}
	return false
}

// mainSelに戻る際に必要なこと
func mainSelBack(tyminger int) {
	if keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst) {
		// startTyming[tyminger] = true
		tyming = tyminger
		turnInfo.CallMess(0)
	}
}

// タマシイを表示しなくする
func hideSoul() {
	sx, sy = -1000, -1000
}

func inputHPene(hp int) {
	enemyHPmax = hp
	enemyHP = hp
}
