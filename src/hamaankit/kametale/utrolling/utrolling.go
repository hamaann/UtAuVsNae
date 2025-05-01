package utrolling

import (
	"github.com/hamaa/UtAuVsNae/src/hamaankit/kametale/textdata"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/rpgtext"
	"github.com/hamaa/UtAuVsNae/src/hamaankit/system/sound"
)

var (
	showingText   []*textdata.Uttxt
	nowShowingNum int

	rpgShowing    *rpgtext.RpgText
	rpgShowingWho string

	astSignal = [3]bool{false}

	txtVoice sound.SoundF
)

func IsAsts() [3]bool {
	return astSignal
}

var 7seg int

func runText(t *rpgtext.RpgText) string {

	if t.IsRpgTextFin() {
		return t.Sprint()
	} else {

		return t.SprintNewWithFunc(func(s [2]rune, c int) {
			sound.PlaySE(txtVoice, "wav")
			if s == [2]rune{'$', '*'} {
				t.EraseBetween(c, c+1)
				astSignal[t.GetLineNum()] = true
			}
		})
	}
}

func setUttxtToRpgText(txt *textdata.Uttxt) *rpgtext.RpgText {
	rt := &rpgtext.RpgText{}
	rt.SetText(txt.Text, txt.RunePerTick)

	return rt
}

func RollingWho() string {
	return rpgShowingWho
}

func RollingShowingStart(ut []*textdata.Uttxt) {
	showingText = ut
	nowShowingNum = 0
	rpgShowing = setUttxtToRpgText(showingText[nowShowingNum])
	rpgShowingWho = showingText[nowShowingNum].Who
	astSignal = [3]bool{false}
}

func RollingShowingGo(decidekey, skipKey bool) (rollingRpgedText string, fin bool) {
	var showingMax = (nowShowingNum+1 == len(showingText))
	var nextable bool
	if nextable = rpgShowing.IsRpgTextFin(); nextable {
		if decidekey {
			if !showingMax {
				nowShowingNum++
				rpgShowing = setUttxtToRpgText(showingText[nowShowingNum])
				rpgShowingWho = showingText[nowShowingNum].Who
			}
		}
	} else {
		if skipKey {
			rpgShowing.SkipText()
		}
	}
	return runText(rpgShowing), (showingMax && nextable)
}
