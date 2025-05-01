package nexting

import "github.com/hajimehoshi/ebiten/v2"

// func main() {
// 	var room = 0

// 	Nexting[int](&room, 5, 50, 0x000000)
// 	for nexton {
// 		fmt.Println(float32(nextroop)/float32(TimeForDeep), room)
// 		NextingPoint()

// 	}
// }

func NextingFin() bool {
	return !nexton
}

var nextroop = 0
var NexterColor = int(0xffffff)
var TimeForDeep = int(1)
var gostoper = make(chan bool)
var nexton = false

func Nexting[someType any](Chenge_var *someType, how_chenge someType, time int, color int, shiftInit *func()) (first bool) {
	if !nexton {
		NexterColor = color
		TimeForDeep = time
		nexton = true

		go func() {
			for {
				if nextroop != TimeForDeep {
					// image.DrawGageAlpha(screen, 0, x_max, 0, y_max, 0x00, 0x00, 0x00, float32(nextroop)/10.0)
					nextroop++
					<-gostoper

				} else {
					if shiftInit != nil {
						(*shiftInit)()
					}
					*Chenge_var = how_chenge

					go func() {
						for {
							if nextroop != 0 {
								nextroop--
								<-gostoper
							} else {

								nexton = false
								break
							}

						}

					}()

					break
				}

			}
		}()

		return true
	}

	return false

}

func nextingPoint() {
	if nexton {
		gostoper <- true
	}
}

func NextingColor(screen *ebiten.Image, x_max, y_max float32, drawbox func(screen *ebiten.Image, x1, x2, y1, y2 float32, color int, alpha float64)) {
	// var r, g, b float32

	// colorDismantling(&r, &g, &b, NexterColor)
	// image.DrawGageAlpha(screen, 0, x_max, 0, y_max, r, g, b, float32(nextroop)/float32(TimeForDeep))
	drawbox(screen, 0, x_max, 0, y_max, NexterColor, float64(nextroop)/float64(TimeForDeep))
	nextingPoint()

}
