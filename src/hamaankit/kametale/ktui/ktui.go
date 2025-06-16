package ktui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	// "github.com/hamaa/vsNae/src/Font"
	// "github.com/hamaa/vsNae/src/BattleParts/turnInfo/actlib"
	// "github.com/hamaa/vsNae/src/Depicter/show"
	// "github.com/hamaa/vsNae/src/Image"
	// "github.com/hamaa/vsNae/src/fontlib"
)

const (
	Jap = iota
	Eng

	AllLang
)

var (
	Lang            int
	textboxFont     text.Face
	serifFont       text.Face
	englishAsteFont text.Face
)

func Setting(defaultLang int, textboxF, serifF, engAsteFont text.Face) {
	Lang = defaultLang
	textboxFont, serifFont, englishAsteFont = textboxF, serifF, engAsteFont
}

// var lang = 0
func UndertaleTextBox(screen *ebiten.Image, x1 float32, x2 float32, y1 float32, y2 float32) {
	w := float32(5)

	// Image.BoxDepict(screen, Image.StdOpTri(), x1, x2, y1, y2, true, 1, 0xffffff, 1, false)
	show.ShowPathStd(screen, show.BoxPath(x1, x2, y1, y2), true, 1, 0xffffff, 1, false)

	// Image.BoxDepict(screen, Image.StdOpTri(), x1+w, x2-w, y1+w, y2-w, true, 1, 0x000000, 1, false)
	show.ShowPathStd(screen, show.BoxPath(x1+w, x2-w, y1+w, y2-w), true, 1, 0x000000, 1, false)
	// Image.BoxDepict(screen, Image.StdOpTri(), x1+(w/2), x2-(w/2), y1+(w/2), y2-(w/2), false, w, 0xff0000, 1, false)
}

///////////////////////////////////////////////////////////////////////////
///////////選択欄//////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////

func UTfight(screen *ebiten.Image, EneName string, EneHP int, EneHPmax int) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang] //Font.BoxMess[lang]
	dyEn := LtuningEng(7)
	// Image.PrintDepStd(screen, fmt.Sprintf("＊ %s", EneName), textboxFont, 30+20+(26*1.5), 250+20, 0)
	show.PrintShowOrg(screen, fmt.Sprintf("＊ %s", EneName), textboxFont, 30+20+(26*1.5), 250+20+dyEn, 0)
	dy := float32(-5.0)
	// Image.BoxDepictStd(screen, 574-83-100, 574-83, 250+30+dy, 250+30+18+dy, true, 1, 0xb50006)
	// Image.BoxDepictStd(screen, 574-83-100, 574-83-100+(float32(EneHP)/float32(EneHPmax)*100), 250+30+dy, 250+30+18+dy, true, 1, 0x00ff00)
	show.ShowPathStd(screen, show.BoxPath(574-83-100, 574-83, 250+30+dy, 250+30+18+dy), true, 1, 0xb50006, 1, false)
	show.ShowPathStd(screen, show.BoxPath(574-83-100, 574-83-100+(float32(EneHP)/float32(EneHPmax)*100), 250+30+dy, 250+30+18+dy), true, 1, 0x00ff00, 1, false)

}

func UTmercy(screen *ebiten.Image, canRun bool) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang]
	dyEn := LtuningEng(7)
	switch Lang {
	case Eng:
		show.PrintShowOrg(screen, "＊", englishAsteFont, 30+20+(26*1.5), 250+20, 0)
		show.PrintShowOrg(screen, "＊ Spare", textboxFont, 30+20+(26*1.5)+dyEn*2, 250+20+dyEn, 0)
		if canRun {
			// Image.PrintDepStd(screen, "＊ にげる", textboxFont, 30+20+(26*1.5), 250+20+float64(26*1.5*1), 0)
			show.PrintShowOrg(screen, "＊", englishAsteFont, 30+20+(26*1.5), 250+20+float64(26*1.5*1), 0)
			show.PrintShowOrg(screen, "＊ Flee", textboxFont, 30+20+(26*1.5)+dyEn*2, 250+20+float64(26*1.5*1)+dyEn, 0)
		}
	default:
		show.PrintShowOrg(screen, "＊ にがす", textboxFont, 30+20+(26*1.5), 250+20+dyEn, 0)
		if canRun {
			// Image.PrintDepStd(screen, "＊ にげる", textboxFont, 30+20+(26*1.5), 250+20+float64(26*1.5*1), 0)
			show.PrintShowOrg(screen, "＊ にげる", textboxFont, 30+20+(26*1.5), 250+20+float64(26*1.5*1)+dyEn, 0)
		}
	}

}

func UTitems(screen *ebiten.Image, soulseetop int, itemslist [3]string, vunum int) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang] //Font.BoxMess[lang]
	dyEn := LtuningEng(7)
	for i := 0; i < vunum; i++ {
		// Image.PrintDepStd(screen, fmt.Sprintf("＊ %s", itemslist[i]), textboxFont, 30+20+(26*1.5), 250+20+float64(26*1.5*i), 0)
		show.PrintShowOrg(screen, "＊", englishAsteFont, 30+20+(26*1.5), 250+20+float64(26*1.5*i), 0)
		show.PrintShowOrg(screen, fmt.Sprintf("＊ %s", itemslist[i]), textboxFont, 30+20+(26*1.5)+dyEn*2, 250+20+float64(26*1.5*i)+dyEn, 0)
	}

}

func Uniform3(num int) int {
	var vunum = 3
	if num < 3 {
		vunum = num
	}
	return vunum
}

func UTactions(screen *ebiten.Image, actNames [6]string) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang] // Font.BoxMess[lang]
	dyEn := LtuningEng(7)
	for i := 0; i < 3; i++ {
		if actNames[2*i] != "" {
			// Image.PrintDepStd(screen, fmt.Sprintf("＊ %s", acts[2*i]), textboxFont, 30+20+(26*1.5), 250+20+float64(26*1.5*i), 0)
			show.PrintShowOrg(screen, "＊", englishAsteFont, 30+20+(26*1.5), 250+20+float64(26*1.5*i), 0)
			show.PrintShowOrg(screen, fmt.Sprintf("＊ %s", actNames[2*i]), textboxFont, 30+20+(26*1.5)+dyEn*2, 250+20+float64(26*1.5*i)+dyEn, 0)
		}
		if actNames[2*i+1] != "" {
			// Image.PrintDepStd(screen, fmt.Sprintf("＊ %s", acts[2*i+1]), textboxFont, 320+(26*1.5), 250+20+float64(26*1.5*i), 0)
			show.PrintShowOrg(screen, "＊", englishAsteFont, 320+(26*1.5), 250+20+float64(26*1.5*i), 0)
			show.PrintShowOrg(screen, fmt.Sprintf("＊ %s", actNames[2*i+1]), textboxFont, 320+(26*1.5)+dyEn*2, 250+20+float64(26*1.5*i)+dyEn, 0)
		}
	}
}

// /////////////////////////////////////////////////////////////////////////
// /////////セリフメッセージ////////////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////
func UtSerifBox(screen *ebiten.Image, x float64, y float64, wide bool, serif string) {

	// seriFont := fontlib.InTermFont[fontlib.Lang] // Font.SerifFont[0]
	if wide {
		// Image.TringleDepict(screen, Image.StdOpTri(), float32(x), float32(x), float32(x)-30, float32(y)+25, float32(y)+45, float32(y)+35, false, 3, 0x000000, 1, false)
		// Image.ArcBoxDepict(screen, Image.StdOpTri(), float32(x), float32(x)+206, float32(y), float32(y)+100, 20, false, 3, 0x000000, 1, false)
		// Image.TringleDepict(screen, Image.StdOpTri(), float32(x), float32(x), float32(x)-30, float32(y)+25, float32(y)+45, float32(y)+35, true, 1, 0xffffff, 1, false)
		// Image.ArcBoxDepict(screen, Image.StdOpTri(), float32(x), float32(x)+206, float32(y), float32(y)+100, 20, true, 1, 0xffffff, 1, false)

		show.ShowPathStd(screen, show.TringlePath(float32(x), float32(x), float32(x)-30, float32(y)+25, float32(y)+45, float32(y)+35), false, 3, 0x000000, 1, false)
		show.ShowPathStd(screen, show.ArcBoxPath(float32(x), float32(x)+206, float32(y), float32(y)+100, 20), false, 3, 0x000000, 1, false)
		show.ShowPathStd(screen, show.TringlePath(float32(x), float32(x), float32(x)-30, float32(y)+25, float32(y)+45, float32(y)+35), true, 1, 0xffffff, 1, false)
		show.ShowPathStd(screen, show.ArcBoxPath(float32(x), float32(x)+206, float32(y), float32(y)+100, 20), true, 1, 0xffffff, 1, false)

	} else { //arcsize15
		// Image.TringleDepict(screen, Image.StdOpTri(), float32(x), float32(x), float32(x)-12, float32(y)+45, float32(y)+45+15, float32(y)+45+(15/2), false, 3, 0x000000, 1, false)
		// Image.ArcBoxDepict(screen, Image.StdOpTri(), float32(x), float32(x)+87, float32(y), float32(y)+110, 15, false, 3, 0x000000, 1, false)
		// Image.TringleDepict(screen, Image.StdOpTri(), float32(x), float32(x), float32(x)-12, float32(y)+45, float32(y)+45+15, float32(y)+45+(15/2), true, 1, 0xffffff, 1, false)
		// Image.ArcBoxDepict(screen, Image.StdOpTri(), float32(x), float32(x)+87, float32(y), float32(y)+110, 15, true, 1, 0xffffff, 1, false)

		show.ShowPathStd(screen, show.TringlePath(float32(x), float32(x), float32(x)-12, float32(y)+45, float32(y)+45+15, float32(y)+45+(15/2)), false, 3, 0x000000, 1, false)
		show.ShowPathStd(screen, show.ArcBoxPath(float32(x), float32(x)+87, float32(y), float32(y)+110, 15), false, 3, 0x000000, 1, false)
		show.ShowPathStd(screen, show.TringlePath(float32(x), float32(x), float32(x)-12, float32(y)+45, float32(y)+45+15, float32(y)+45+(15/2)), true, 1, 0xffffff, 1, false)
		show.ShowPathStd(screen, show.ArcBoxPath(float32(x), float32(x)+87, float32(y), float32(y)+110, 15), true, 1, 0xffffff, 1, false)
	}
	// Image.PrintDepCol(screen, Image.StdOpTx(), string(serif), seriFont, x+9, y+10, 13*1.5, 0x000000, 1)
	show.PrintShowStd(screen, serif, serifFont, x+9, y+10, 13*1.5, 0x000000, 1)
}

///////////////////////////////////////////////////////////////////////////
///////////テキストメッセージ////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////

func UTstdText(screen *ebiten.Image, dx, dy float64, txmess string, ast [3]bool) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang] //Font.BoxMess[lang]

	for i := 0; i < 3; i++ {
		if ast[i] {
			// Image.PrintDep(screen, Image.StdOpTx(), "＊", textboxFont, 30+20, 250+20+float64(26*1.5*i), 0)

			show.PrintShowOrg(screen, "＊", englishAsteFont, 30+20, 250+20+float64(26*1.5*i), 0)
		}
	}
	// Image.PrintDep(screen, Image.StdOpTx(), string(txmess), textboxFont, 30+20+(26*1.5), 250+20, (26 * 1.5))
	// show.PrintShowStd(screen, string(txmess), textboxFont, 30+20+(26*1.5), 250+20, (26 * 1.5), 0xffffff, 1)
	dyEn := LtuningEng(7)
	show.PrintShowOrg(screen, txmess, textboxFont, 30+20+(26*1.5)+dx, 250+20+dyEn+dy, 26*1.5)
}

func LcheckEng() bool {
	return Lang == Eng
}

func LtuningEng(num float64) float64 {
	if LcheckEng() {
		return num
	}
	return 0
}

func Aste() string {
	if LcheckEng() {
		return "*"
	}
	return "＊"
}

func ShowUTstdMessage(screen *ebiten.Image, txmess string, ast [3]bool, talkingMode bool, character *ebiten.Image) {
	// textboxFont := fontlib.InBoxFont[fontlib.Lang] //Font.BoxMess[lang]
	// outward.BattleBox(screen, 30, 30+575, 250, 250+140)
	// for i := 0; i < 3; i++ {
	// 	if ast[i] {
	// 		Image.PrintDep(screen, Image.StdOpTx(), "＊", textboxFont, 30+20, 250+20+float64(26*1.5*i), 0)
	// 	}
	// }
	// Image.PrintDep(screen, Image.StdOpTx(), string(txmess), textboxFont, 30+20+(26*1.5), 250+20, (26 * 1.5))
	high := float32(330.0)
	shiftX := 0.0

	if talkingMode {
		high = 10
		if character != nil {
			shiftX = 105
		}
	}
	dyEn := LtuningEng(7)
	UndertaleTextBox(screen, 30, 30+575, high, high+140) //480
	for i := 0; i < 3; i++ {
		if ast[i] {
			// Image.PrintDep(screen, Image.StdOpTx(), "＊", textboxFont, shiftX+30+20, float64(high)+20+float64(26*1.5*i), 0)

			// if fontlib.LcheckEng() {
			// 	show.PrintShowOrg(screen, fontlib.Aste(), fontlib.EmergFont[fontlib.EnglishAsteFont], shiftX+30+20, float64(high)+20+float64(26*1.5*i), 0)
			// } else {
			// 	show.PrintShowOrg(screen, fontlib.Aste(), textboxFont, shiftX+30+20, float64(high)+20+float64(26*1.5*i), 0)
			// }
			show.PrintShowOrg(screen, "＊", englishAsteFont, shiftX+30+20, float64(high)+20+float64(26*1.5*i), 0)
		}
	}
	// Image.PrintDep(screen, Image.StdOpTx(), string(txmess), textboxFont, shiftX+30+20+(26*1.5), float64(high)+20, (26 * 1.5))
	show.PrintShowOrg(screen, txmess, textboxFont, shiftX+30+20+(26*1.5), float64(high)+20+dyEn, (26 * 1.5))

	if talkingMode {
		if character != nil {
			show.ShowImgStd(screen, character, 1, 97, 80, 0, false, 1, 1, 1)
		}
	}
}
