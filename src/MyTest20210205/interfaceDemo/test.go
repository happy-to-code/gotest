package main

import (
	"fmt"
	"sync"
	"time"
)

// EventHub 事件中心接口
type EventHub interface {
	RunServer()
	CreateTBByLedgerID(path, ledgerId string)
	SendMsg(et string)
}

type mqttEvent struct {
	lock    sync.Mutex
	MsgID   uint16
	UID     uint64
	Enabled bool
	Addr    string
}

func (mq *mqttEvent) RunServer() {
	fmt.Println("--------RunServer-----------")
}

func (mq *mqttEvent) CreateTBByLedgerID(path, ledgerId string) {
	fmt.Println("--------CreateTBByLedgerID-----------", path, ":", ledgerId)
}

func (mq *mqttEvent) SendMsg(et string) {
	fmt.Println("--------SendMsg-----------", et)
}

var e EventHub

func NewEventHub(addr string, dbPath string, qos uint8) {
	fmt.Printf("eeeeeeeeeeeeeee%seeeeeeeeeeee%seeee:%d", addr, dbPath, qos)
	if e == nil {
		e = &mqttEvent{Addr: addr}
	}
}
func main() {
	NewEventHub("addr", "a/b", 8)
	time.Sleep(3)
}
