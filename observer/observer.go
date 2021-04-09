package observer

import "fmt"

type Observer interface {
	Notify(state string)
}

type DBObserver struct {
}

func (dbo *DBObserver) Notify(state string) {
	fmt.Println("数据库观察者收到通知：", state)
}

type FrontEndObserver struct {
}

func (feo *FrontEndObserver) Notify(state string) {
	fmt.Println("前端观察者收到通知：", state)
}

type Task struct {
	observers []Observer
	state     string
}

func NewTask(observers []Observer) *Task {
	return &Task{
		observers: observers,}
}

func (t *Task) ChangeState(state string) {
	t.state = state

	fmt.Println("任务进入", state, "状态")
	for _, observer := range t.observers {
		observer.Notify(state)
	}
}

func Run() {

	dbObserver := &DBObserver{}
	frontEndObserver := &FrontEndObserver{}
	task := NewTask([]Observer{dbObserver, frontEndObserver})
	task.ChangeState("开始状态")
	task.ChangeState("结束状态")

}
