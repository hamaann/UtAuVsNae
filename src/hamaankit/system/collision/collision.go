package collision

import "math"

// ベースあり箱型
func BoxColB(targX, targY, baseX, baseY, addX, addY, wide, high, rad float64) bool {
	baseX, baseY = BaseRad(baseX, baseY, addX, addY, rad)
	return BoxCol(targX, targY, baseX, baseY, wide, high, rad)
}

// ベースあり楕円型
func OvalColB(targX, targY, baseX, baseY, addX, addY, wide, high, rad float64) bool {
	baseX, baseY = BaseRad(baseX, baseY, addX, addY, rad)
	return OvalCol(targX, targY, baseX, baseY, wide, high, rad)
}

// 用途不明
func BaseRad(baseX, baseY, addX, addY, rad float64) (x, y float64) {
	/*r := math.Sqrt(addX*addX + addY*addY)
	xx := baseX + r*math.Cos(rad)
	yy := baseY + r*math.Sin(rad)*/

	s, c := math.Sin(rad), math.Cos(rad)
	xx := baseX + (addX * c) - (addY * s)
	yy := baseY + (addX * s) + (addY * c)
	return xx, yy
}

// チューブ型
func TubeCol(targX float64, targY float64, rad float64, coreX float64, coreY float64, rangeOn float64, rangeUnd float64) bool {

	if rad > math.Pi/2 {
		for rad > math.Pi/2 {
			rad -= math.Pi
		}
	}

	if rad < -math.Pi/2 {
		for rad < -math.Pi/2 {
			rad += math.Pi
		}
	}

	f := math.Tan(rad)*(targX-coreX) + coreY

	// if !(f+(rangeOn/math.Cos(rad)) >= targY && targY >= f-(rangeUnd/math.Cos(rad))) {
	// 	fmt.Println(f+(rangeOn/math.Cos(rad)), targY, f-(rangeUnd/math.Cos(rad)))
	// }
	return f+(rangeOn/math.Cos(rad)) >= targY && targY >= f-(rangeUnd/math.Cos(rad))
}

// 箱型
func BoxCol(targX float64, targY float64, x float64, y float64, wide float64, high float64, rad float64) bool {
	//fmt.Println(TubeCol(targX, targY, rad, x, y, high/2, high/2) && TubeCol(targX, targY, rad+(math.Pi/2), x, y, wide/2, wide/2))
	return TubeCol(targX, targY, rad, x, y, high/2, high/2) && TubeCol(targX, targY, rad+(math.Pi/2), x, y, wide/2, wide/2)
}

// 円型
func RoundCol(targX float64, targY float64, x float64, y float64, dstRange float64) bool {
	a := targX - x
	b := targY - y
	return math.Sqrt((a*a)+(b*b)) <= dstRange
}

// 楕円型
func OvalCol(targX float64, targY float64, x float64, y float64, wide float64, high float64, rad float64) bool {
	c, s := math.Cos(rad), math.Sin(rad)
	x_ := ((targX-x)*c + (targY-y)*s) / (wide / 2)
	y_ := (-(targX-x)*s + (targY-y)*c) / (high / 2)
	return x_*x_+y_*y_ <= 1
}
