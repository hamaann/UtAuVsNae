package chara

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/show"
)

type CharaFace struct {
	FaceImg []*ebiten.Image
	Face    int
	x, y    float64
}

func (cf *CharaFace) View(sc *ebiten.Image) {
	show.ShowImgStd(sc, cf.FaceImg[cf.Face], 1, cf.x, cf.y, 0, false, 1, 1, 1)
}

func (cf *CharaFace) Set(x, y float64, imgs ...*ebiten.Image) {
	cf.x, cf.y = x, y
	cf.Face = 0
	cf.FaceImg = imgs
}

func (cf *CharaFace) SetXY(x, y float64) {
	cf.x, cf.y = x, y

}

func (cf *CharaFace) AddXY(dx, dy float64) {
	cf.x += dx
	cf.y += dy

}

type UtChara struct {
	View      *CharaFace
	ATK, DEF  int
	HP, maxHP int
}

func (c *UtChara) SetStatus(atk, def, hp int) {
	c.maxHP, c.HP = hp, hp
	c.ATK, c.DEF = atk, def
}
