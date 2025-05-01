package barrages

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show/showkage"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/collision"
	"github.com/hamaa/UtAuVsNae/src/load"
)

var MidX, MidY float64

const (
	TestBull = iota
	TestBull2
	TestBull3B
	TestBull3O
	ErrorBull
	AllBull
)

const DefItime = 80

type Bullet struct {
	Showing    BulletView
	Coll       BulletColl
	Invtime    int
	Damage     int
	InBox      bool
	WillDelete bool
	DecreeMode int
}

const (
	Modedef = iota
	ModeBlue
	ModeOrenge
)

type (
	BulletView func(screen *ebiten.Image, x float64, y float64, rad float64, t int)
	BulletColl func(targX float64, targY float64, x float64, y float64, rad float64, t int) bool
)

var (
	testB     *ebiten.Image
	errMessJP *ebiten.Image
)

var Bu [AllBull]Bullet

func Makingbullet(gx, gy float64) {
	MidX, MidY = gx/2, gy/2

	show.LoadImage(&testB, "material/img/Nae_command_set/fight2.png")
	show.LoadImage(&errMessJP, "material/img/this_is_error_jp.png")

	s := func(screen *ebiten.Image, x float64, y float64, rad float64, t int) {
		showkage.ShowImgKage(screen, showkage.KageOptGlich(showkage.StdShowOptKage(errMessJP, 1, x, y, rad, 0.2, false, 1, 1),
			errMessJP, float32(time.Now().Nanosecond()), 3, 10, 0, 210), showkage.Glitch, errMessJP)
	}
	c := func(targX, targY, x, y, rad float64, t int) bool {
		return false
	}

	(&Bu[ErrorBull]).InputBullet(s, c, true, DefItime, 0, false, Modedef)

	s = func(screen *ebiten.Image, x float64, y float64, rad float64, t int) {
		//x, y = collision.BaseRad(x, y, 30, 0, rad)

		// if t%2 == 0 {

		//Image.Depict(screen, testB, StdOpInB(), 1, x, y, rad, 1, false, false, false)
		//show.ShowPath(screen, StdOpTriInB(true, false), show.LineBoxPath(x, y, 110/2, 40/2, rad), true, 1, 0xffffff, 1)
		// } else {
		// 	Image.DepictStdFill(screen, testB, 1, x, y, rad, 0xffffff, 1, false)
		// }
		//show.ShowPathStd(screen, show.LineBoxPath(x, y, 110/2, 40/2, rad), true, 1, 0xffffff, 1, false)

		ShowImgBull(screen, testB, x, y, 1, 30, 0, rad, 1, 1, 1)
		//ShowBoxBull(screen, x, y, 30, 0, 110/2, 40/2, rad, true, 1, 0xffffff, 1)

	}

	c = func(targX, targY, x, y, rad float64, t int) bool {
		return colBx(targX, targY, x, y, 30, 0, 110, 40, rad)
		// x, y = collision.BaseRad(x, y, 20, 0, rad)
		// return collision.BoxCol(targX, targY, x, y, 110, 40, rad)
		//return collision.OvalCol(targX, targY, x, y, 110, 40, rad)
	}

	(&Bu[TestBull]).InputBullet(s, c, true, DefItime, 25, false, Modedef)
	(&Bu[TestBull2]).InputBullet(s, c, false, DefItime, 25, true, Modedef)

	s = func(screen *ebiten.Image, x float64, y float64, rad float64, t int) {
		ShowBoxBull(screen, x, y, 30, 0, 150/2, 40/2, rad, true, 1, 0x00ffff, 1)
	}
	c = func(targX, targY, x, y, rad float64, t int) bool {
		return colBx(targX, targY, x, y, 30, 0, 150, 40, rad)

	}
	(&Bu[TestBull3B]).InputBullet(s, c, false, DefItime, 25, true, ModeBlue)

	s = func(screen *ebiten.Image, x float64, y float64, rad float64, t int) {
		ShowBoxBull(screen, x, y, 30, 0, 150/2, 40/2, rad, true, 1, 0xffaa00, 1)
	}
	(&Bu[TestBull3O]).InputBullet(s, c, false, DefItime, 25, true, ModeOrenge)

	// fmt.Println(">>", Bu[TestBull].InBox, Bu[TestBull].Invtime, Bu[TestBull].Damage)
	// ("%v", Bu[TestBull].Showing)

}

// func (b *Bullet) InputBulletMain(BulletView,BulletColl) {

// }

func (b *Bullet) InputBullet(show BulletView, coll BulletColl, inbox bool, invtime int, damage int, willdelete bool, decreeMode int) {
	b.Showing = show
	b.Coll = coll
	b.InBox = inbox

	b.Invtime = invtime
	b.Damage = damage
	b.WillDelete = willdelete
	b.DecreeMode = decreeMode
}

func StdOpInB() *ebiten.DrawImageOptions {
	ip := &ebiten.DrawImageOptions{} //Image.StdOp()
	ip.Blend = ebiten.BlendSourceAtop

	return ip
}

func StdOpTriInB(Fill bool, antiAlias bool) *ebiten.DrawTrianglesOptions {
	ip := show.StdShowOptTri(Fill, antiAlias) //&ebiten.DrawTrianglesOptions{}
	ip.Blend = ebiten.BlendSourceAtop

	return ip
}

func ShowImgBull(screen *ebiten.Image, Pic *ebiten.Image, baseX, baseY, size, x, y float64, rad, alpha float64, sizeX, sizeY float64) {
	baseX, baseY = collision.BaseRad(baseX, baseY, x, y, rad)
	opp := show.StdShowOpt(Pic, size, baseX, baseY, rad, alpha, false, sizeX, sizeY)
	opp.Blend = ebiten.BlendSourceAtop
	show.ShowImg(screen, Pic, opp)
}

func ShowBoxBull(screen *ebiten.Image, baseX, baseY float64, x, y, wide, high, rad float64, Fill bool, nonFillLineSize float32, Color int, alpha float64) {
	baseX, baseY = collision.BaseRad(baseX, baseY, x, y, rad)
	show.ShowPath(screen, StdOpTriInB(Fill, false), show.LineBoxPath(baseX, baseY, wide, high, rad), Fill, nonFillLineSize, Color, alpha)
}

func colBx(targX, targY, x, y, addX, addY, wide, high, rad float64) bool {
	x, y = collision.BaseRad(x, y, addX, addY, rad)
	return collision.BoxCol(targX, targY, x, y, wide, high, rad)
}

func colOv(targX, targY, x, y, addX, addY, wide, high, rad float64) bool {
	x, y = collision.BaseRad(x, y, addX, addY, rad)
	return collision.OvalCol(targX, targY, x, y, wide, high, rad)
}

func init() {

	load.InitalPro[load.InBarrages] = func(gx, gy float64) {
		Makingbullet(gx, gy)
	}
}
