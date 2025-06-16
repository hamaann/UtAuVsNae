package rooms

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/field/event"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/avatar"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/load"
)

const AllRoom = 3

var RoomView [AllRoom]func(order avatar.Integrity, screen *ebiten.Image)
var RoomEvent [AllRoom]func(mx, my *float64, mdir *int, speed float64, room *int)

var (
	NullCharacter Turner
	floatMan      Turner
	things        [2]*ebiten.Image
	testYuzu      *ebiten.Image
	letsGo        *ebiten.Image
)

var mainCharaX, mainCharaY float64
var mainCharaDirection int

func letUpdateMe(myX, myY float64, myDirec int) {
	mainCharaX, mainCharaY = myX, myY
	mainCharaDirection = myDirec
}

func letLookMe(order avatar.Integrity, screen *ebiten.Image) {
	me := NullCharacter[mainCharaDirection]
	avatar.PutObject(order, screen, me, mainCharaX, mainCharaY, 1, 1, 1, 0)
}

var fieldMessData map[string][]*textdata.Uttxt

func txtData(key string) []*textdata.Uttxt {
	// t := (*fieldMessData)[key]
	// fmt.Println((*fieldMessData)[key])
	// fmt.Println((fieldMessData), (fieldMessData)[key])
	return (fieldMessData)[key]
}

func init() {
	textdata.GetTextData(&fieldMessData, "material/data/fieldMess.csv", 18)
	// fmt.Println((fieldMessData))

	load.InitalPro[load.InRooms] = func(gx, gy float64) {
		LoadTurner(&NullCharacter, "material/img/mapObject/virtual")
		LoadTurner(&floatMan, "material/img/mapObject/float_man")
		show.LoadImage(&testFloor, "material/img/floor/interim_floor.png")
		show.LoadImage(&testYuzu, "material/img/mapObject/Yuzu.png")
		show.LoadImage(&things[0], "material/img/mapObject/who.png")
		show.LoadImage(&things[1], "material/img/mapObject/who_is_this.png")
		show.LoadImage(&letsGo, "material/img/letsGo.png")

		var roomNum int

		/*when*/
		roomNum = 0
		RoomView[roomNum] = func(order avatar.Integrity, screen *ebiten.Image) {
			letLookMe(order, screen)
			avatar.PutFloor(order, screen, testFloor, 0, 0, 1, 1)
			avatar.PutFloor(order, screen, letsGo, 400, -180, 1, 1)
			avatar.PutFloor(order, screen, letsGo, 0, -180, 1, 1)
			avatar.PutObject(order, screen, things[0], -100, 0, 0, 1, 1, 0)
			avatar.PutObject(order, screen, things[0], 100, 0, 0, 1, 1, 0)
			avatar.PutObject(order, screen, things[1], 130, 50, 0, 1, 1, 0)

		}

		RoomEvent[roomNum] = func(mx, my *float64, mdir *int, speed float64, room *int) {
			letUpdateMe(*mx, *my, *mdir)
			UTcamera(*mx, gx, 10, -500, 500, *my, gy, 10, -150, 150, 25)

			event.TouchCabin(mx, my, speed, -500, 500, -150, 150, 0)
			event.ShiftRoomBox(mx, my, 1, room, 20, 20-500, 0, 500-10, 500, -150, 150, 3)
			event.TouchBox(mx, my, speed, 130-35, 130+35, 50-40, 50+10, 0)

			event.TalkingBox(mx, my, mdir, speed, 0, 130-35, 130+35, 50-40, 50+10, 30, txtData("box_test"), false, nil, true)

		}

		/*when*/
		roomNum = 1
		RoomView[roomNum] = func(order avatar.Integrity, screen *ebiten.Image) {
			letLookMe(order, screen)
			avatar.PutFloor(order, screen, testFloor, 0, 0, 1, 1)

			avatar.PutObject(order, screen, floatMan[teruteruDir], -100, 0, 0, 1, 1, 0)

			avatar.PutObject(order, screen, testYuzu, 0, 300, 0, 1, 1, 0)
		}

		var goingEnc = false
		RoomEvent[roomNum] = func(mx, my *float64, mdir *int, speed float64, room *int) {
			letUpdateMe(*mx, *my, *mdir)
			// avatar.CameraSet(math.Min(*mx, 20+(gx/2)), 0)
			UTcamera(*mx, gx, 10, -500, 500, *my, gy, 10, -150, 150, 25)
			event.TouchCabin(mx, my, speed, -500, +500, -150, 150, 0)
			event.ShiftRoomBox(mx, my, 0, room, 20, 500-20, 0, -500, -500+10, -150, 150, 3)

			teruteruTelling := event.TalkingBox(mx, my, mdir, speed, 1, -100-35, -100+35, 0-40, 0+10, 30, txtData("nae_test"), true, &goingEnc, true)
			if teruteruTelling {
				teruteruDir = event.EachDir(*mdir)
			} else {
				teruteruDir = hmutil.DOWN
			}
			// event.Test1Box(mx, my, speed, dir, 100, 200, -100, 100, 0, 85-38, 453, &goingEnc)
			// if dir == 0 {

			event.Encount(&goingEnc, event.EncStdSx, event.EncStdSy)
			// }
			// if goingEnc {
			// 	fmt.Println(goingEnc)
			// }
			// event.TouchBox(mx, my, speed, dir, 130-35, 130+35, 50-40, 50+10, 0)
		}
	}

}

var ( //Objects
	teruteruDir int
)

// func funcEncount() {
// 	event.Encount(true, event.EncStdSx, event.EncStdSy)
// }

func UTcamera(mx, gx, addgx, rx1, rx2 float64, my, gy, addgy, ry1, ry2 float64, mHigh float64) {
	if rx2-(gx/2)+addgx > rx1+(gx/2)-addgx {
		mx = math.Min(mx, rx2-(gx/2)+addgx)
		if mx != rx2-(gx/2)+addgx {
			mx = math.Max(rx1+(gx/2)-addgx, mx)
		}
	} else {
		mx = rx1 + (rx2-rx1)/2
	}

	if ry2-(gy/2)+addgy > ry1+(gy/2)-addgy {
		my = math.Min(my, ry2-(gy/2)+addgy)
		if my != ry2-(gy/2)+addgy {
			my = math.Max(ry1+(gy/2)-addgy, my)
		}
	} else {
		my = ry1 + (ry2-ry1)/2
	}
	avatar.CameraSet(mx, my, mHigh)
}

type Turner [4]*ebiten.Image

var testFloor *ebiten.Image

func LoadTurner(PicsPointer *Turner, path string) {
	for r := 0; r < 4; r++ {
		st := ""
		switch r {
		case hmutil.DOWN:
			st = "down"
		case hmutil.UP:
			st = "up"
		case hmutil.LEFT:
			st = "left"
		case hmutil.RIGHT:
			st = "right"
		}

		show.LoadImage(&(*PicsPointer)[r], fmt.Sprintf("%s/%s.png", path, st))
	}
}
