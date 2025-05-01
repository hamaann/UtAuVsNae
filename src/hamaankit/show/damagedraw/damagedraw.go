package damagedraw

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/avatar"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
)

const damageCounterMax = 100.0

var damageFont text.Face

func SetDamageFont(path string, size float64) {
	const stdDpi = 72.0

	show.GetNewFont(&damageFont, path, size, stdDpi)
}

type damageNum struct {
	damage  int
	counter int
	x, y    float64
	z       float64
}

var allViewingDamage []*damageNum

func DamageGo(thisDamage int, tx, ty, tz float64) {
	b := &damageNum{
		damage:  thisDamage,
		counter: damageCounterMax,
		x:       tx,
		y:       ty,
		z:       tz,
	}
	// fmt.Println(b.damage)
	for i := range allViewingDamage {
		if allViewingDamage[i] == nil {
			allViewingDamage[i] = b
			return
		}
	}
	allViewingDamage = append(allViewingDamage, b)
}

func DrawAllDamagesAvatar(order avatar.Integrity, screen *ebiten.Image) {
	for i := range allViewingDamage {
		if allViewingDamage[i] == nil {
			continue
		}
		dm := allViewingDamage[i]
		defZ := dm.z
		defZ -= 10
		// avatar.PutMultiObject(order, screen, func(screen *ebiten.Image, x0, y0, sizeb float64) {
		// 	show.PrintShowStd(screen, fmt.Sprintf("%d", dm.damage), damageFont, x0, y0, 1, 0xffffff, float32(dm.counter)/damageCounterMax)
		// }, 0, 0, dm.x, dm.y, defZ+float64(dm.counter*10)/damageCounterMax, 0, 1)
		// fmt.Printf("%d ", dm.damage)
		avatar.PutMultiPicture2D(order, screen, func(screen *ebiten.Image, x0, y0, sizeb float64) {
			// fmt.Printf("%d\n", dm.damage)
			// fmt.Println(0xff * float64(dm.counter) / damageCounterMax)
			show.PrintShowStd(screen, fmt.Sprintf("%d", dm.damage), damageFont, x0, y0, 1, 0xffffff, float32(0xff*dm.counter)/damageCounterMax)
		}, dm.x, dm.y-defZ-float64(dm.counter*10)/damageCounterMax, 1, false)
	}
	// fmt.Println()
}

func UpdateAllDamagesAvatar() {
	for i := range allViewingDamage {
		if allViewingDamage[i] == nil {
			continue
		}

		dm := allViewingDamage[i]
		dm.counter--
		//
		if dm.counter == 0 {
			allViewingDamage[i] = nil
		}

	}
	//
}
