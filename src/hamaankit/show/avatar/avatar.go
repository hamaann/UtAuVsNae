package avatar

import (
	//hamaankit/show/
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
)

var (
	GameWidth        float64
	GameHeight       float64
	CameraX, CameraY float64
	CameraDistanse   float64
)

type Integrity int

var cameheight float64

var DrawTop func(screen *ebiten.Image, gx, gy float64)

func AvaterWorld(screen *ebiten.Image, drawWorld func(order Integrity, screen *ebiten.Image)) {
	// DebugShowAbouLocationSet()

	for j_y := ScreenMin(); j_y <= ScreenMax()+int(HighestObject); j_y++ {
		drawWorld(Integrity(j_y), screen)
	}

	if DrawTop != nil {
		DrawTop(screen, GameWidth, GameHeight)
	}

}

func AvaterScreenSize(screenWidth, screenHeight int) {
	avaterInit(float64(screenWidth), float64(screenHeight))
}

func avaterInit(gx, gy float64) {

	GameHeight = gy
	GameWidth = gx
	CameraX, CameraY, CameraDistanse = 0, 0, 1
	cameheight = gy * (25.0 / 480.0)
	// objectOrder = make(map[int]func(screen *ebiten.Image))

}

// var objectOrder map[int]func(screen *ebiten.Image) //= make(chan int)

func PutObject(order Integrity, screen, pic *ebiten.Image, x, y, z, size, alpha, underset float64) {

	inHighestObject(float64(pic.Bounds().Dy()))
	if ScreenMin() < int(y) && ScreenMax()+int(HighestObject) > int(y) {
		if order == Integrity(y) {

			show.ShowImg(screen, pic, StdObjOpt(pic, size, x, y, z, underset, alpha, 1, 1, true))
		}

	}
}

func PutFloor(order Integrity, screen, pic *ebiten.Image, x, y, size, alpha float64) {
	if order == Integrity(ScreenMin()) {
		show.ShowImg(screen, pic, StdObjOpt(pic, size, x, y, 0, 0, alpha, 1, 1, false))
	}
}

func PutOuter(order Integrity, screen, pic *ebiten.Image, x, y, size, alpha float64) {
	if order == Integrity(ScreenMax()+int(HighestObject)) {
		show.ShowImg(screen, pic, StdObjOpt(pic, size, x, y, 0, 0, alpha, 1, 1, false))
	}
}

func PutMultiObject(order Integrity, screen *ebiten.Image, drawOb func(screen *ebiten.Image, x0, y0, sizeb float64), high, wideMid, x, y, z, underset, size float64) {
	inHighestObject(high)
	if ScreenMin() < int(y) && ScreenMax()+int(HighestObject) > int(y) {
		if order == Integrity(y) {
			nx, ny, nsize := ShowObjectPlace(GameWidth, GameHeight, CameraX, CameraY, CameraDistanse, x, y, size)
			drawOb(screen, nx-wideMid*nsize, ny-high*nsize, nsize)
		}
	}
}

func PutMultiPicture2D(order Integrity, screen *ebiten.Image, drawPic func(screen *ebiten.Image, x0, y0, sizeb float64), x, y, size float64, floor bool) {
	t := (order == Integrity(ScreenMax()+int(HighestObject)))
	if floor {
		t = (order == Integrity(ScreenMin()))
	}

	if t {
		nx, ny, nsize := ShowObjectPlace(GameWidth, GameHeight, CameraX, CameraY, CameraDistanse, x, y, size)
		drawPic(screen, nx, ny, nsize)
	}
}

func ShowObjectPlace(Gwidth, Gheight, cameX, cameY, cameDst float64, x, y, size float64) (tx, ty, tsize float64) {
	// image.ShowUnder(screen, img, size, x_base+((x-z_x)*(1/SH)), y_base+((y-z_y)*(1/SH)), 0, ch_x/SH, ch_y/SH)
	baseX, baseY := (Gwidth / 2), (Gheight/2)+(cameheight/cameDst)
	size = size / cameDst
	//show.ShowImgStd(screen,pic,size,baseX+((x-cameX)*(1/cameDst)),baseY+((y-cameY)*(1/cameDst)),0,true,1,1,1)
	return baseX + ((x - cameX) * (1 / cameDst)), baseY + ((y - cameY) * (1 / cameDst)), size
}

func StdObjOpt(Pic *ebiten.Image, size, x, y, z, setUnderHigh float64, alpha float64, sizeX, sizeY float64, standingObject bool) *ebiten.DrawImageOptions {
	nx, ny, nsize := ShowObjectPlace(GameWidth, GameHeight, CameraX, CameraY, CameraDistanse, x, y, size)
	return show.StdShowOptNew(Pic, nsize, nx, ny, 0, alpha, standingObject, setUnderHigh-z, 1, 1)
}

func CameraSet(x, y, height float64) {
	CameraX, CameraY = x, y
	cameheight = height
}

func AddCameraDist(distAdd float64) {
	CameraDistanse += distAdd

}

func ScreenMin() int {
	return int(CameraY - (float64((GameHeight*CameraDistanse)/2.0) + cameheight))
}

func ScreenMax() int {
	return int(CameraY + (float64((GameHeight*CameraDistanse)/2.0) + cameheight))
}

var HighestObject = 0.0

func inHighestObject(high float64) {
	// highestObject = math.Max(high, highestObject)
	if HighestObject < high {
		HighestObject = high
		// fmt.Println(HighestObject)
	}
}

// var drawed bool

// func DebugShowAbouLocation(screen *ebiten.Image) {
// 	if !drawed {
// 		// debugeimaging.DebugViewMCB(screen, GameWidth, GameHeight, CameraX, CameraY, CameraDistanse, float32(cameheight))
// 		drawed = true
// 	}
// }

// func DebugShowAbouLocationSet() {
// 	drawed = false
// }

func AvaterMouseXY(mx, my float64) (x, y float64) {
	// mx, my := getmus.MouseX, getmus.MouseY
	return CameraX + float64(mx) - (GameWidth / 2), CameraY + float64(my) + cameheight - (GameHeight / 2)
}
