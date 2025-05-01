package hmutil

import (
	"math"
	"math/rand"
	"time"
)

func JstBool(b bool, bjst *int) bool {
	if !b {
		*bjst = 0
	} else if *bjst < 2 {
		*bjst++
	}

	return *bjst == 1
}

const (
	UP = iota
	DOWN
	LEFT
	RIGHT

	ALLDIR
)

func InvDir(dir int) int {
	if dir%2 == 0 {
		return dir + 1
	} else {
		return dir - 1
	}
}

func DirectionToString(dir int) string {
	s := ""
	switch dir {
	case UP:
		s = "back"
	case DOWN:
		s = "front"
	case LEFT:
		s = "left"
	case RIGHT:
		s = "right"
	}

	return s
}

func DirectionAdd(x, y float64, add float64, dir int) (nx, ny float64) {
	switch dir {
	case UP:
		return x, y - add
	case DOWN:
		return x, y + add
	case LEFT:
		return x - add, 0
	case RIGHT:
		return x + add, 0
	}

	return x, y
}

// bmax,bmin：atk=100,def=0の時の最大ダメージ、最小ダメージ
func DamCalculationStd(damage int, DamRate float64, myAtk, yourDef int) int {
	p := DamRate / 100.0
	bmax, bmin := float64(damage)*(1+p), float64(damage)*(1-p)
	h := float64(myAtk) / float64(10*yourDef)
	s := int64(1 + myAtk*yourDef*damage)
	r := float64(EasyRand(int(bmin), int(bmax+1), s))
	base := bmin * (1.0 + (h))

	return int(base + r)
}

// min <= ranum < max
func EasyRand(min, max int, seed int64) (ranum int) {
	rand.Seed(time.Now().UnixNano() * int64(max-min) * (seed + 1))
	return min + rand.Intn(max-min)
}

var rotateDirs = [ALLDIR]int{UP, LEFT, DOWN, RIGHT}

func DirToR(baseDir, nowDir int) float64 {
	var goingRotate = -1
	for i := 0; ; i++ {
		if baseDir == rotateDirs[i%4] && goingRotate == -1 {
			goingRotate = i
		}

		if goingRotate != -1 {
			if nowDir == rotateDirs[i%4] {
				return (math.Pi / 2) * float64(i-goingRotate)
			}
		}
	}
}

func MenuSelecter(PosiBool bool, NegaBool bool, Menum *int, min int, max int, delta int) (shifted bool) {
	sft := false
	// if abreast >= keyparallelmax {
	// 	log.Fatal("overflow from  keyparallelmax ")
	// }

	if PosiBool {
		*Menum = (*Menum + delta) % max
		sft = true

	}

	if NegaBool {
		*Menum -= delta
		if *Menum < min {
			*Menum = max - (min - *Menum)
		}
		sft = true

	}

	return sft
}
