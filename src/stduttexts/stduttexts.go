package stduttexts

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/iconlib"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/utrolling"
	// "github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	// "github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
)

var (
	utextType int

	x, y              float64
	serifWideMode     bool
	boxtCharacterMode bool

	TextRollingNow bool

	decide, skip bool
)

const (
	Nomal = iota
	Boxt
	Serif
)

// var

func StartTextdata(roll []*textdata.Uttxt, txttype int, seriX, seriY float64, boxtChara, seriWide bool) {
	// fmt.Println(*roll[0])
	utrolling.RollingShowingStart(roll)
	TextRollingNow = true
	utextType = txttype
	boxtCharacterMode = boxtChara
	x, y = seriX, seriY
	serifWideMode = seriWide
}

func Control(dcd, skp bool) {
	decide, skip = dcd, skp
	// if decide {
	// 	fmt.Println("de")
	// }
}

func PrintUt(sc *ebiten.Image) {
	if TextRollingNow {
		w := iconlib.WhoTxtSt(utrolling.RollingWho())
		// d, c := keylib.JstInpKey(keylib.Decide, &keylib.DecideJst), keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst)

		txt, fin := utrolling.RollingShowingGo(decide, skip)
		switch utextType {
		case Nomal:
			ktui.UTstdText(sc, 0, 0, txt, utrolling.IsAsts())
		case Boxt:
			ktui.ShowUTstdMessage(sc, txt, utrolling.IsAsts(), boxtCharacterMode, w.Img)
		case Serif:
			ktui.UtSerifBox(sc, x, y, serifWideMode, txt)
		}

		if fin && decide {
			TextRollingNow = false
		}
	}

}
