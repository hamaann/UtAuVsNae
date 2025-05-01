package getmus

import "github.com/hajimehoshi/ebiten/v2"

var (
	MouseX, MouseY           int
	MouseL, MouseMid, MouseR bool
	WheelX, WheelY           float64
)

func MouseInput() {
	MouseX, MouseY = ebiten.CursorPosition()
	MouseL, MouseR = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft), ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	MouseMid = ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	WheelX, WheelY = ebiten.Wheel()
}


