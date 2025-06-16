package item

import (
	"fmt"

	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/ktui"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
)

// var ItemStatusMessages = textdata.GetTextData("material/data/itemMess.csv")

type ItemUT int

const (
	Nullitem = ItemUT(iota)
	Hotcake
	BadMemory
	Candy
	IceBall
	BSPai

	Allitems
)

var itemList [8]ItemUT

var AllItemSet [Allitems]*Itemthings

var (
	eating sound.SoundF
	Heal   sound.SoundF
	Hurt   sound.SoundF
)

func InitItem() {
	for i := range itemList {
		itemList[i] = Nullitem
	}

	sound.LoadSound(&eating, "material/sounds/se/snd_swallow.wav")
	sound.LoadSound(&Heal, "material/sounds/se/snd_heal_c.wav")
	sound.LoadSound(&Hurt, "material/sounds/se/snd_hurt1_c.wav")
	LoadItemsTxt()
}

func LoadItemsTxt() {
	var itm map[string][]*textdata.Uttxt
	textdata.GetTextData(&itm, "material/data/itemMess.csv", -1)

	setAllItemSet(Hotcake, 100, "ホットケーキ", itm)
	setAllItemSet(BadMemory, -1, "いやなおもいで", itm)
	setAllItemSet(Candy, 10, "モンスターあめ", itm)
	setAllItemSet(IceBall, 45, "ゆきだるまのかけら", itm)
	setAllItemSet(BSPai, 10000, "バタースコッチパイ", itm)

}

func setAllItemSet(num ItemUT, healPower int, name string, list map[string][]*textdata.Uttxt) {

	AllItemSet[num] = &Itemthings{
		Name:  name,
		Dotxt: (list)[name],
		Heal:  healPower,
	}
}

type Itemthings struct {
	Name  string
	Dotxt []*textdata.Uttxt
	Heal  int
}

func ItemEffeHp(listnum int, Hp *int, HpMax int) (txt []*textdata.Uttxt, iname ItemUT) {
	itm := itemList[listnum]
	ItemEffeTxtInit(itm)
	food := AllItemSet[itm]
	hm := food.Heal
	*Hp += hm
	UseItem(listnum)
	if *Hp <= 0 {
		*Hp = HpMax
	}
	var abHP string
	switch ktui.Lang {
	case ktui.Eng:
		abHP = fmt.Sprintf("You recovered %d HP!", hm)
		if hm < 0 {
			abHP = fmt.Sprintf("You decreased %d HP.", -hm)
			EatingSound(Hurt)
		} else {
			EatingSound(Heal)
		}
		if *Hp >= HpMax {
			*Hp = HpMax
			abHP = "Your HP was maxed out."
		}
	default:
		abHP = fmt.Sprintf("HPが　%dかいふくした！", hm)
		if hm < 0 {
			abHP = fmt.Sprintf("HPが　%dへった。", -hm)
			EatingSound(Hurt)
		} else {
			EatingSound(Heal)
		}
		if *Hp >= HpMax {
			*Hp = HpMax
			abHP = "HPが　まんタンになった。"
		}
	}

	// turnInfo.Mestexts[turnInfo.ForItems[itm-1][ktui.Lang]] = (ItemName(itm).Dotxt + "\n" + abHP)
	// food.Dotxt[0].Text=append((food.Dotxt[0].Text),byte('\n'))
	food.Dotxt[0].Text = food.Dotxt[0].Text + "\n" + abHP
	return food.Dotxt, itm
}

func ItemEffeTxtInit(itm ItemUT) {
	// itm := itemList[listnum]
	food := AllItemSet[itm]
	// fmt.Println("first>>", food.Dotxt[0].Text)
	rnTx := []rune(food.Dotxt[0].Text)
	for i := range rnTx {
		// fmt.Printf("%d %d ", rnTx[len(rnTx)-i-1], '\n')
		// fmt.Println(rnTx[len(rnTx)-i-1] == '\n')
		if rnTx[len(rnTx)-i-1] == '\n' {
			// fmt.Println("del>>", rnTx[(len(rnTx)-i-1):])
			rnTx = rnTx[:(len(rnTx) - i - 1)]
			food.Dotxt[0].Text = string(rnTx)
			return
		}
	}
}

func EatingSound(wavsn sound.SoundF) {
	pp := sound.PlayBGM(eating, "wav")
	go func(m sound.Player, wavs *sound.SoundF) {
		for sound.Check(m) {
		}
		sound.PlaySE(*wavs, "wav")
	}(pp, &wavsn)
}

func ItemKey(key int) (item ItemUT) {
	return itemList[key]
}

func AddItem(item ItemUT) {
	for i := range itemList {
		if itemList[i] == Nullitem {
			itemList[i] = item
			break
		}
	}
}

func UseItem(itemnum int) {
	itemList[itemnum] = Nullitem
	there := 0
	for i := range itemList {
		if itemList[i] != Nullitem {
			itemList[there] = itemList[i]
			there++
		}
	}
	for there < len(itemList) {
		itemList[there] = Nullitem
		there++
	}
}

func HowItems() int {
	for i := range itemList {
		if itemList[i] == Nullitem {
			return i
		}
	}
	return len(itemList)
}

// func ItemName(item int) Itemthings {
// 	// name := ""

// 	switch item {
// 	case Hotcake:
// 		if ktui.LcheckEng() {
// 			return Itemthings{"Pancake", "You ate the Pancake.", 100}
// 		}
// 		return Itemthings{"ホットケーキ", "ホットケーキを　たべた。", 100}
// 	case BadMemory:
// 		if ktui.LcheckEng() {
// 			return Itemthings{"BadMemory", "You consume the Bad Memory.", -1}
// 		}
// 		return Itemthings{"いやなおもいで", "いやなおもいでを　のみこんだ。", -1}
// 	case Candy: //MnstrCndy
// 		if ktui.LcheckEng() {
// 			return Itemthings{"MnstrCndy", "You ate the Monster Candy.", 10}
// 		}
// 		return Itemthings{"モンスターあめ", "モンスターあめを　たべた。", 10}
// 	case IceBall:
// 		if ktui.LcheckEng() {
// 			return Itemthings{"SnowPiece", "You ate the Snowman Piece.", 45}
// 		}
// 		return Itemthings{"ゆきだるまのかけら", "ゆきだるまのかけらを　たべた。", 45}
// 	case BSPai: //ButtsPie
// 		if ktui.LcheckEng() {
// 			return Itemthings{"ButtsPie", "You ate the Butterscotch Pie.", 1000}
// 		}
// 		return Itemthings{"バタースコッチパイ", "バタースコッチパイを　たべた。", 1000}
// 	default:
// 		return Itemthings{"", "", 0}
// 	}

// }
