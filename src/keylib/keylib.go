package keylib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/getkey"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
)

var KeyOff = false

var (
	Decide    *getkey.InpKey
	DecideJst int
	Cross     [hmutil.ALLDIR]*getkey.InpKey
	CrossJst  [hmutil.ALLDIR]int
	Cancel    *getkey.InpKey
	CancelJst int
	Menu      *getkey.InpKey
	MenuJst   int
)

func JstInpKey(k *getkey.InpKey, kjst *int) bool {
	return hmutil.JstBool(k.Pushed, kjst)
}

func JustDir(dir int) bool {
	return JstInpKey(Cross[dir], &CrossJst[dir])
}

func init() {
	Decide = getkey.MakeInpKey(ebiten.KeyZ, ebiten.KeyEnter)
	Cancel = getkey.MakeInpKey(ebiten.KeyX, ebiten.KeyShift)
	Menu = getkey.MakeInpKey(ebiten.KeyC, ebiten.KeyControl)

	Cross[hmutil.DOWN] = getkey.MakeInpKey(ebiten.KeyArrowDown, ebiten.KeyS)
	Cross[hmutil.UP] = getkey.MakeInpKey(ebiten.KeyArrowUp, ebiten.KeyW)
	Cross[hmutil.LEFT] = getkey.MakeInpKey(ebiten.KeyArrowLeft, ebiten.KeyA)
	Cross[hmutil.RIGHT] = getkey.MakeInpKey(ebiten.KeyArrowRight, ebiten.KeyD)
}

func LibAllKeyCheck() {
	if KeyOff {
		Decide.Off()
		Cancel.Off()
		Menu.Off()
		for i := range hmutil.ALLDIR {
			Cross[i].Off()
		}

		return
	}
	Decide.Check()
	Cancel.Check()
	Menu.Check()
	for i := range hmutil.ALLDIR {
		Cross[i].Check()
	}
}
