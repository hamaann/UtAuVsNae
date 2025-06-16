package fieldmove

import (
	"github.com/hamaa/UtAuVsNae/src/field/event"
	"github.com/hamaa/UtAuVsNae/src/field/rooms"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
	"github.com/hamaa/UtAuVsNae/src/keylib"
)

const DefaultUtSpeed = 5.0

var (
	myX     float64
	myY     float64
	myDirec int
	mySpeed = DefaultUtSpeed
)

func MyFieldPlace() (x, y float64) {
	return myX, myY
}

func MyDirection() (dir int) {
	return myDirec
}

func Location(nowRoomNumber *int) {
	for d := 0; d <= 3; d++ {
		if n := GoingToDir(&myX, &myY, mySpeed, d); n != -1 {

			myDirec = n

		}

	}

	rooms.RoomEvent[*nowRoomNumber](&myX, &myY, &myDirec, mySpeed, nowRoomNumber)
	event.LetsMenu()
}

func GoingToDir(sx *float64, sy *float64, speed float64, godir int) (dir int) {

	if keylib.Cross[godir].Pushed {

		if event.CheckNotMoveable() {

			return -1
		}
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
		return godir
	}

	return -1
}
