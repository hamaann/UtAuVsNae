package textdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultTextPerSpeed = 2

type Uttxt struct {
	Who         string
	Text        string
	RunePerTick int
}

func GetTextData(pathName string) (list *map[string][]*Uttxt) {
	// var list *[]string
	var texts map[string][]*Uttxt
	fp, err := os.Open(pathName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	loadUttexts(scanner, &texts)
	return &texts
}

func loadUttexts(sc *bufio.Scanner, list *map[string][]*Uttxt) {
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
			if len(t) >= 3 {
				rpt, _ = strconv.Atoi(t[2])

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
			(*list)[listname] = append((*list)[listname], &Uttxt{
				Who:         t[0],
				Text:        t[1],
				RunePerTick: rpt,
			})
		}
	}
}

// func loadStrings(sc *bufio.Scanner, list *[]string) {
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
