package main

import "fmt"

type IMessage interface {
	Print()
}
type BaseMessage struct {
	msg string
}

func (message *BaseMessage) Print() {
	fmt.Println("baseMessage:msg", message.msg)
}

type SubMessage struct {
	BaseMessage
}

func (message *SubMessage) Print() {
	fmt.Println("subMessage:msg", message.msg)
}

func interface_use(i IMessage) {
	i.Print()
}

func main() {

	baseMessage := new(BaseMessage)
	baseMessage.msg = "base"
	interface_use(baseMessage)

	subMessage := new(SubMessage)
	subMessage.msg = "sub"
	interface_use(subMessage)

	fmt.Println("----------------------------")
	iMessage := make([]IMessage, 0, 10)
	iMessage = append(iMessage, baseMessage)
	iMessage = append(iMessage, subMessage)

	for i, msg := range iMessage {
		fmt.Print(i)
		interface_use(msg)
	}
}
