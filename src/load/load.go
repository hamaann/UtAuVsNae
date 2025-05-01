package load

const (
	InRooms = iota
	InFontlib
	InOutward
	InMenuUT
	InBarrages
	InGeno
	InBattle

	AllInitalP
)

var InitalPro [AllInitalP]func(gx, gy float64)

var nowLoading = 0

func LoadedSystem(gx, gy float64) bool {
	if nowLoading >= AllInitalP {
		return true
	}

	InitalPro[nowLoading](gx, gy)
	// time.Sleep(time.Second)
	nowLoading++

	return false
}

func ProgressLoadSystem() float64 {
	return float64(nowLoading) / AllInitalP
}
