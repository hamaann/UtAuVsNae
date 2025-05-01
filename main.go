package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hamaa/UtAuVsNae/src/battle/battle"
	"github.com/hamaa/UtAuVsNae/src/field/rooms"
	"github.com/hamaa/UtAuVsNae/src/fontlib"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/avatar"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show/showkage"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/nexting"
	"github.com/hamaa/UtAuVsNae/src/keylib"
	"github.com/hamaa/UtAuVsNae/src/load"
)

const DoubleSize = 1 //Image.DoubleSize
const DebugMode = !true

var QRtomove *ebiten.Image

const (
	GameX = 640
	GameY = 480
)

var room int

const (
	first = iota
	loading
	cushion
	field
	startfighting
	fighting
	allroom
)

type Game struct {
	playerName    string
	nowRoomNumber int
}

func (g *Game) Update() error {
	keylib.LibAllKeyCheck()

	switch room {
	case first:
		room = loading
		g.playerName = "？"
	case loading:
		if load.LoadedSystem(GameX, GameY) {//消してた、この後消していいかも
			room = cushion
		}
	case cushion:
		// if keyinp.KboolJstMany(keyinp.Kok1, keyinp.Kok2) {
		if keylib.JstInpKey(keylib.Decide, &keylib.DecideJst) {
			room = field //startfighting

		}
	case field:
		// location.UpdateWorld()
	case startfighting:
		battle.BTinit(g.playerName)
		room = fighting
	case fighting:
		battle.BttUpdate(GameX, GameY)
	}

	return nil
}

var onscreen *ebiten.Image

// var glichscreen *ebiten.Image
var ycic int

func (g *Game) Draw(screen *ebiten.Image) {
	switch room {
	case first:
	case loading:
		loadingView(screen)
	case cushion:
		// show.ShowPathStd(screen, show.LineBoxPath(100, 100, 110, 40, math.Pi*float64(ycic)/100), true, 1, 0xff0000, 1, false)
		show.PrintShowOrg(screen, "zを押してスタート\n\n<ミッション>\nてるてる坊主みたいな子に話しかけてください。\n「これはエラー」と出てきたらクリア\n操作方法はほぼ↓に準拠", fontlib.InBoxFont, 0, 0, 40)
		show.ShowImgStd(screen, QRtomove, 1, 200, GameY-100, 0, false, 1, 1, 1)
	case field:
		avatar.AvaterWorld(screen, rooms.RoomView[g.nowRoomNumber])
	case startfighting:
	case fighting:
		// Image.BoxDepictStd(onscreen, 0, GameX, 0, GameY, true, 1, 0x000000) //黒背景
		show.ShowPathStd(onscreen, show.BoxPath(0, GameX, 0, GameY), true, 1, 0x000000, 1, false)
		battle.BttDraw(onscreen, GameX, GameY)
	}

	ycic++

	nexting.NextingColor(screen, GameX, GameY, func(screen *ebiten.Image, x1, x2, y1, y2 float32, color int, alpha float64) {
		show.ShowPathStd(screen, show.BoxPath(x1, x2, y1, y2), true, 1, color, alpha, false)
	})
	/*	if Image.AccessableCheck() {
		scrrad, scrdx, scrdy, sizeX, sizeY, cbsize, scrglx, scrgly, scrSeed := Image.MainScreenAccess()
		ShowOnScreen(onscreen, screen, scrrad, scrdx, scrdy, sizeX, sizeY, GameX, GameY, cbsize, scrglx, scrgly, scrSeed)
	} else {*/
	ShowOnScreen(onscreen, screen, 0, 0, 0, 1, 1, GameX, GameY, 0, 0, 0, 0) //math.Pi*float64(ycic)/100
	/*}*/
	//math.Sin(math.Pi*float64(ycic)/100)
	//math.Abs(20*math.Sin(math.Pi*float64(ycic)/50))

	onscreen.Clear()

	if DebugMode {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("-debugMode-\n%f\n%f", ebiten.ActualFPS(), ebiten.CurrentFPS()))
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	var im *ebiten.Image
	show.LoadImage(&im, "material/img/rand.png")
	// showkage.LoadAllKage(im, &kageproTrFill, &kageproNoiFill, &kageproTest, &kageproWave, &kageproGlitch, &kageproGrad, &kageproAbr)

	// Image.MainScreenPermission(false)
	ebiten.SetWindowSize(640*DoubleSize, 480*DoubleSize)
	ebiten.SetWindowTitle("vs Nae")
	// showshader.SetShowShader(GameX, GameY, &shaderChrom)

	onscreen = ebiten.NewImage(GameX, GameY)
	// glichscreen = ebiten.NewImage(GameX, GameY)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if DebugMode {
		// debugeimaging.DebugModeON()
	}
	g := &Game{}
	g.nowRoomNumber = first
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}

func init() {
	room = first
	// sound.BattEm = SeDir
	// for i := 0; i != -1; {
	// 	fmt.Printf("loading %d %%\r", 100*i/10)
	// 	// i = battle.Initing(GameX, GameY, i)
	// }
	// fmt.Printf("\ndone\n")

	show.LoadImage(&QRtomove, "material/img/how_to_move.png")

}

func loadingView(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("System Loading...(%.1f%%)", load.ProgressLoadSystem()*100), (GameX/2)-70, (GameY / 2))
}

func ShowOnScreen(onscr *ebiten.Image, screen *ebiten.Image, rad float64, dx float64, dy float64, sizeX float64, sizeY float64, Mx float64, My float64, CBsize float64, glx float32, gly int, seed float32) {

	op := show.StdShowOpt(onscr, 1, (Mx/2)+dx, (My/2)+dy, rad, 1, false, sizeX, sizeY)
	//
	if gly > 3 {
		imgg := onscr
		// fmt.Println(seed)
		// seed = float32(time.Now().UnixNano() % 10000)
		showkage.ShowImgKage(screen, showkage.KageOptGlich(showkage.StdShowOptKage(imgg, 1, (Mx/2)+dx, (My/2)+dy, rad, 1, false, sizeX, sizeY),
			imgg, seed, glx, gly, float32(CBsize), 210), showkage.Glitch, imgg)
	} else if CBsize > 0 {
		//showshader.Chromaticaberration(screen, onscr, CBsize, Cturn, 1, (Mx/2)+dx, (My/2)+dy, rad, false, sizeX, sizeY, 1)
		imgg := onscr
		showkage.ShowImgKage(screen, showkage.KageOptAberrate(showkage.StdShowOptKage(imgg, 1, (Mx/2)+dx, (My/2)+dy, rad, 1, false, sizeX, sizeY),
			imgg, float32(CBsize), 210), showkage.Aberration, imgg)
	} else {
		show.ShowImg(screen, onscr, op)
	}

	// if len(Image.Glicher) > 0 {
	// 	// fmt.Println(">>  ", Image.Glicher)
	// 	for i := range Image.Glicher {
	// 		g := Image.Glicher[i]
	// 		Image.BoxDepictStd(glichscreen, float32(g.X1), float32(g.X2), float32(g.Y1), float32(g.Y2), true, 1, 0x000000)
	// 	}

	// 	for i := range Image.Glicher {
	// 		g := Image.Glicher[i]

	// 		im := onscr.SubImage(Image.GiveRecter(g.X1, g.X2, g.Y1, g.Y2)).(*ebiten.Image)

	// 		x1, x2, y1, y2 := g.X1, g.X2, g.Y1, g.Y2

	// 		show.ShowImgStd(glichscreen, im, 1, (x1+(x2-x1)/2.0)+g.Dx, (y1+(y2-y1)/2.0)+g.Dy, 0, false, 1, 1, 1)
	// 		// fmt.Println(i, ">>  ", Image.Glicher[i], " >> ", (x1+(x2-x1)/2)+g.Dx, (y1+(y2-y1)/2)+g.Dy)

	// 	}
	// 	if CBsize > 0 {
	// 		showshader.Chromaticaberration(screen, glichscreen, CBsize, Cturn, 1, (Mx/2)+dx, (My/2)+dy, rad, false, sizeX, sizeY, 1)
	// 	} else {
	// 		show.ShowImg(screen, glichscreen, op)
	// 	}
	// 	glichscreen.Clear()

	// 	//screenShot(screen)
	// 	// log.Fatal(Image.Glicher)
	// }

}

// 	r := math.Sqrt(addX*addX + addY*addY)
// 	xx := baseX + r*math.Cos(rad)
// 	yy := baseY + r*math.Sin(rad)
// 	return xx, yy
// }

var shout = false

func screenShot(screen *ebiten.Image) {
	if !shout {
		shout = true
		f, _ := os.Create("shout.png")
		defer f.Close()
		png.Encode(f, screen)
	}
}
