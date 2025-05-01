package avatarobject

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/avatar"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/show/damagedraw"

	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/hmutil"
)

func (fob *Fightings) OutpStatus() string {
	var ot = ""
	for i := range fob.status {
		if fob.status[i] != nil && fob != nil {
			ot += fmt.Sprintln(fob.status[i].modeName, fob.status[i].counter)
		}
	}
	return ot
}

type AvatarObject struct {
	x, y, z    float64
	dir, sign  int
	objectView func(order avatar.Integrity, screen *ebiten.Image, x0, y0, z0 float64, drection, signal int)
}

func defaultObject(x0, y0, z0 float64, dir0 int, sign0 int) *AvatarObject {
	// ob = &AvatarObject{}
	// ob.x, ob.y, ob.z = x0, y0, z0
	// ob.dir = dir0
	// ob.objectView = func(order avatar.Integrity, screen *ebiten.Image, x0, y0, z0 float64, drection, signal int) {

	// }
	// ob.sign = sign0

	ob := AvatarObject{
		x0, y0, z0, dir0, sign0, func(order avatar.Integrity, screen *ebiten.Image, x0, y0, z0 float64, drection, signal int) {},
	}
	return &ob
}

// avatar移動
func (ob *AvatarObject) Move(dx, dy, dz float64) {
	ob.x += dx
	ob.y += dy
	ob.z += dz

}

// 方向設定
func (ob *AvatarObject) SetDir(d int) {
	ob.dir = d
}

// 外見設定
func (ob *AvatarObject) SetView(f func(order avatar.Integrity, screen *ebiten.Image, x0, y0, z0 float64, drection, signal int)) {
	ob.objectView = f
	// fmt.Println(ob)
}

// avater描画
func (ob *AvatarObject) View(order avatar.Integrity, screen *ebiten.Image) {
	ob.objectView(order, screen, ob.x, ob.y, ob.z, ob.dir, ob.sign)
}

// 描画用信号送信
func (ob *AvatarObject) PushSign(signal int) {
	ob.sign = signal
}

func (ob *AvatarObject) Where() (x, y float64) {
	return ob.x, ob.y
}

func (ob *AvatarObject) Which() (dir int) {
	return ob.dir
}

func (ob *AvatarObject) LookedCamera(height float64) {
	avatar.CameraSet(ob.x, ob.y, height)
}

func (ob *AvatarObject) MoveD(dir int, sp float64) {
	switch dir {
	case hmutil.UP:
		ob.Move(0, -sp, 0)
	case hmutil.DOWN:
		ob.Move(0, sp, 0)
	case hmutil.LEFT:
		ob.Move(-sp, 0, 0)
	case hmutil.RIGHT:
		ob.Move(sp, 0, 0)
	}
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Effector func(target *Fightings, thisDir int) (delete bool)
type Bullet struct {
	Obj          *AvatarObject
	Targets      []*Fightings
	bulletEffect Effector
	BulletRange  func(x0, y0, targX, targY float64, counter int, rotate float64) bool
	BulletMove   func(Obj *AvatarObject, counter int) (delete bool)
}

func SetBullet(t ...*Fightings) *Bullet {
	return &Bullet{
		defaultObject(0, 0, 0, 0, 0), t, func(target *Fightings, thisDir int) (delete bool) { return false }, func(x0, y0, targX, targY float64, counter int, rotate float64) bool { return false }, func(Obj *AvatarObject, counter int) (delete bool) { return false },
	}

}

func (bb *Bullet) Run(x, y, z float64, dir int) {
	ob := *bb.Obj
	bs := Bullet{}
	bs = *bb
	b := &bs
	b.Obj = &ob
	b.Obj.x, b.Obj.y, b.Obj.z = x, y, z
	b.Obj.dir = dir

	for i := range bulletsRunningInTheAvataerWorld {
		if bulletsRunningInTheAvataerWorld[i] == nil {
			bulletsRunningInTheAvataerWorld[i] = b
			return
		}
	}
	bulletsRunningInTheAvataerWorld = append(bulletsRunningInTheAvataerWorld, b)
}

func (b *Bullet) SetEffect(ef Effector) {
	b.bulletEffect = ef
}

func (b *Bullet) touchAndEffect() (delete bool) {
	for i := range b.Targets {
		if b.Targets[i] == nil {
			continue
		}
		t := b.Targets[i]
		if b.BulletRange(b.Obj.x, b.Obj.y, t.Obj.x, t.Obj.y, b.Obj.sign, hmutil.DirToR(hmutil.DOWN, b.Obj.dir)) {
			if b.bulletEffect(t, b.Obj.dir) {
				return true
				// Fin(b)
			}
		}
	}

	return false
}

var bulletsRunningInTheAvataerWorld []*Bullet

func AllBulletRunning() {
	for i := range bulletsRunningInTheAvataerWorld {
		// setSliceFilter(i)
		if bulletsRunningInTheAvataerWorld[i] == nil {
			continue
		}
		b := *bulletsRunningInTheAvataerWorld[i]
		if b.BulletMove(b.Obj, b.Obj.sign) || b.touchAndEffect() {
			bullFinN(i)
			// bulletsRunningInTheAvataerWorld[i] = nil
			continue
		}

		b.Obj.sign++

	}
}

// func Fin(b *Bullet) {
// 	b.BulletMove = func(Obj *AvatarObject, counter int) (delete bool) {
// 		return true
// 	}
// }

func bullFinN(num int) {
	bulletsRunningInTheAvataerWorld[num] = nil
}

func AllBulletView(order avatar.Integrity, screen *ebiten.Image) {
	for i := range bulletsRunningInTheAvataerWorld {
		if bulletsRunningInTheAvataerWorld[i] == nil {
			continue
		}
		b := *bulletsRunningInTheAvataerWorld[i]
		b.Obj.View(order, screen)
	}
}

func MakeEffectorGiveDamege(damage int, damageRate float64, atk int, backPower, backPhaseLong, voidPhaseLong, stopPhaseLong int, touchDead bool) Effector {
	return func(target *Fightings, thisDir int) (delete bool) {
		direction := thisDir
		target.Obj.SetDir(hmutil.InvDir(direction))
		target.Damaged(hmutil.DamCalculationStd(damage, damageRate, atk, target.def), direction, backPower, backPhaseLong, voidPhaseLong, stopPhaseLong)
		return touchDead
	}
}

func MakeEffectorGiveDamegeWithShow(x, y, z float64, damage int, damageRate float64, atk int, backPower, backPhaseLong, voidPhaseLong, stopPhaseLong int, touchDead bool) Effector {
	return func(target *Fightings, thisDir int) (delete bool) {
		direction := thisDir
		target.Obj.SetDir(hmutil.InvDir(direction))
		dam := hmutil.DamCalculationStd(damage, damageRate, atk, target.def)

		if target.Damaged(dam, direction, backPower, backPhaseLong, voidPhaseLong, stopPhaseLong) {
			damagedraw.DamageGo(dam, x, y, z)
		}
		return touchDead
	}
}

func AllBulletObjectsOutput() int {
	return len(bulletsRunningInTheAvataerWorld)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Status Object
type condition struct {
	counter  int
	modeName string
}

// Fightings Object
type Fightings struct {
	Obj              *AvatarObject
	hp, hpMax        int
	speed            float64
	atkBase, defBase int
	atk, def         int
	status           []*condition
}

func MakeFighter(x0, y0, z0 float64, dir0 int, sign0 int, hp int, speed float64, atk, def int) *Fightings {
	f := &Fightings{}
	f.Obj = defaultObject(x0, y0, z0, dir0, sign0)
	f.hpMax, f.hp = hp, hp
	f.speed = speed
	f.atk, f.def = atk, def
	f.atkBase, f.defBase = atk, def
	return f
}

func (fob *Fightings) ShowStatus() (atk, def, hp int) {
	return fob.atk, fob.def, fob.hp
}

func (fob *Fightings) ShowStatusBase() (atk, def, hp int) {
	return fob.atkBase, fob.defBase, fob.hpMax
}

func (fob *Fightings) Buffed(atkBf, defBf int, phaseTime int) {
	fob.atk += atkBf
	fob.def += defBf

	if atkBf > 0 {
		fob.Grant(fmt.Sprintf("buff_ a %d", atkBf), phaseTime)
	}
	if atkBf < 0 {
		fob.Grant(fmt.Sprintf("debuff_ a %d", atkBf), phaseTime)
	}

	if defBf > 0 {
		fob.Grant(fmt.Sprintf("buff_ d %d", defBf), phaseTime)
	}
	if defBf < 0 {
		fob.Grant(fmt.Sprintf("debuff_ d %d", defBf), phaseTime)
	}
}

// Fightingsへの効果の有効化
func (fob *Fightings) StatusEffecting() {

	for i := -1; fob.allStatusLoop(&i); {
		if fob.status[i] == nil {
			continue
		}
		switch fob.status[i].modeName {
		case "dashing":
			dashSize := float64(3)
			fob.MoveFob(fob.Obj.dir, fob.speed*dashSize)
		default:
			fob.whenHurt(i)
			fob.whenBuffed(i)
		}
	}

	fob.Counting()
}

// 死亡チェック
func (fob *Fightings) DeadCheck() bool {
	return fob.hp <= 0
}

// HP回復
func (fob *Fightings) Heal(healingHelth int) {
	fob.hp += healingHelth
	if fob.hp > fob.hpMax {
		fob.hp = fob.hpMax
	}
}

// 被ダメージ＆ノックバック状態付与
func (fob *Fightings) Damaged(damage int, direction int, backPower, backPhaseLong, voidPhaseLong, stopPhaseLong int) (Effected bool) {
	if fob.InTheStatusNear("void") == -1 {
		fob.hp -= damage
		if fob.hp < 0 {
			fob.hp = 0
		}
		fob.Grant(conHurt(backPower, backPhaseLong, direction))
		fob.Grant("void", backPhaseLong+voidPhaseLong)
		fob.Grant("stop", backPhaseLong+voidPhaseLong+stopPhaseLong)

		return true
	}

	return false
}

// hurtingのステータス作成
func conHurt(backPower, backLong int, dir int) (statusName string, statusTime int) {
	return fmt.Sprintf("hurting %c %d", directionToRune(dir), backPower), backLong
}

// ノックバック挙動
// テンプレート：hurting U 10
func (fob *Fightings) whenHurt(stnm int) {
	// if strings.Index(fob.status[stnm].modeName, "hurting") == 0 {

	// fmt.Println(fob.modeNameHead("hurting", stnm))
	if fob.modeNameHead("hurting", stnm) {
		st, dir, power := "", ' ', 0
		rd := fob.statusReader(stnm)
		fmt.Fscanf(rd, "%s %c %d", &st, &dir, &power)
		// fmt.Println(dir, power)
		// fmt.Fscanf(strings.NewReader(fob.status[stnm].modeName), "%s %s %d", nil, &dir, &power)
		// bc := len("hurting")
		switch dir {
		case 'U':
			fob.MoveFob(hmutil.UP, float64(power))
		case 'D':
			fob.MoveFob(hmutil.DOWN, float64(power))
		case 'L':
			fob.MoveFob(hmutil.LEFT, float64(power))
		case 'R':
			fob.MoveFob(hmutil.RIGHT, float64(power))
		}
	}
}

func (fob *Fightings) whenBuffed(stnm int) {
	if fob.status[stnm].counter <= 2 {
		// fmt.Println(fob.status[stnm].modeName)
		if fob.modeNameHead("buff_", stnm) {
			st := ""
			w, buffing := ' ', 0
			fmt.Fscanf(fob.statusReader(stnm), "%s %c %d", &st, &w, &buffing)
			switch w {
			case 'a':
				fob.atk -= buffing
				// fmt.Println(fob.atk)
			case 'd':
				fob.def -= buffing
			}
			fob.FinishN(stnm)
			return
		}

		if fob.modeNameHead("debuff_", stnm) {
			st := ""
			w, buffing := ' ', 0
			fmt.Fscanf(fob.statusReader(stnm), "%s %c %d", &st, &w, &buffing)
			switch w {
			case 'a':
				fob.atk += buffing
			case 'd':
				fob.def += buffing
			}
			fob.FinishN(stnm)
			return
		}
	}
}

func (fob *Fightings) modeNameHead(nameHead string, statusNumber int) bool {
	// s := "hur"
	// if statusNumber == 0 {
	// 	fmt.Println(fob.status[statusNumber].modeName, nameHead, strings.Index(fob.status[statusNumber].modeName, nameHead))
	// }
	return strings.Index(fob.status[statusNumber].modeName, nameHead) == 0
}

func (fob *Fightings) statusReader(statusNumber int) *strings.Reader {
	// fmt.Println(fob.status[statusNumber].modeName)
	return strings.NewReader(fob.status[statusNumber].modeName)
}

// 操作に応じた移動
func (fob *Fightings) MoveStd(dir int) {
	if fob.InTheStatusNear("dashing", "stop") == -1 {
		fob.Obj.SetDir(dir)
		fob.MoveFob(dir, fob.speed)
	}
}

// 移動
func (fob *Fightings) MoveFob(dir int, sp float64) {
	// sp *= fob.speed
	switch dir {
	case hmutil.UP:
		fob.Obj.Move(0, -sp, 0)
	case hmutil.DOWN:
		fob.Obj.Move(0, sp, 0)
	case hmutil.LEFT:
		fob.Obj.Move(-sp, 0, 0)
	case hmutil.RIGHT:
		fob.Obj.Move(sp, 0, 0)
	}
}

// 状態付与
func (fob *Fightings) Grant(statusName string, statusTime int) {
	for i := range fob.status {
		if fob.status[i] == nil {
			fob.status[i] = &condition{statusTime, statusName}
			return
		}
	}
	fob.status = append(fob.status, &condition{statusTime, statusName})
}

// 状態強制終了
func (fob *Fightings) Finish(statusName string) {
	w := fob.InTheStatusAll(statusName)
	for n := range w {
		fob.FinishN(w[n])
		// fob.status[w[n]] = nil
	}
}

func (fob *Fightings) FinishN(statusNumber int) {
	fob.status[statusNumber] = nil
}

// 状態カウント
func (fob *Fightings) Counting() {
	for i := range fob.status {
		if fob.status[i] != nil {
			if fob.status[i].counter > 0 {
				fob.status[i].counter--
			} else {
				fob.status[i] = nil
			}
		}
	}
}

// 最近対象ステータス探索
func (fob *Fightings) InTheStatusNear(statusName ...string) (where int) {
	for i := range fob.status {
		for j := range statusName {
			if fob.status[i] != nil {
				if fob.status[i].modeName == statusName[j] {
					return i
				}
			}
		}
	}
	return -1
}

// 全対象ステータス探索
func (fob *Fightings) InTheStatusAll(statusName string) (where []int) {
	w := make([]int, 0)
	for i := range fob.status {
		if fob.status[i] != nil {
			if fob.status[i].modeName == statusName {
				w = append(w, i)
			}
		}
	}
	return w
}

func (fob *Fightings) allStatusLoop(n *int) bool {
	*n++
	return len(fob.status) > *n
}

func directionToRune(dir int) rune {
	s := ' '
	switch dir {
	case hmutil.UP:
		s = 'U'
	case hmutil.DOWN:
		s = 'D'
	case hmutil.LEFT:
		s = 'L'
	case hmutil.RIGHT:
		s = 'R'
	}

	return s
}
