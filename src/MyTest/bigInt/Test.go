package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Block struct {
	Nonce     uint64
	TimeStamp uint64
	Data      []byte
	Hash      []byte
}

func main() {
	var block Block
	block.Nonce = 9800
	block.TimeStamp = uint64(time.Now().Unix())
	block.Data = []byte("HelloWorld")
	block.Hash = []byte("123456789")

	fmt.Printf("blockInfo:%+v\n", block)

	tem := [][]byte{
		Uint64ToByte(block.Nonce),
		Uint64ToByte(block.TimeStamp),
		block.Data,
		block.Hash,
	}
	var blockInfo []byte
	blockInfo = bytes.Join(tem, []byte{})

	sum256 := sha256.Sum256(blockInfo)
	fmt.Println("sum256:", sum256)
	fmt.Printf("sum256:%x", sum256)

	fmt.Println("=================")
	var sz [2]int = [2]int{1, 2}
	fmt.Printf("sz的类型和值：%s,%v\n", reflect.TypeOf(sz), sz)
	fmt.Printf("sz[:]的类型和值：%s,%v\n", reflect.TypeOf(sz[:]), sz[:])

}

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
