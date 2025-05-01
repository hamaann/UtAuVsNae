package attack

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/bulletattack/barrages"
	"github.com/hamaa/UtAuVsNae/src/battle/bulletattack/barrageturn"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/keylib"
)

func AttackSetUp(pattern string) (boxType int) {
	return barrageturn.InitDanmaku(pattern)
}

func CallAttacking() (until bool) {
	return barrageturn.CallDanmakuData()
}

func SoulAttaking(sx *float64, sy *float64, invT *int, Hp *int, IsMove bool) (damed bool) {
	//fmt.Println(barrages.Bu[barrages.TestBull].Coll(*sx, *sy, 240+85, 250+70, rrrad, 0))
	// if barrages.Bu[barrages.TestBull].Coll(*sx, *sy, 240+85, 250+70, rrrad, 0) {
	// 	*invT = barrages.Bu[barrages.TestBull].Invtime
	// 	*Hp -= barrages.Bu[barrages.TestBull].Damage
	// }

	// fmt.Println(barrageturn.Living)
	for i := range barrageturn.BulletNum {
		if barrageturn.Living[i] {
			n := barrageturn.BulletNum[i]
			if barrages.Bu[n].Coll(*sx, *sy, barrageturn.Bx[i], barrageturn.By[i], barrageturn.BulletRad[i], barrageturn.Bt[i]) &&
				colorModeRun(barrages.Bu[n].DecreeMode, IsMove) {
				*invT = barrages.Bu[n].Invtime
				*Hp -= barrages.Bu[n].Damage
				if barrages.Bu[n].WillDelete {
					barrageturn.Living[i] = false
				}
				break
			}
		}
	}

	return *invT > 0
}

func colorModeRun(modeColor int, moving bool) (damage bool) {
	switch modeColor {
	case barrages.ModeBlue:
		return moving
	case barrages.ModeOrenge:
		return !moving
	default:
		return true
	}
}

func ViewAttack(screen *ebiten.Image, inB bool) {
	// if inB == barrages.Bu[barrages.TestBull].InBox {
	// 	barrages.Bu[barrages.TestBull].Showing(screen, 0, 250+70, rrrad, 0)
	// 	rrrad += math.Pi / 100
	// }
	// bn += 10
	// bn %= 300
	for i := range barrageturn.BulletNum {
		if barrageturn.Living[i] {
			n := barrageturn.BulletNum[i]
			if inB == barrages.Bu[n].InBox {

				barrages.Bu[n].Showing(screen, barrageturn.Bx[i], barrageturn.By[i], barrageturn.BulletRad[i], barrageturn.Bt[i])
			}
		}
	}

}

func SoulMoving(sx *float64, sy *float64, runspead float64, box *[4]float64) (moved bool) {
	mvd := [4]bool{false, false, false, false}
	mvd[0] = GoToDir(sx, sy, runspead, hmutil.DOWN, box)
	mvd[1] = GoToDir(sx, sy, runspead, hmutil.UP, box)
	mvd[2] = GoToDir(sx, sy, runspead, hmutil.LEFT, box)
	mvd[3] = GoToDir(sx, sy, runspead, hmutil.RIGHT, box)
	return mvd[0] || mvd[1] || mvd[2] || mvd[3]
}

func GoToDir(sx *float64, sy *float64, speed float64, godir int, Colb *[4]float64) (moved bool) {
	mvd := false
	if keylib.JustDir(godir) {
		switch godir {
		case hmutil.DOWN:
			*sy += speed

		case hmutil.LEFT:
			*sx -= speed

		case hmutil.RIGHT:
			*sx += speed

		case hmutil.UP:
			*sy -= speed

		}
		mvd = true

	}
	if collingGoOut(sx, sy, speed, godir, *Colb, -10) {
		for i := 0; i < 4; i++ {
			if outGo(i, sx, sy, *Colb, -5) {
				collingGoOut(sx, sy, speed*2, i, *Colb, -5)
			}
		}
	}

	return mvd
}

func collingGoOut(mx *float64, my *float64, speed float64, i_pos int, Colb [4]float64, Add float64) (in bool) {

	if !(*mx < Colb[1]+Add && *mx > Colb[0]-Add && *my > Colb[2]-Add && *my < Colb[3]+Add) {
		switch i_pos {
		case hmutil.DOWN:
			*my -= speed

		case hmutil.LEFT:
			*mx += speed
		case hmutil.RIGHT:
			*mx -= speed
		case hmutil.UP:
			*my += speed
		}
		return true
	}

	return false

}

func outGo(i_pos int, mx *float64, my *float64, Colb [4]float64, Add float64) (out bool) {
	switch i_pos {
	case hmutil.DOWN:

		return Colb[3]+Add < *my

	case hmutil.LEFT:

		return Colb[0]-Add > *mx

	case hmutil.RIGHT:
		return Colb[1]+Add < *mx
	case hmutil.UP:

		return Colb[2]-Add > *my

	}
	return false
}

// func Initing(gx, gy float64) {
// 	barrages.Initing(gx, gy)
// }
