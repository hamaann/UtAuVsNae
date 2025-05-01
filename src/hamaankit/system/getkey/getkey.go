package getkey

import "github.com/hajimehoshi/ebiten/v2"

type InpKey struct {
	key *[]ebiten.Key

	Pushed bool
}

func (ky *InpKey) Check() {
	ky.Pushed = inputMix(*ky.key...)

}

func (ky *InpKey) Off() {
	ky.Pushed = false

}

func MakeInpKey(Keys ...ebiten.Key) *InpKey {
	k := &InpKey{}
	k.key = &Keys
	k.Pushed = false

	return k
}

func KeysInput(inpk ...*InpKey) {
	for i := range inpk {
		inpk[i].Check()
	}
}

func inputMix(Keys ...ebiten.Key) bool {
	ky := false
	for i := range Keys {
		ky = ky || ebiten.IsKeyPressed(Keys[i])
	}

	return ky
}

func inputOne(Key ebiten.Key) bool {
	return ebiten.IsKeyPressed(Key)
}

// var allkey []*InpKey

// func MakeInpKeyWithAllKeys(Keys ...ebiten.Key) *InpKey {
// 	k := MakeInpKey(Keys...)
// 	allkey = append(allkey, k)
// 	return k
// }

// func KeysInputAll() {
// 	KeysInput(allkey...)
// }
