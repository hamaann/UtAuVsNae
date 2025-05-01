package geno

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/load"
)

const halfSizeBar = 320 - 30

var barSay float32
var barStop bool

func SetBar() {
	barSay = -halfSizeBar
	barStop = false
	darkAndWhite = 0
}

func GoingBar(speed float32) (miss bool) {
	if !barStop {
		barSay += speed

		return barSay > halfSizeBar
	}

	return false
}

func StopBar(stop bool, damax, damin int, Hp *int) (stoped bool) {
	if stop {
		barStop = true
		damedMotion = dmonNum
		if notThrough() {
			ddd := (float64(damax-damin))*
				(1-math.Abs(float64(barSay)/halfSizeBar)) + float64(damin)
			HpView = *Hp
			*Hp -= int(ddd)
			herdamage = int(ddd)
			HpViewDec = ddd / decrTime

			sound.PlaySE(slash, "wav")
		} else {
			HpView = *Hp
			herdamage = 0
			HpView = 0
		}
		// damaging = int(ddd * 100)

		// decrDam = damaging / decrTime
		// decrDam /= 100
		// lostDamage = (damaging / 100) - (decrDam * decrTime)
		// damaging = decrDam * decrTime

	}
	return stop
}

func RetDamaging() int {
	return damedMotion
}

// func retBarSayPl() float64 {
// 	return float64(barSay)
// }

const (
	dmonNum  = 130
	decrTime = 80
)

var (
	damedMotion int
	HpView      int
	HpViewDec   float64
	herdamage   int
)

func ViewSlashDam(screen *ebiten.Image, gx float32, eneHP int, eneHPmax int, thiY float32) {
	if damedMotion > 0 {
		if decrTime+10 == damedMotion && notThrough() {
			sound.PlaySE(herHurt, "wav")
		}
		if decrTime+10 > damedMotion {

			if notThrough() {
				showEnemyHealth(screen, gx, HpView, eneHPmax, 50, thiY)
			}
			if damedMotion >= 10 {
				HpView -= int(HpViewDec)
				if damedMotion == 10 {
					HpView = eneHP
				}
			}
			halfstring := math.Log10(float64(herdamage)) / 2
			damY := float64(thiY) - 30
			goingdam := (float64(damedMotion-10) / float64(decrTime)) * math.Pi

			if notThrough() {
				// Image.PrintDepStd(screen, fmt.Sprintln(herdamage), fontlib.DamageFont, float64(gx/2)-(halfstring*1.3*20), damY-10*math.Sin(goingdam), 10)
				// show.PrintShowOrg(screen, fmt.Sprintln(herdamage), fontlib.DamageFont, float64(gx/2)-(halfstring*1.3*20), damY-10*math.Sin(goingdam), 10)
				PrintDamage(screen, herdamage, float64(gx/2)-(halfstring*1.3*20)-10, damY-10*math.Sin(goingdam))
				//Font.DamageFont
			} else {
				halfstring = 2
				// Image.PrintDepStd(screen, fmt.Sprintln("MISS"), fontlib.DamageFont, float64(gx/2)-(halfstring*1.3*20), damY-10*math.Sin(goingdam), 10)
				// show.PrintShowOrg(screen, fmt.Sprintln("MISS"), fontlib.DamageFont, float64(gx/2)-(halfstring*1.3*20), damY-10*math.Sin(goingdam), 10)
				PrintMiss(screen, float64(gx/2)-(halfstring*1.3*20)-10, damY-10*math.Sin(goingdam))
				//Font.DamageFont
			}
			//fmt.Println(goingdam)
		}

		damedMotion--
		// fmt.Println(HpView, eneHP)
	}
}

func showEnemyHealth(screen *ebiten.Image, gx float32, eneHP int, eneHPmax int, HpSmaller float32, thisY float32) {
	if eneHP < 0 {
		eneHP = 0
	}
	half := (float32(eneHPmax) / HpSmaller) / 2
	// Image.BoxDepictStd(screen, (gx/2)-half, (gx/2)+half, thisY, thisY+18, true, 1, 0x404040)
	show.ShowPathStd(screen, show.BoxPath((gx/2)-half, (gx/2)+half, thisY, thisY+18), true, 1, 0x404040, 1, false)
	// Image.BoxDepictStd(screen, (gx/2)-half, (gx/2)+(float32(eneHP)/float32(eneHPmax)*(2*half))-half, thisY, thisY+18, true, 1, 0x00ff00)
	show.ShowPathStd(screen, show.BoxPath((gx/2)-half, (gx/2)+(float32(eneHP)/float32(eneHPmax)*(2*half))-half, thisY, thisY+18), true, 1, 0x00ff00, 1, false)
}

// func PrintDamage(){
// 	if barStop{

// 	}
// }

var darkAndWhite int

func AttackBar(screen *ebiten.Image, gx float32) {
	dk := false
	if barStop {
		dk = ((darkAndWhite/10)%2 == 0)
	}
	itsBar(screen, (gx/2)+barSay, dk)
	darkAndWhite += 1
}

func itsBar(screen *ebiten.Image, x float32, dark bool) {
	out := 0x000000
	if dark {
		out = 0xffffff
	}
	ad := float32(2.5)

	// Image.BoxDepictStd(screen, x-6, x+6, 250+5, 250+140-5, true, 1, out)
	show.ShowPathStd(screen, show.BoxPath(x-6, x+6, 250+5, 250+140-5), true, 1, out, 1, false)
	// Image.BoxDepictStd(screen, x-6+ad, x+6-ad, 250+5+ad, 250+140-5-ad, true, 1, (out+0xffffff)%(0xffffff*2))
	show.ShowPathStd(screen, show.BoxPath(x-6+ad, x+6-ad, 250+5+ad, 250+140-5-ad), true, 1, (out+0xffffff)%(0xffffff*2), 1, false)
}

var (
	slash   sound.SoundF
	herHurt sound.SoundF
)

func init() {
	load.InitalPro[load.InGeno] = func(gx, gy float64) {
		sound.LoadSound(&slash, "material/sounds/se/snd_laz_c.wav")
		sound.LoadSound(&herHurt, "material/sounds/se/snd_damage_c.wav")
		var damageNumbers []*ebiten.Image
		show.LoadImageUnion(&damageNumbers, "material/img/damages.png", 30, 30, 13)
		missRune = damageNumbers[10:]
		damageRune[0] = damageNumbers[9]
		for i := 0; i < 9; i++ {
			damageRune[i+1] = damageNumbers[i] //1,2,3...8,9,0を0,1,2,3...8,9になおす
		}
		// damageRune = append(damageRune, damageNumbers[:9]...)

	}

}

func notThrough() bool {
	return barSay <= halfSizeBar
}

var missRune = make([]*ebiten.Image, 3)
var damageRune = make([]*ebiten.Image, 10)

func PrintDamage(screen *ebiten.Image, num int, x, y float64) {
	var showedNum = 0
	var prNum int
	howBig := float64(int(math.Log10(float64(num))) + 1)
	for i := 0.0; i < howBig; i++ {
		prNum = int(float64(num)/math.Pow10(int(howBig-i)-1)) - showedNum
		show.ShowImg(screen, damageRune[prNum], SideShowOpt(1, x+(i*31), y, 0, 1, 1, 1))
		showedNum += prNum
		showedNum *= 10
	}

}

func PrintMiss(screen *ebiten.Image, x, y float64) {
	for i := 0.0; i < 4; i++ {
		show.ShowImg(screen, missRune[int(math.Min(i, 2))], SideShowOpt(1, x+(i*31), y, 0, 1, 1, 1))
	}
}

// 非中心画像描画オプション
func SideShowOpt(size float64, x float64, y float64, rad float64, alpha float64, sizeX float64, sizeY float64) *ebiten.DrawImageOptions {
	Opt := &ebiten.DrawImageOptions{}

	//Opt.GeoM.Translate(sw, sh)

	// 画像を拡大/縮小する
	Opt.GeoM.Scale(size*sizeX, size*sizeY)

	// 画像を画面の左上を中心に回転させる（縦横半分めり込んでいるので、中心で回転することになる)
	Opt.GeoM.Rotate(rad /* / 180 * math.Pi*/)

	// 好きな位置へ移動させる
	Opt.GeoM.Translate(x, y)

	Opt.ColorScale.ScaleAlpha(float32(alpha))

	return Opt
}
