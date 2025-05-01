package showkage

import (
	img "image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type KageType struct {
	shader    *ebiten.Shader
	optionSet KageOpt
}

func (k *KageType) Load(kageFile *[]byte) {
	LoadKage(&k.shader, kageFile)
}

func (k *KageType) Setting(option KageOpt) {
	k.optionSet = option
}

func (k *KageType) ShowImgKageType(screen *ebiten.Image, Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, sizeX float64, sizeY float64) {
	ShowImgKageStd(screen, k.optionSet, k.shader, Pic, size, x, y, rad, alpha, setUnder, sizeX, sizeY)
}

type KageOpt func(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image) (op *ebiten.DrawRectShaderOptions)

func ShowImgKageStd(screen *ebiten.Image, optf KageOpt, shader *ebiten.Shader, Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, sizeX float64, sizeY float64) {
	ShowImgKage(screen, optf(StdShowOptKage(Pic, size, x, y, rad, alpha, setUnder, sizeX, sizeY), Pic), shader, Pic)
}

func ShowImgKage(screen *ebiten.Image, opt *ebiten.DrawRectShaderOptions, shader *ebiten.Shader, Pic *ebiten.Image) {
	screen.DrawRectShader(Pic.Bounds().Dx(), Pic.Bounds().Dy(), shader, opt)
}

func StdShowOptKage(Pic *ebiten.Image, size float64, x float64, y float64, rad float64, alpha float64, setUnder bool, sizeX float64, sizeY float64) *ebiten.DrawRectShaderOptions {
	Opt := &ebiten.DrawRectShaderOptions{}
	w, h := Pic.Size()
	// 係数で画像を拡大/縮小したときの大きさを計算しておく
	var sw, sh float64 = float64(w) * sizeX * size, float64(h) * sizeY * size

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

	//  Opt.ColorScale.ScaleAlpha(float32(alpha))

	return Opt
}

func LoadKage(shad **ebiten.Shader, kageF *[]byte) {
	if kageF != nil {
		var err error
		*shad, err = ebiten.NewShader(*kageF)
		if err != nil {
			log.Fatal(err)
		}

		*kageF = nil
	}

}

var randSeeds *ebiten.Image

const RSimgSizeX, RSimgSizeY = 1536, 2048

var (
	TrialFill  *ebiten.Shader
	NoiseFill  *ebiten.Shader
	Test       *ebiten.Shader
	Waving     *ebiten.Shader
	Glitch     *ebiten.Shader
	Grad       *ebiten.Shader
	Aberration *ebiten.Shader
)

// embedで仕入れたファイル達([]byte)ポインタを受け取る。
func LoadAllKage(randSeedImg *ebiten.Image, kages ...*[]byte) {
	// gameX, gameY = Gwidth, Gheight
	if randSeedImg != nil {
		randSeeds = randSeedImg
	}
	LoadKage(&TrialFill, kages[0])
	LoadKage(&NoiseFill, kages[1])
	LoadKage(&Test, kages[2])
	LoadKage(&Waving, kages[3])
	LoadKage(&Glitch, kages[4])
	LoadKage(&Grad, kages[5])
	LoadKage(&Aberration, kages[6])
}

func KageOptFillT(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, color []float64) *ebiten.DrawRectShaderOptions {
	opt.Uniforms = make(map[string]interface{})
	opt.Uniforms["FillColor"] = color
	opt.Images[0] = Pic
	return opt
}
func KageOptFillN(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, seed float32) *ebiten.DrawRectShaderOptions {
	opt.Uniforms = make(map[string]interface{})
	opt.Uniforms["Seed"] = seed
	opt.Images[0] = Pic
	return opt
}

func KageOptTest(screen *ebiten.Image, opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, cellSize float32) *ebiten.DrawRectShaderOptions {
	opt.Images[0] = Pic
	opt.Uniforms = make(map[string]any, 1)
	opt.Uniforms["CellSize"] = cellSize
	return opt
}

// vertShift==1:縦
func KageOptWave(screen *ebiten.Image, opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, vertShift int, size, notWidth, move float64) *ebiten.DrawRectShaderOptions {
	opt.Images[0] = Pic
	opt.Uniforms = make(map[string]any, 1)
	opt.Uniforms["Vertical"] = vertShift
	opt.Uniforms["Size"] = size
	opt.Uniforms["Narrow"] = notWidth
	opt.Uniforms["Phase"] = move
	return opt
}

func KageOptGlich(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, seed float32, shiftMax float32, glHeight int, distanse float32, colorpattern int) *ebiten.DrawRectShaderOptions {
	opt.Uniforms = make(map[string]interface{})
	opt.Uniforms["Seed"] = seed
	// fmt.Println(seed, shiftMax)
	opt.Uniforms["ShiftMax"] = shiftMax
	opt.Uniforms["High"] = glHeight
	opt.Uniforms["Distance"] = distanse
	opt.Uniforms["Pattern"] = colorpattern
	opt.Images[0] = Pic
	baseX, baseY := 0, 0
	PicInfo := Pic.Bounds()
	opt.Images[1] = randSeeds.SubImage(img.Rect(GetRandForImg(&baseX, &baseY, &PicInfo, seed))).(*ebiten.Image)
	// opt.Images[1] = randSeeds.SubImage(img.Rect(baseX, baseY, baseX+PicInfo.Dx(), baseY+PicInfo.Dy())).(*ebiten.Image)
	return opt
}

func GetRandForImg(bx, by *int, pic *img.Rectangle, seed float32) (x0 int, y0 int, x1 int, y1 int) {
	// rand.Seed(time.Now().UnixNano() * 777)
	rand.Seed(int64(seed * 777))
	*bx = rand.Intn(RSimgSizeX - (*pic).Dx())
	// rand.Seed(time.Now().UnixNano() * 888)
	rand.Seed(int64(seed * 888))
	*by = rand.Intn(RSimgSizeY - (*pic).Dy())
	return *bx, *by, *bx + (*pic).Dx(), *by + (*pic).Dy()
}

// ぼかし
func KageOptGrad(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, width float32) *ebiten.DrawRectShaderOptions {
	opt.Uniforms = make(map[string]interface{})
	opt.Uniforms["Width"] = width

	opt.Images[0] = Pic
	return opt
}

func KageOptAberrate(opt *ebiten.DrawRectShaderOptions, Pic *ebiten.Image, distanse float32, colorpattern int) *ebiten.DrawRectShaderOptions {
	opt.Uniforms = make(map[string]interface{})
	opt.Uniforms["Distance"] = distanse
	opt.Uniforms["Pattern"] = colorpattern

	opt.Images[0] = Pic
	return opt
}
