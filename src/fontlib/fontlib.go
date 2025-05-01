package fontlib

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
)

var (
	NameFontBT      text.Face
	NameFontFD      text.Face
	NumsLove        text.Face
	KRfont          text.Face
	InBoxFont       text.Face
	InTermFont      text.Face
	DamageFont      text.Face
	StatasFont      text.Face
	EnglishAsteFont text.Face
)

func SetFonts(lang int, jap int, eng int) {

	switch lang {
	case jap:
		show.GetNewFont(&InBoxFont, "material/fonts/JF_Dot_Shinonome14.ttf", 1.3*20, 72)
		show.GetNewFont(&InTermFont, "material/fonts/BIZUDGothic_Regular.ttf", 13, 72)
		show.GetNewFont(&NameFontFD, "material/fonts/JF_Dot_Shinonome14.ttf", 14, 72)
		show.GetNewFont(&NameFontBT, "material/fonts/JF_Dot_Shinonome14.ttf", 1.05*15, 72)

	case eng:
		show.GetNewFont(&InBoxFont, "material/fonts/DePixelHalbfett.ttf", 1.3*12, 72)
		show.GetNewFont(&InTermFont, "material/fonts/DePixelHalbfett.ttf", 13*0.8, 72)
		show.GetNewFont(&NameFontFD, "material/fonts/4x4kanafont.ttf", 14, 72)
		show.GetNewFont(&NameFontBT, "material/fonts/4x4kanafont.ttf", 1.05*15, 72)
	}

	show.GetNewFont(&NumsLove, "material/fonts/4x4kanafont.ttf", 1.25*13, 72)
	show.GetNewFont(&DamageFont, "material/fonts/4x4kanafont.ttf", 1.3*20, 72)
	show.GetNewFont(&StatasFont, "material/fonts/4x4kanafont.ttf", 9.5, 72) //9.5
	show.GetNewFont(&KRfont, "material/fonts/4x4kanafont.ttf", 1.2*10, 72)

	show.GetNewFont(&EnglishAsteFont, "material/fonts/JF_Dot_Shinonome14.ttf", 1.3*20, 72)

}

// func LcheckEng() bool {
// 	return Lang == Eng
// }

// func LtuningEng(num float64) float64 {
// 	if LcheckEng() {
// 		return num
// 	}
// 	return 0
// }

// func Aste() string {
// 	if LcheckEng() {
// 		return "*"
// 	}
// 	return "＊"
// }
