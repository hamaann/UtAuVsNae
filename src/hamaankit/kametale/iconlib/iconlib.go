package iconlib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
)

func WhoTxtSt(name string) (who *whotxt) {
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
	Img   *ebiten.Image
	Voice *sound.SoundF
}

var (
	DefauC     *whotxt
	DefmanIcon *whotxt
	NaeIcon    *whotxt
)

func LoadUtAuSounds() {
	DefauC = LoadNillCharacterIcon()
	DefmanIcon = LoadCharacterIcon("material/sounds/se/SND_TXT1.wav", "material/img/sad_man.png")
	NaeIcon = LoadCharacterIcon("material/sounds/se/SND_TXTNae.wav", "material/img/Nae_face.png")
}

func LoadNillCharacterIcon() (wt *whotxt) {
	w := &whotxt{}
	w.Img = nil
	var v sound.SoundF
	sound.LoadSound(&v, "material/sounds/se/SND_TXT1.wav")
	w.Voice = &v
	return w
}

func LoadCharacterIcon(soundSouce, imageSouce string) (wt *whotxt) {
	w := &whotxt{}
	var v sound.SoundF
	sound.LoadSound(&v, soundSouce)
	w.Voice = &v
	show.LoadImage(&w.Img, imageSouce)
	return w
}
