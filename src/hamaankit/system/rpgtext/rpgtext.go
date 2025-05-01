package rpgtext

type RpgText struct {
	text              []rune
	callingRune, tick int
	fin               bool
	tpRune            int
	line              int
}

func (t *RpgText) IsRpgTextFin() bool {
	return t.fin
}

func (t *RpgText) SetText(tx string, tickPerRune int) {
	t.text = []rune(tx)
	t.callingRune = 0
	t.tick = 0
	t.line = 0
	t.fin = false
	t.tpRune = tickPerRune
}

func (t *RpgText) Sprint() string {
	return string(t.text[:t.callingRune])
}

func (t *RpgText) SprintNew() string {

	return t.SprintNewWithFunc(nil)
}

func (t *RpgText) EraseBetween(delhead, delhip int) {

	t.text = append(t.text[:delhead], t.text[delhip+1:]...)

	if delhead <= t.callingRune && t.callingRune <= delhip {
		t.callingRune = delhead
	}
}

func (t *RpgText) SprintNewWithFunc(f func(willShow [2]rune, callingnum int)) string {
	if !t.IsRpgTextFin() {
		t.tick++
		if t.tick%t.tpRune == 0 {
			if t.callingRune < len(t.text) {

				if f != nil && t.callingRune+1 < len(t.text) {
					f([2]rune{t.text[t.callingRune], t.text[t.callingRune+1]}, t.callingRune)

				}
				t.callingRune++
				if t.callingRune+1 < len(t.text) {
					if t.text[t.callingRune] == '\n' {
						t.callingRune++
						t.line++
					}
				}

			} else {
				t.fin = true
			}
		}
	}
	return t.Sprint()
}

func (t *RpgText) SkipText() {
	t.callingRune = len(t.text)
	t.fin = true
}

func (t *RpgText) GetLineNum() int {
	return t.line
}
