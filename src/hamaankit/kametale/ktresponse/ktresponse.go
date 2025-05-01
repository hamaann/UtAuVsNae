package ktresponse

var actProgress int

func AddActProg(delta int) {
	actProgress += delta
}

func GetNowProg() int {
	return actProgress
}

type Act struct {
	name   string
	Ans    *Anser
	effect func()
}

func (act *Act) SetEffect(f func()) {
	act.effect = f
}

func NewAct(actName string, actAns *Anser, actEffect func()) *Act {
	return &Act{
		name:   actName,
		Ans:    actAns,
		effect: actEffect,
	}

}

type Anser struct {
	message []string
	serif   []string
}

func (ans *Anser) NewMessage(m ...string) {
	ans.message = m
}

func (ans *Anser) NewSerif(s ...string) {
	ans.serif = s
}

func (ans *Anser) ResetMessage() {
	ans.message = make([]string, 0)
}

func (ans *Anser) ResetSerif() {
	ans.serif = make([]string, 0)
}
