package stduttexts

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/utrolling"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
)

var (
	utextType int

	x, y          float64
	serifWideMode bool

	TextRollingNow bool

	decide, skip bool
)

const (
	Nomal = iota
	Boxt
	Serif
)

// var

func StartTextdata(roll []*textdata.Uttxt, txttype int, seriX, seriY float64, seriWide bool) {
	utrolling.RollingShowingStart(roll)
	TextRollingNow = true
	x, y = seriX, seriY
	serifWideMode = seriWide
}

func Control(dcd, skp bool) {
	decide, skip = dcd, skp
}

func PrintUt(sc *ebiten.Image) {
	w := whoTxtSt(utrolling.RollingWho())
	// d, c := keylib.JstInpKey(keylib.Decide, &keylib.DecideJst), keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst)
	txt, fin := utrolling.RollingShowingGo(decide, skip)
	switch utextType {
	case Nomal:
		ktui.UTstdText(sc, 0, 0, txt, utrolling.IsAsts())
	case Boxt:
		ktui.ShowUTstdMessage(sc, txt, utrolling.IsAsts(), false, w.img)
	case Serif:
		ktui.UtSerifBox(sc, x, y, serifWideMode, txt)
	}

	if fin && decide {
		TextRollingNow = false
	}

}

func whoTxtSt(name string) (who *whotxt) {
	switch name {
	case "Nae":
		return NaeIcon
	case "who":
		return DefmanIcon
	case "st":
		return DefauC
	default:
		return DefauC
	}
}

type whotxt struct {
	img   *ebiten.Image
	voice *sound.SoundF
}

var (
	DefauC     = LoadNillCharacterIcon()
	DefmanIcon = LoadCharacterIcon("material/sounds/se/SND_TXT1.wav", "material/img/sad_man.png")
	NaeIcon    = LoadCharacterIcon("material/sounds/se/SND_TXTNae.wav", "material/img/Nae_face.png")
)

func LoadNillCharacterIcon() (wt *whotxt) {
	w := &whotxt{}
	w.voice = nil
	show.LoadImage(&w.img, "material/sounds/se/SND_TXT1.wav")
	return w
}

func LoadCharacterIcon(soundSouce, imageSouce string) (wt *whotxt) {
	w := &whotxt{}
	sound.LoadSound(w.voice, soundSouce)
	show.LoadImage(&w.img, imageSouce)
	return w
}
