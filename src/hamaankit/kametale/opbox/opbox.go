package opbox

import (
	"math"
)

const allPattbox = 3

// x1,x2,y1,y2
var PattBox = [allPattbox][4]float64{
	{30, 30 + 575, 250, 250 + 140},
	{240, 640 - 230, 225, 250 + 140},
	{0, 640, 0, 480},
}

var (
	Shifting [2][2]float64 //
	shiflong [2]float64
	shiftime int
)

const (
	cos = iota
	sin
)

func ShiftBox(now [4]float64, how [4]float64, time int) {
	a, b := (how[x1] - now[x1]), (how[y1] - now[y1])

	Shifting[0][cos] = a / math.Sqrt(a*a+b*b)
	Shifting[0][sin] = b / math.Sqrt(a*a+b*b)
	shiflong[0] = math.Sqrt(a*a+b*b) / float64(time)

	a, b = (how[x2] - now[x2]), (how[y2] - now[y2])
	// fmt.Println(b, a, b/a, math.Atan(b/a), math.Cos(-0), math.Sin(-0))
	Shifting[1][cos] = a / math.Sqrt(a*a+b*b)
	Shifting[1][sin] = b / math.Sqrt(a*a+b*b)
	shiflong[1] = math.Sqrt(a*a+b*b) / float64(time)

	shiftime = time

	if a*a+b*b == 0 {
		for i := range Shifting {
			for j := range Shifting[i] {
				Shifting[i][j] = 0
			}
		}
	}
	// fmt.Println(Shifting, shiflong, (how[y1]-now[y1])/(how[x1]-now[x1]), b/a, shiflong[1]*Shifting[1][0])
}

func ShiftBoxChenger(shiftthings *[4]float64) {
	if shiftime > 0 {
		shiftthings[x1] += shiflong[0] * Shifting[0][cos] //math.Cos(Shifting[0])
		shiftthings[y1] += shiflong[0] * Shifting[0][sin] //math.Sin(Shifting[0])
		shiftthings[x2] += shiflong[1] * Shifting[1][cos] //math.Cos(Shifting[1])
		shiftthings[y2] += shiflong[1] * Shifting[1][sin] //math.Sin(Shifting[1])
		shiftime--
	}
}

const (
	x1 = iota
	x2
	y1
	y2
)

func plmn(x float64) float64 {
	if x == 0 {
		return 1
	}

	return x / math.Abs(x)
}

func opAtan(x1 float64, x2 float64) float64 {
	n := math.Atan(x1 / x2)
	if x1 == 0 {
		n = plmn(x2) * math.Pi / 2
	}

	return n

}
