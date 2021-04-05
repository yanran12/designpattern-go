package command

import "fmt"

type Command interface {
	Invoke(tv *TV)
}

type TV struct {
}

func (tv *TV) Receive(command Command) {
	command.Invoke(tv)
}

type OpenCommand struct {
}

func (openCommand *OpenCommand) Invoke(tv *TV) {
	fmt.Println("打开电视")
}

type CloseCommand struct {
}

func (closeCommand *CloseCommand) Invoke(tv *TV) {
	fmt.Println("关闭电视")
}

type ChannelOneCommand struct {
}

func (channelOneCommand *ChannelOneCommand) Invoke(tv *TV) {
	fmt.Println("切换到1频道")
}

func Run() {

	openCommand := &OpenCommand{}
	channelOneCommand := &ChannelOneCommand{}
	closeCommand := &CloseCommand{}

	tv := &TV{}

	tv.Receive(openCommand)
	tv.Receive(channelOneCommand)
	tv.Receive(closeCommand)
}
