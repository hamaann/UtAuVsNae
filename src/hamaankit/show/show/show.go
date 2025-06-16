package show

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

// 画像読み込み
func LoadImage(PicPointer **ebiten.Image, Path string) {
	var imgLoaderr error
	*PicPointer, _, imgLoaderr = ebitenutil.NewImageFromFile(Path)
	if imgLoaderr != nil {
		log.Fatal("about:", Path, "\n", imgLoaderr)
	}
}

func LoadImageUnion(PicPointer *[]*ebiten.Image, Path string, Width, Height int, ImageNum int) {
	var unionIm *ebiten.Image
	LoadImage(&unionIm, Path)
	// var uniImFellApart []*ebiten.Image
	*PicPointer = make([]*ebiten.Image, ImageNum)
	stW, stH := 0, 0
	for i := 0; i < ImageNum; i++ {
		(*PicPointer)[i] = unionIm.SubImage(image.Rect(stW, stH, stW+Width, stH+Height)).(*ebiten.Image)
		stW += Width
		if stW+Width > unionIm.Bounds().Dx() {
			stW = 0
			stH += Height
			if stH+Height > unionIm.Bounds().Dy() {
				fmt.Println("Images you want is too much")
				break
			}
		}
	}
	// PicPointer = &uniImFellApart
	// fmt.Println(len(*PicPointer))

}

// 通常画像描画オプション
func StdShowOpt(Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, sizeX float64, sizeY float64) *ebiten.DrawImageOptions {
	Opt := &ebiten.DrawImageOptions{}
	// Opt.GeoM.Scale(sizeX, sizeY)
	w, h := Pic.Size()
	// 係数で画像を拡大/縮小したときの大きさを計算しておく
	var sw, sh float64 = float64(w) * sizeX * size, float64(h) * sizeY * size
	//Opt.GeoM.Translate(sw, sh)

	// 画像を拡大/縮小する
	Opt.GeoM.Scale(size*sizeX, size*sizeY)

	// 縮小したサイズに合わせて、画面の左上に縦横半分めり込む形にする
	Opt.GeoM.Translate(-sw/2, -sh/2)
	if setUnder {
		Opt.GeoM.Translate(0, -sh/2)
	}

	// 画像を画面の左上を中心に回転させる（縦横半分めり込んでいるので、中心で回転することになる)
	Opt.GeoM.Rotate(rad /* / 180 * math.Pi*/)

	// 好きな位置へ移動させる
	Opt.GeoM.Translate(x, y)

	Opt.ColorScale.ScaleAlpha(float32(alpha))

	return Opt
}

// 通常画像描画オプション
func StdShowOptNew(Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, setUnderHigh float64, sizeX float64, sizeY float64) *ebiten.DrawImageOptions {
	Opt := &ebiten.DrawImageOptions{}
	// Opt.GeoM.Scale(sizeX, sizeY)
	w, h := Pic.Size()
	// 係数で画像を拡大/縮小したときの大きさを計算しておく
	var sw, sh float64 = float64(w) * sizeX * size, float64(h) * sizeY * size
	//Opt.GeoM.Translate(sw, sh)

	// 画像を拡大/縮小する
	Opt.GeoM.Scale(size*sizeX, size*sizeY)

	// 縮小したサイズに合わせて、画面の左上に縦横半分めり込む形にする
	Opt.GeoM.Translate(-sw/2, -sh/2)
	if setUnder {
		Opt.GeoM.Translate(0, (-sh/2)+setUnderHigh)
	}

	// 画像を画面の左上を中心に回転させる（縦横半分めり込んでいるので、中心で回転することになる)
	Opt.GeoM.Rotate(rad /* / 180 * math.Pi*/)

	// 好きな位置へ移動させる
	Opt.GeoM.Translate(x, y)

	Opt.ColorScale.ScaleAlpha(float32(alpha))

	return Opt
}

func ShowImg(screen *ebiten.Image, Pic *ebiten.Image, Opt *ebiten.DrawImageOptions) {
	screen.DrawImage(Pic, Opt)
}

func ShowImgStd(screen *ebiten.Image, Pic *ebiten.Image, size float64, x float64, y float64, rad float64, setUnder bool, sizeX float64, sizeY float64, alpha float64) {
	ShowImg(screen, Pic, StdShowOpt(Pic, size, x, y, rad, alpha, setUnder, sizeX, sizeY))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func StdShowOptFill(Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, sizeX float64, sizeY float64) *colorm.DrawImageOptions {
	Opt := &colorm.DrawImageOptions{}
	// Opt.GeoM.Scale(sizeX, sizeY)
	w, h := Pic.Size()
	// 係数で画像を拡大/縮小したときの大きさを計算しておく
	var sw, sh float64 = float64(w) * sizeX * size, float64(h) * sizeY * size
	//Opt.GeoM.Translate(sw, sh)

	// 画像を拡大/縮小する
	Opt.GeoM.Scale(size*sizeX, size*sizeY)

	// 縮小したサイズに合わせて、画面の左上に縦横半分めり込む形にする
	Opt.GeoM.Translate(-sw/2, -sh/2)
	if setUnder {
		Opt.GeoM.Translate(0, -sh/2)
	}

	// 画像を画面の左上を中心に回転させる（縦横半分めり込んでいるので、中心で回転することになる)
	Opt.GeoM.Rotate(rad /* / 180 * math.Pi*/)

	// 好きな位置へ移動させる
	Opt.GeoM.Translate(x, y)

	return Opt
}

func ShowImgFill(screen *ebiten.Image, Pic *ebiten.Image, Opt *colorm.DrawImageOptions, fillColor int, alpha float64) {
	var r, g, b float64
	ColorDismantlingIM(&r, &g, &b, fillColor)
	// if fillColor == 0x4e0000 {
	// 	fmt.Println(r, g, b, alpha)
	// }

	var cm colorm.ColorM
	cm.Scale(0, 0, 0, alpha)
	cm.Translate(r/0xff, g/0xff, b/0xff, 0)
	colorm.DrawImage(screen, Pic, cm, Opt)
}

func ShowImgFillStd(screen *ebiten.Image, Pic *ebiten.Image, size float64, x float64, y float64, rad float64, setUnder bool, sizeX float64, sizeY float64, fillColor int, alpha float64) {
	ShowImgFill(screen, Pic, StdShowOptFill(Pic, size, x, y, rad, alpha, setUnder, sizeX, sizeY), fillColor, alpha)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BoxPath(x1 float32, x2 float32, y1 float32, y2 float32) *vector.Path {
	var path vector.Path
	path.MoveTo(x1, y1)
	path.LineTo(x1, y2)
	path.LineTo(x2, y2)
	path.LineTo(x2, y1)

	path.Close()

	return &path
}

func LineBoxPath(x, y, wide, high, rad float64) *vector.Path {
	var path vector.Path
	path.MoveTo(baseRad(x, y, wide, high, rad))   //x+wide, y+high)
	path.LineTo(baseRad(x, y, wide, -high, rad))  //x+wide, y-high)
	path.LineTo(baseRad(x, y, -wide, -high, rad)) //x-wide, y-high)
	path.LineTo(baseRad(x, y, -wide, high, rad))  //x-wide, y+high)
	path.LineTo(baseRad(x, y, wide, high, rad))

	path.Close()

	return &path
}

func CirclePath(x float32, y float32, radius float32, rad1 float32, rad2 float32) *vector.Path {
	var path vector.Path
	path.MoveTo(x, y)
	path.Arc(x, y, radius, rad1, rad2, vector.Clockwise)
	path.Close()

	return &path
}

func ArcBoxPath(x1 float32, x2 float32, y1 float32, y2 float32, arcsize float32) *vector.Path {
	var path vector.Path

	path.MoveTo(x1+arcsize, y1)
	path.LineTo(x2-arcsize, y1)
	path.ArcTo(x2, y1, x2, y1+arcsize, arcsize)
	path.LineTo(x2, y1+arcsize)
	path.LineTo(x2, y2-arcsize)
	path.ArcTo(x2, y2, x2-arcsize, y2, arcsize)
	path.LineTo(x2-arcsize, y2)
	path.LineTo(x1+arcsize, y2)
	path.ArcTo(x1, y2, x1, y2-arcsize, arcsize)
	path.LineTo(x1, y2-arcsize)
	path.LineTo(x1, y1+arcsize)
	path.ArcTo(x1, y1, x1+arcsize, y1, arcsize)

	path.Close()

	return &path
}

// 2024_9_27///////////////////////////
func FunctionPath(f func(t float32) (x, y float32), tmin, tmax float32, resolution float32) *vector.Path {
	var path vector.Path
	path.MoveTo(f(tmin))
	for tt := tmin + resolution; tt < tmax; tt += resolution {
		path.LineTo(f(tt))
		// fmt.Println(f(tt))
	}

	return &path
}

/////////////////////////////////////////

// 2025_2_19///////////////////////////
func IsoscelesPath(x, y, w, h float32, turn bool) *vector.Path {
	if !turn {
		return TringlePath(x, x+(w/2), x-(w/2), y+h, y, y)
	} else {
		return TringlePath(x+w, x, x, y, y+(h/2), y-(h/2))
	}
}

//////////////////////////////////////////

func TringlePath(x1 float32, x2 float32, x3 float32, y1 float32, y2 float32, y3 float32) *vector.Path {
	var path vector.Path
	path.MoveTo(x1, y1)
	path.LineTo(x2, y2)
	path.LineTo(x3, y3)

	path.Close()
	return &path
}

func ShowPathStd(screen *ebiten.Image, path *vector.Path, Fill bool, nonFillLineSize float32, Color int, alpha float64, antiAlias bool) {
	ShowPath(screen, StdShowOptTri(Fill, antiAlias), path, Fill, nonFillLineSize, Color, alpha)
}

func ShowPath(screen *ebiten.Image, Opt *ebiten.DrawTrianglesOptions, path *vector.Path, Fill bool, nonFillLineSize float32, Color int, alpha float64) {
	retPathDrawer(path, Fill, nonFillLineSize, Color, alpha).ShowTriangles(screen, Opt)
}

func (pd *pathDrawer) ShowTriangles(screen *ebiten.Image, Opt *ebiten.DrawTrianglesOptions) {
	screen.DrawTriangles(*(pd.vs), *(pd.is), whiteSubImage, Opt)
}

func StdShowOptTri(Fill bool, antiAlias bool) *ebiten.DrawTrianglesOptions {
	Opt := &ebiten.DrawTrianglesOptions{}
	Opt.AntiAlias = antiAlias
	if Fill {
		Opt.FillRule = ebiten.EvenOdd
	}

	return Opt
}

type pathDrawer struct {
	vs *[]ebiten.Vertex
	is *[]uint16
}

func retPathDrawer(path *vector.Path, Fill bool, nonFillLineSize float32, Color int, alpha float64) (pd *pathDrawer) {
	var vs []ebiten.Vertex
	var is []uint16
	if Fill {
		vs, is = (path).AppendVerticesAndIndicesForFilling(nil, nil)
	} else {
		op := &vector.StrokeOptions{}
		op.Width = nonFillLineSize
		op.LineJoin = vector.LineJoinRound
		vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op)
	}

	var r64, g64, b64 float64
	ColorDismantlingIM(&r64, &g64, &b64, Color)
	r, g, b := float32(r64), float32(g64), float32(b64)
	for i := range vs {
		vs[i].DstX = (vs[i].DstX)
		vs[i].DstY = (vs[i].DstY)
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = r / float32(0xff)
		vs[i].ColorG = g / float32(0xff)
		vs[i].ColorB = b / float32(0xff)
		// vs[i].ColorR = 0xfe / float32(0xff)
		// vs[i].ColorG = 0xad / float32(0xff)
		// vs[i].ColorB = 0xffffffff / float32(0xff)
		vs[i].ColorA = float32(alpha)
	}
	return &pathDrawer{&vs, &is}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func StdShowOptPri(x float64, y float64, lineSpace float64, Color int, alpha float32) *text.DrawOptions {
	Opt := &text.DrawOptions{}
	Opt.GeoM.Translate(x, y)
	Opt.LineSpacing = lineSpace

	var r64, g64, b64 float64
	ColorDismantlingIM(&r64, &g64, &b64, Color)
	// r, g, b := float32(r64), float32(g64), float32(b64)

	// Opt.ColorScale.SetR(r)
	// Opt.ColorScale.SetG(g)
	// Opt.ColorScale.SetB(b)

	// Opt.ColorScale.SetA(alpha)

	// if Color != 0 {
	// 	fmt.Printf("<<%x %x\n", int(Opt.ColorScale.B()), int(b))
	// }
	// Opt.ColorScale.ScaleWithColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	Opt.ColorScale.SetA(0xff)
	var textColor color.RGBA
	textColor.R, textColor.G, textColor.B, textColor.A = uint8(r64), uint8(g64), uint8(b64), uint8(alpha)
	Opt.ColorScale.ScaleWithColor(textColor)
	// Opt.ColorScale.SetA(alpha)
	return Opt
}

func PrintShow(screen *ebiten.Image, Opt *text.DrawOptions, format string, afont text.Face) {

	text.Draw(screen, format, afont, Opt)
}

func PrintShowStd(screen *ebiten.Image, format string, afont text.Face, x float64, y float64, lineSpace float64, Color int, alpha float32) {

	PrintShow(screen, StdShowOptPri(x, y, lineSpace, Color, alpha), format, afont)
}

func PrintShowOrg(screen *ebiten.Image, format string, afont text.Face, x float64, y float64, lineSpace float64) {
	Opt := &text.DrawOptions{}
	Opt.GeoM.Translate(x, y)
	Opt.LineSpacing = lineSpace
	PrintShow(screen, Opt, format, afont)
}

const DoubleSize = 1

// フォントロード
func GetNewFont(ff *text.Face, fname string, fsiz float64, fdpi float64) {
	fsiz /= DoubleSize
	fdpi *= DoubleSize
	ftBinary, err := ioutil.ReadFile(fname)
	tt, err := opentype.Parse(ftBinary)
	if err != nil {
		log.Fatal(err)
	}

	fff, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fsiz,
		DPI:     fdpi,
		Hinting: font.HintingVertical,
	})

	*ff = text.NewGoXFace(fff)
	if err != nil {
		log.Fatal(err)
	}

}

func GetNewFontBin(ff *text.Face, ftBinary []byte, fsiz float64, fdpi float64) {
	fsiz /= DoubleSize
	fdpi *= DoubleSize
	tt, err := opentype.Parse(ftBinary)
	if err != nil {
		log.Fatal(err)
	}

	fff, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fsiz,
		DPI:     fdpi,
		Hinting: font.HintingVertical,
	})

	*ff = text.NewGoXFace(fff)
	if err != nil {
		log.Fatal(err)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// rgb 16進数表示のcolorをr成分、g成分、b成分に分解します。
func ColorDismantlingIM(r *float64, g *float64, b *float64, color int) {
	var rr, gg, bb int
	rr = color / int(math.Pow(0x10, 4))
	*r = float64(rr)
	gg = (color / int(math.Pow(0x10, 2))) - (rr * int(math.Pow(0x10, 2)))
	*g = float64(gg)
	bb = (color / int(math.Pow(0x10, 0))) - (rr * int(math.Pow(0x10, 4))) - (gg * int(math.Pow(0x10, 2)))
	*b = float64(bb)

	// fmt.Println(*r, *g, *b)
}

func baseRad(baseX, baseY, addX, addY, rad float64) (x, y float32) {
	// r := math.Sqrt(addX*addX + addY*addY)
	// //rad += math.Atan(addY / addX) //* (math.Abs(addX) / addX)
	// xx := baseX + r*math.Cos(rad+math.Acos(addX/r))
	// yy := baseY + r*math.Sin(rad+math.Asin(addY/r))
	//  fmt.Println(xx, yy, r, rad, addX, addY)
	s, c := math.Sin(rad), math.Cos(rad)
	xx := baseX + (addX * c) - (addY * s)
	yy := baseY + (addX * s) + (addY * c)

	return float32(xx), float32(yy)
}
