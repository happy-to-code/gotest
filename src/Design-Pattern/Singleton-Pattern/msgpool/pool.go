package msgpool

import "sync"

// 消息池
type messagePool struct {
	pool *sync.Pool
}

// 消息
type Message struct {
	Count int
}

// 消息池单例
var msgPool = &messagePool{
	// 如果消息池里没有消息，则新建一个Count值为0的Message实例
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}

// 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 往消息池里添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// 从消息池里获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
