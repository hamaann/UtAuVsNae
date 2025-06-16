package textdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RollingTypeName(name string) int {
	switch name {
	case "ctrl":
		return 0
	case "auto":
		return 1
	case "auto_ctrl":
		return 2
	}
	return 0
}

const defaultTextPerSpeed = 2

type Uttxt struct {
	Who         string
	Text        string
	RunePerTick int
	RollingType int
}

func GetTextData(list *map[string][]*Uttxt, pathName string, returnNum int) {
	// var list *[]string
	// var texts map[string][]*Uttxt
	fp, err := os.Open(pathName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	loadUttexts(scanner, list, returnNum) //&texts)

	// return texts
}

func loadUttexts(sc *bufio.Scanner, list *map[string][]*Uttxt, returnNum int) {
	// unicode.IsSpace(',')
	*list = make(map[string][]*Uttxt)
	sc.Split(bufio.ScanLines)
	listname := ""
	// *list = make([]string, 0)
	for sc.Scan() {
		stt := string(sc.Bytes())
		removeUnwantedStrings(&stt)
		t := strings.Split(stt, ",")
		if t[0] == "name" {
			listname = t[1]
			(*list)[listname] = make([]*Uttxt, 0)
		} else if !(t[0] == "" && t[1] == "") {
			rpt := defaultTextPerSpeed
			rltp := 0
			if len(t) >= 3 {
				if t[2] != "" {
					rpt, _ = strconv.Atoi(t[2])
				}
				if len(t) >= 4 {
					rltp = RollingTypeName(t[3])
				}

			}

			// if len(t) > 3 {
			// for i := range t {
			// 	if i < 3 {
			// 		continue
			// 	} else {
			// 		t[2] = t[2] + "\n" + t[i]
			// 	}

			// }
			// }
			if t[0] == "" {
				sz := len((*list)[listname])
				(*list)[listname][sz-1].Text = (*list)[listname][sz-1].Text + "\n" + t[1]
				continue
			}

			t := &Uttxt{
				Who:         t[0],
				Text:        t[1],
				RunePerTick: rpt,
				RollingType: rltp,
			}
			if returnNum > 0 {
				t.TextAutoLine(returnNum)
				// t.TextShiftUtBox(returnNum)
			}

			(*list)[listname] = append((*list)[listname], t)

		}
	}

}

func IsMode(rolltype int, typeName string) bool {
	return RollingTypeName(typeName) == rolltype
}

func (txt *Uttxt) TextAutoLine(alineNum int) {
	looked := 0
	beforsp := 0
	rn := []rune(txt.Text)
	if len(rn) < alineNum {
		return
	}

	// fmt.Printf("spacedData>>")
	for i := 0; i < len(rn); i++ {
		if rn[i] == ' ' || rn[i] == '　' {
			if i-looked < alineNum {
				beforsp = i
				// fmt.Printf("%d ", i)
			} else {
				i = beforsp
				rn[i] = '\n'
				looked = i
				// fmt.Printf("<%d> ", i)
			}
		}

		if rn[i] == '\n' {
			looked = i
		}

		if !(i+1 < len(rn)) {
			if !(i-looked < alineNum) {
				i = beforsp
				rn[i] = '\n'
				looked = i
				// fmt.Printf("{%d} ", i)
			}
		}
	}

	txt.Text = string(rn)
	// fmt.Println("\n", txt.Text)
}

// func (txt *Uttxt) TextShiftUtBox(retNumText int) {
// 	rn := []rune(txt.Text)
// 	if len(rn) < retNumText {
// 		return
// 	}
// 	sp := make([]int, 0)
// 	sk := make([]int, 0)
// 	rn = append(rn, ' ')
// 	for i := range rn {
// 		if rn[i] == ' ' || rn[i] == '　' {
// 			sp = append(sp, i)
// 		}

// 		if rn[i] == '\n' {
// 			sk = append(sk, i)
// 		}
// 	}

// 	looked := 0
// 	added := 0
// 	skipCount := 0
// 	fmt.Println("[txt]>>>", txt.Text, sp, sk)
// 	for j := range sp {
// 		if len(sk) > skipCount {
// 			if sp[j] > sk[skipCount] {
// 				skipCount++
// 				looked = sp[j] + 1
// 			}
// 		}
// 		if sp[j]+added-looked > retNumText {
// 			j--
// 			fmt.Println("	>", sp[j+1], looked, "::", "!", string(rn[:sp[j+1]+added-looked]), sp[j+1]+added-looked)
// 			subrn := []rune{'\n'}
// 			subrn = append(subrn, rn[sp[j]+1+added-looked:]...)
// 			rn = append(rn[:sp[j]+1+added-looked], subrn...)
// 			looked = sp[j] + 1
// 			j++
// 			// j--
// 			// added++

// 		}
// 	}

// 	txt.Text = string(rn)
// 	fmt.Println("[txt:shifted]>>>", txt.Text, added, looked)
// 	return
// }

// func loadStrings(sc *bufio.Scanner, li()st *[]string) {
// 	// unicode.IsSpace(',')
// 	sc.Split(bufio.ScanLines)
// 	*list = make([]string, 0)
// 	for sc.Scan() {
// 		stt := string(sc.Bytes())
// 		removeUnwantedStrings(&stt)
// 		*list = append(*list, stt)
// 	}
// }

func removeUnwantedStrings(st *string) {
	def, af, _ := strings.Cut(*st, string([]byte{239, 187, 191}))
	*st = fmt.Sprintf("%s%s", def, af)
}

func NanUttxt() *Uttxt {
	return &Uttxt{
		Who:         "st",
		Text:        "",
		RunePerTick: 2,
		RollingType: 0,
	}
}
