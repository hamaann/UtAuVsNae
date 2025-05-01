package menuUT

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/outward"
	"github.com/hamaa/UtAuVsNae/src/fontlib"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/load"
)

var (
	menuOn     = false
	menuLook   int
	menuLookOn = false
	itemLook   int
	itemLookOn = false
	itemUsing  int
)

// var saveItemNum int

const (
	lookItem = iota
	lookStat
	lookPhone
)

func MenuControl(cantMove *bool) {
	if !*cantMove || menuOn {
		if keylib.JstInpKey(keylib.Menu, &keylib.MenuJst) && !menuLookOn {
			*cantMove = !*cantMove
			menuOn = !menuOn
			menuLook = lookItem

			if menuOn {
				sound.PlaySE(SeSelMen, "wav")
			}
		}

		// if keyinp.KboolJstMany(keyinp.Kbk1, keyinp.Kbk2) {
		if keylib.JstInpKey(keylib.Cancel, &keylib.CancelJst) {
			if itemLookOn {
				itemLookOn = false
			} else if menuLookOn {
				menuLookOn = false
			} else {
				*cantMove = false
				menuOn = false
			}
		}

		if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) && menuOn {
			if !menuLookOn {
				if menuLook == lookItem {
					menuLookOn = true
					itemLook = 0
				}
			} else {
				switch menuLook {
				case lookItem:
					if !itemLookOn {
						itemLookOn = true
						itemUsing = 0
					}
				}

			}
			sound.PlaySE(SeDeciMen, "wav")
		}

		if menuOn {
			if !menuLookOn {
				dw := keylib.JustDir(hmutil.DOWN)
				up := keylib.JustDir(hmutil.UP)
				if hmutil.MenuSelecter(dw, up, &menuLook, 0, 3, 1) {
					//sound
					sound.PlaySE(SeSelMen, "wav")

				}
			} else {
				switch menuLook {
				case lookItem:
					if !itemLookOn {
						dw := keylib.JustDir(hmutil.DOWN)
						up := keylib.JustDir(hmutil.UP)
						if hmutil.MenuSelecter(dw, up, &itemLook, 0, items.HowItems(), 1) {
							//sound
							sound.PlaySE(SeSelMen, "wav")
						}
					} else {
						lf := keylib.JustDir(hmutil.LEFT)
						rg := keylib.JustDir(hmutil.RIGHT)
						if hmutil.MenuSelecter(rg, lf, &itemUsing, 0, 3, 1) {
							//sound
							sound.PlaySE(SeSelMen, "wav")
						}
					}
				}
			}
		}
	}
}

func MenuLook(screen *ebiten.Image, Name string, LOVE, HP, HPmax, Gold int) {
	if menuOn {
		dyEn := ktui.LtuningEng(7)
		ktui.UndertaleTextBox(screen, 31, 31+143, 52, 52+110)
		ktui.UndertaleTextBox(screen, 31, 31+143, 52+110+6, 316)
		// show.PrintShowStd(screen, Name, NameJPFont, 31+15, 70, 1, 0xffffff, 1)

		// show.PrintShowStd(screen, fmt.Sprintf("LV  %d\nHP  %d/%d\nG  %d", LOVE, HP, HPmax, Gold), StatasFont, 31+15, 100, 20, 0xffffff, 1)
		show.PrintShowOrg(screen, Name, fontlib.NameFontFD, 31+15, 70, 1)

		show.PrintShowOrg(screen, fmt.Sprintf("LV   %d\nHP   %d/%d\nG     %d", LOVE, HP, HPmax, Gold), fontlib.StatasFont, 31+15, 100, 20)
		//50,17*2
		show.PrintShowOrg(screen, "ITEM", fontlib.InBoxFont, 31+25*2, 52+110+6+23+dyEn, 1)
		show.PrintShowOrg(screen, "STAT", fontlib.InBoxFont, 31+25*2, 52+110+6+23+18*2+dyEn, 1)
		show.PrintShowOrg(screen, "PHONE", fontlib.InBoxFont, 31+25*2, 52+110+6+23+18*4+dyEn, 1)
		locked(screen, 52+110+6+23+18*2)
		locked(screen, 52+110+6+23+18*4)

		if !menuLookOn {
			var menuSoulY float64
			switch menuLook {
			case lookItem:
				menuSoulY = 52 + 110 + 6 + 23 + 13
			case lookStat:
				menuSoulY = 52 + 110 + 6 + 23 + 18*2 + 13
			case lookPhone:
				menuSoulY = 52 + 110 + 6 + 23 + 18*4 + 13
			}
			show.ShowImgStd(screen, outward.Soul, 1, 31+25+5, menuSoulY, 0, false, 1, 1, 1)
		} else {
			switch menuLook {
			case lookItem:
				ktui.UndertaleTextBox(screen, 31+143+17, 31+143+17+370, 52, 445)
				ViewItems(screen) //-27*2
				show.PrintShowOrg(screen, "USE", fontlib.InBoxFont, 31+143+17+50, 445-27*2+dyEn, 1)
				show.PrintShowOrg(screen, "INFO", fontlib.InBoxFont, 31+143+17+150, 445-27*2+dyEn, 1)
				show.PrintShowOrg(screen, "DROP", fontlib.InBoxFont, 31+143+17+268, 445-27*2+dyEn, 1)
				show.ShowPathStd(screen, show.BoxPath(31+143+17+5, 31+143+17+370-5, 445-27*2, 445-27), true, 1, 0x000000, 0.5, false)

				if !itemLookOn {
					show.ShowImgStd(screen, outward.Soul, 1, 31+143+17+25+5, 55+18*2*float64(itemLook+1)+13, 0, false, 1, 1, 1)
				} else {
					adddrop := 0.0
					if itemUsing == 2 {
						adddrop = 10
					}
					show.ShowImgStd(screen, outward.Soul, 1, 31+143+17+35+float64(105*itemUsing)+adddrop, 445-27*2+13, 0, false, 1, 1, 1)
				}
				// fmt.Println(">>")
			}
		}
		//item:y:455:x:370
		// show.ShowPathStd(screen, show.BoxPath(31+25*2, 31+143-5, 52+110+6+23+18*4, 52+110+6+23+18*4+18*2), true, 1, 0x000000, 0.5, false)

	}
}

func ViewItems(screen *ebiten.Image) {
	// fmt.Println(items.ItemList)
	dyEn := ktui.LtuningEng(7)
	for it := 0; it < items.HowItems(); it++ {
		nm := items.ItemName(items.ItemList[it])
		show.PrintShowOrg(screen, nm.Name, fontlib.InBoxFont, 31+143+17+60, 55+18*2*float64(it+1)+dyEn, 1)
		// fmt.Println(">>", nm.Name)
	}
}

func locked(screen *ebiten.Image, y float32) {
	show.ShowPathStd(screen, show.BoxPath(31+25*2, 31+143-5, y, y+18*2), true, 1, 0x000000, 0.5, false)
}

const nnn = 17 * 1.5

func GetMenuON() bool {
	return menuOn
}

var (
	SeSelMen  sound.SoundF
	SeDeciMen sound.SoundF
)

func init() {
	load.InitalPro[load.InMenuUT] = func(gx, gy float64) {
		sound.LoadSound(&SeSelMen, "material/sounds/se/snd_squeak.wav")
		sound.LoadSound(&SeDeciMen, "material/sounds/se/snd_select.wav")
	}
}

// var (
// 	NameJPFont text.Face
// 	StatasFont text.Face
// )

// func init() {
// 	load.InitalPro[load.InMenuUT] = func(gx, gy float64) {
// 		show.GetNewFont(&NameJPFont, "material/fonts/JF_Dot_Shinonome14.ttf", 14, 72)
// 		show.GetNewFont(&StatasFont, "material/fonts/4x4kanafont.ttf", 9.5, 72) //9.5
// 	}
// }
