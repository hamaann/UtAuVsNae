package barrageturn

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hamaa/UtAuVsNae/src/battle/bulletattack/barrages"
	// "golang.org/x/exp/rand"
)

var (
	BulletNum []int
	BulletRad []float64
	Bx, By    []float64
	Bt        []int
	Living    []bool
)

var nowAttPatt = ""
var counter = 0

func InitDanmaku(pattern string) (boxType int) {
	nowAttPatt = pattern
	counter = 0
	switch pattern {
	case "test3":
		return 0
	default:
		return 1
	}
}

func CallDanmakuData() (until bool) {

	counter++ //最初のループ：counter==1

	switch nowAttPatt {
	case "null":
		return false
	case "test1":
		if readyB(2) {
			setB(0, barrages.TestBull2, -200, float64(easyRand(225, 250+140, 924)), 0)
			setB(1, barrages.TestBull2, 200+barrages.MidX*2, float64(easyRand(225, 250+140, 925)), 0)
		}
		// if counter > 100 {
		shiftB(0, 8, modeX)
		shiftB(1, -8, modeX)
		if returnB(0, modeX) > 200+barrages.MidX*2 {
			setBm(0, 0, float64(easyRand(225, 250+140, 924)), 0)
			setBm(1, barrages.MidX*2, float64(easyRand(225, 250+140, 925)), 0)
		}
		// }
		return finaleB(500, false)
	case "test2":

		if readyB(1) {
			setB(0, barrages.TestBull, 240, closingMidle(225, 250+140), math.Pi/2)
			// Image.MainScreenPermission(true)
		}
		// if counter > 100 {
		shiftB(0, math.Pi/50, modeRad)
		shiftB(0, 0.5, modeX)
		// Image.ChromaticaLeave(math.Abs(50 * math.Sin(math.Pi*float64(counter)/100)))

		// }
		return finaleB(500, true)
	case "test3":
		if readyB(2) {
			setB(0, barrages.TestBull3B, -200, 250+140-50, -math.Pi/2)
			setB(1, barrages.TestBull3O, 200+barrages.MidX*2, 250+140-50, -math.Pi/2)
		}
		// if counter > 100 {
		shiftB(0, 4, modeX)
		shiftB(1, -4, modeX)
		if returnB(0, modeX) > 200+barrages.MidX*2 {
			setBm(0, 0, 250+140-50, -math.Pi/2)
			setBm(1, barrages.MidX*2, 250+140-50, -math.Pi/2)
		}

		return finaleB(1000, false)
	default:
		if readyB(1) {
			setB(0, barrages.ErrorBull, barrages.MidX, 250+140-80, 0)
		}
		return true
	}

	return true
}

// 弾幕のx,y,rad,tをどう動かすか
func readyB(howBullet int) bool {
	if counter == 1 {
		MakeBullet(howBullet)
	}

	return counter == 1
}

const (
	modeX = iota
	modeY
	modeRad
	modeT
)

func shiftB(slnum int, move float64, modeElement int) {
	switch modeElement {
	case modeX:
		shiftBall(slnum, move, 0, 0, 0)
	case modeY:
		shiftBall(slnum, 0, move, 0, 0)
	case modeRad:
		shiftBall(slnum, 0, 0, move, 0)
	case modeT:
		shiftBall(slnum, 0, 0, 0, int(move))
	}
}

func shiftBall(slnum int, dx, dy, drad float64, dt int) {
	BulletRad[slnum] += drad
	Bx[slnum] += dx
	By[slnum] += dy

	Bt[slnum] += dt
}

func returnB(slnum int, modeElement int) float64 {
	switch modeElement {
	case modeX:
		return Bx[slnum]
	case modeY:
		return By[slnum]
	case modeRad:
		return BulletRad[slnum]
	case modeT:
		return float64(Bt[slnum])
	}

	fmt.Println("正しくないエレメントが与えられました。")
	log.Fatal(modeElement)
	return 0
}

// ｔ初期化(t = 0)
func setB(slnum int, number int, x float64, y float64, rad float64) {
	setBwitnT(slnum, number, x, y, rad, 0)
}

// ｔ,numberの値不変
func setBm(slnum int, x float64, y float64, rad float64) {
	setBwitnT(slnum, BulletNum[slnum], x, y, rad, Bt[slnum])
}

// ｔ指定
func setBwitnT(slnum int, number int, x float64, y float64, rad float64, t int) {
	BulletNum[slnum] = number
	BulletRad[slnum] = rad
	Bx[slnum], By[slnum] = x, y
	Bt[slnum] = t
	Living[slnum] = true
}

func finaleB(endcou int, screenOff bool) (notEnd bool) {

	if !(counter < endcou) {
		MakeBullet(0)
		if screenOff {
			// Image.MainScreenPermission(false)
		}
	}
	return counter < endcou

}

func MakeBullet(howBullet int) {
	BulletNum = make([]int, howBullet)
	BulletRad = make([]float64, howBullet)
	Bx, By = make([]float64, howBullet), make([]float64, howBullet)
	Bt = make([]int, howBullet)
	Living = make([]bool, howBullet)
}

func closingMidle(min, max float64) (mid float64) {
	return min + (max-min)/2
}

// min <= ranum < max
func easyRand(min, max int, seed int64) (ranum int) {
	rand.Seed(time.Now().UnixNano() * int64(max-min) * (seed + 1))
	return min + rand.Intn(max-min)
}
