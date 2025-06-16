package event

import (
	"github.com/hamaa/UtAuVsNae/src/field/menuUT"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/nexting"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/stduttexts"
)

// func init() {
// 	debugeimaging.DebugModeON()
// }

func TouchBox(mx, my *float64, speed float64, x1, x2, y1, y2, add float64) {
	// debugeimaging.BebugMapCollingBox(float32(x1), float32(x2), float32(y1), float32(y2), float32(add), 0x0000ff)

	// rep := false
	// for {
	// 	for m := range *pusDir {
	// 		if (*pusDir)[2] || (*pusDir)[3] {
	// 			if (m == 0 || m == 1) && !rep {
	// 				continue
	// 			}
	// 		}
	// if (*pusDir)[m] {
	if t, m := InZoneDir(*mx, *my, x1, x2, y1, y2, add); t {
		switch m {
		case hmutil.DOWN:

			*my -= speed

		case hmutil.LEFT:
			// if InZone(*mx, *my, x1, x2, y1, y2, add) {
			*mx += speed
			// }
		case hmutil.RIGHT:
			// if InZone(*mx, *my, x1, x2, y1, y2, add) {
			*mx -= speed
			// }
		case hmutil.UP:
			// if InZone(*mx, *my, x1, x2, y1, y2, add) {
			*my += speed
			// }
		}
	}
	// 		}
	// 	}
	// 	rep = InZone(*mx, *my, x1, x2, y1, y2, add)
	// 	if !rep {
	// 		break
	// 	}
	// 	// if mdir[0] == mdir[1] {
	// 	// 	break
	// 	// }
	// }

}

func TouchCabin(mx, my *float64, speed float64, x1, x2, y1, y2, add float64) {
	TouchBox(mx, my, speed, x1-100, x1, y1-100, y2+100, add)
	TouchBox(mx, my, speed, x2, x2+100, y1-100, y2+100, add)
	TouchBox(mx, my, speed, x1-100, x2+100, y2, y2+100, add)
	TouchBox(mx, my, speed, x1-100, x2+100, y1-100, y1, add)
}

func ShiftRoomBox(mx, my *float64, nextRoom int, nowRoom *int, howLong int, outX, outY float64, x1, x2, y1, y2, add float64) {
	// debugeimaging.BebugMapCollingBox(float32(x1), float32(x2), float32(y1), float32(y2), float32(add), 0xff00ff)
	if InZone(*mx, *my, x1, x2, y1, y2, add) {
		// if godir==dir{

		if nexting.Nexting[int](nowRoom, nextRoom, howLong, 0x000000, nil) {
			go func(cm *bool) {

				for !nexting.NextingFin() {
					if *nowRoom == nextRoom {
						*mx, *my = outX, outY
					}
					// fmt.Println(roomshift.NowNexting())
				}
				*cm = false
			}(&cantMove)
		}
		cantMove = true
		// }
	}
}

var talkDecideJst int

func TalkingBox(mx, my *float64, mdir *int, speed float64, messnum int, x1, x2, y1, y2, add float64, roll []*textdata.Uttxt, talkWithChara bool, shiftingSign *bool, newSign bool) (talking bool) {
	nowTalk := false
	// debugeimaging.BebugMapCollingBox(float32(x1), float32(x2), float32(y1), float32(y2), float32(add), 0x00ffff)
	// if dir == 0 { //mdir {
	if t, d := InZoneDir(*mx, *my, x1, x2, y1, y2, add); t && !menuUT.GetMenuON() && d == *mdir {

		// if boxtloc.GoingMessage(messagedata.PushMessages(messnum), &cantMove) {
		if keylib.JstInpKey(keylib.Decide, &talkDecideJst) && !stduttexts.TextRollingNow {

			cantMove = true
			stduttexts.StartTextdata(roll, stduttexts.Boxt, 0, 0, talkWithChara, false)
			go func(cm *bool, sgn *bool, newsgn bool) {
				for stduttexts.TextRollingNow {

				}
				if sgn != nil {
					*sgn = newsgn

				}
				*cm = false
			}(&cantMove, shiftingSign, newSign)

		}

		// }
		nowTalk = cantMove
	}
	// }

	return nowTalk
}

func Test1Box(mx, my *float64, speed float64, x1, x2, y1, y2, add float64, soulx, souly float64, encou *bool) {
	// debugeimaging.BebugMapCollingBox(float32(x1), float32(x2), float32(y1), float32(y2), float32(add), 0xffff00)
	// if dir == 0 {
	if InZone(*mx, *my, x1, x2, y1, y2, add) {
		// Encount(hmutil.KboolJstMany(hmutil.Kok1, hmutil.Kok2), soulx, souly)
		// if hmutil.KboolJstMany(hmutil.Kok1, hmutil.Kok2) {
		// 	Encounting = true
		// 	goingX, goingY = soulx, souly
		// 	battleGoing = 0
		// }
		// if Encounting {
		// 	battleGoing++
		// }
		if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
			*encou = true
		}
	}
	// }
}

// 箱 [x1,x2,y1,y2]
func InZone(mx, my, x1, x2, y1, y2 float64, Add float64) bool {
	return mx < x2+Add && mx > x1-Add && my > y1-Add && my < y2+Add

}

func InZoneDir(mx, my, x1, x2, y1, y2 float64, Add float64) (inz bool, dir int) {
	t := InZone(mx, my, x1, x2, y1, y2, Add)
	nomalDiagonal := (my >= ((y2-y1)/(x2-x1))*(mx-x1)+y1)
	abnomalDiagonal := (my >= ((y1-y2)/(x2-x1))*(mx-x1)+y2)
	var dirr int
	switch {
	case nomalDiagonal && abnomalDiagonal:
		dirr = hmutil.UP
	case !nomalDiagonal && !abnomalDiagonal:
		dirr = hmutil.DOWN
	case nomalDiagonal && !abnomalDiagonal:
		dirr = hmutil.RIGHT
	case !nomalDiagonal && abnomalDiagonal:
		dirr = hmutil.LEFT
	}
	return t, dirr
}

func EachDir(dir int) int {
	if dir%2 == 0 {
		dir++
	} else {
		dir--
	}
	return dir
}

func CheckNotMoveable() bool {
	return cantMove
}

func LetsMenu() {
	menuUT.MenuControl(&cantMove)
}

var cantMove = false

const (
	EncStdSx = 85 - 38
	EncStdSy = 453
)

func Encount(signal *bool, soulx, souly float64) {
	if *signal {
		*signal = false

		Encounting = true
		goingX, goingY = soulx, souly
		battleGoing = 0
		cantMove = true
	}
	if Encounting {
		battleGoing++
	}
}

var (
	Encounting     = false
	goingX, goingY float64
	battleGoing    int
)

func RetEncountedData() (sgx, sgy float64, forStart int) {
	return goingX, goingY, battleGoing
}
