package state

import "fmt"

type Light struct {
	state State
}

func (l *Light) ChangeSwitch() {
	l.state.ChangeSwitch()
}

type State interface {
	ChangeSwitch()
}

type OnState struct {
	light *Light
}

func (ons *OnState) ChangeSwitch() {
	fmt.Println("关灯")
	ons.light.state = &OffState{light: ons.light}
}

type OffState struct {
	light *Light
}

func (offs *OffState) ChangeSwitch() {
	fmt.Println("开灯")
	offs.light.state = &OnState{light: offs.light}
}

func Run() {

	light := &Light{}
	light.state = &OffState{light}
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
	light.ChangeSwitch()
}
