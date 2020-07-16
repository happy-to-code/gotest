package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// sessionId函数用来生成一个session ID，即session的唯一标识符
func sessionId() []byte {
	b := make([]byte, 32)
	// ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	// rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return []byte{}
	}
	fmt.Println("b:", b)
	return b
}
func main() {
	idBytes := sessionId()
	fmt.Printf("idBytes :%x\n", idBytes)
	fmt.Println("bytes===", hex.EncodeToString(idBytes))
}
