package mediator

import "fmt"

type ChatRoom struct {
	name string
}

func (cr *ChatRoom) SendMsg(msg string) {
	fmt.Println(cr.name + " : " + msg)
}
func (cr *ChatRoom) RegisterUser(u *User) {
	u.cr = cr
}

type User struct {
	name string
	cr   *ChatRoom
}

func (u *User) SendMsg(msg string) {
	if u.cr != nil {
		u.cr.SendMsg(u.name + " : " + msg)
	}
}

func Run() {
	AUser := &User{name: "AUser"}
	BUser := &User{name: "BUser"}

	chatRoom := &ChatRoom{name: "chatRoom1"}
	chatRoom.RegisterUser(AUser)
	chatRoom.RegisterUser(BUser)

	AUser.SendMsg("hello")
	BUser.SendMsg("hello")

}
