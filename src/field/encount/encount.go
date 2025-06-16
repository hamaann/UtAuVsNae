package encount

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/battle/outward"
	"github.com/hamaa/UtAuVsNae/src/field/event"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/avatar"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
	"github.com/hamaa/UtAuVsNae/src/load"
)

func EncountingWorld(screen *ebiten.Image, myX, myY float64, bigRoom *int, encountedRoom int) {
	// avatar.PutObject(avatar.Integrity(myY), screen, me, myX, myY, 0, 1, 1, 0)
	sx, sy, bt := event.RetEncountedData()
	EncountView(screen, myX, myY, sx, sy, bt, bigRoom, encountedRoom)
	// fmt.Println(event.Encounting)
}

var (
	toGoX, toGoY float64
	oldX, oldY   float64
)

func EncountView(screen *ebiten.Image, myX, myY, sgx, sgy float64, forStart int, bigRoom *int, encountedRoom int) {
	const (
		soulBlinkTime   = 30
		soulMovingTime  = 25
		soulWaitingTime = 15
	)

	if forStart <= soulBlinkTime {
		if !((forStart/5)%2 == 0) {
			if forStart%5 == 0 {
				sound.PlaySE(SEbattleSet, "wav")

			}
			avatar.PutObject(avatar.Integrity(myY-20), screen, outward.Soul, myX, myY-20, 0, 1, 1, 0)

		}
		// fmt.Println(forStart)

		if forStart == soulBlinkTime {
			sound.PlaySE(SEbattleFall, "wav")
			oldX, oldY, _ = avatar.ShowObjectPlace(gameX, gameY, avatar.CameraX, avatar.CameraY, avatar.CameraDistanse, myX, myY-20, 1)
			toGoX, toGoY = OnefGoing(oldX, sgx, oldY, sgy, soulMovingTime)
		}
	} else if forStart <= soulBlinkTime+soulMovingTime {
		show.ShowPathStd(screen, show.BoxPath(0, float32(gameX), 0, float32(gameY)), true, 1, 0x000000, 1, false)
		toGoTime := float64(forStart - soulBlinkTime)
		show.ShowImgStd(screen, outward.Soul, 1, oldX+toGoX*toGoTime, oldY+toGoY*toGoTime, 0, false, 1, 1, 1)
		// fmt.Println(oldX+toGoX*toGoTime, oldY+toGoY*toGoTime)
	} else if forStart <= soulBlinkTime+soulMovingTime+soulWaitingTime {
		show.ShowPathStd(screen, show.BoxPath(0, float32(gameX), 0, float32(gameY)), true, 1, 0x000000, 1, false)
		show.ShowImgStd(screen, outward.Soul, 1, oldX+toGoX*soulMovingTime, oldY+toGoY*soulMovingTime, 0, false, 1, 1, 1)
		if forStart == soulBlinkTime+soulMovingTime+soulWaitingTime {
			*bigRoom = encountedRoom
		}
	}

}

func OnefGoing(nx, wx, ny, wy float64, time int) (x, y float64) {
	return (wx - nx) / float64(time), (wy - ny) / float64(time)
}

var gameX, gameY float64
var (
	SEbattleSet  sound.SoundF
	SEbattleFall sound.SoundF
)

func init() {
	load.InitalPro[load.InEncount] = func(gx, gy float64) {

		// nowRoomNumber = 0
		gameX, gameY = gx, gy
		event.Encounting = false
		sound.LoadSound(&SEbattleSet, "material/sounds/se/snd_noise.wav")
		sound.LoadSound(&SEbattleFall, "material/sounds/se/snd_battlefall.wav")
		// LoadTurner(&nullCharacter, "material/img/mapObject/virtual")
		// show.LoadImage(,"material/img/mapObject/virtual/down.png")
	}
}
