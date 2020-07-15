package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {

	i := uint64(123456789)
	fmt.Println(i)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)

	fmt.Println(b[:])
	i = uint64(binary.BigEndian.Uint64(b))
	fmt.Println("i:::", i)

	var num uint64 = 1278900
	toByte := Uint64ToByte(num)
	fmt.Println("toByte:", toByte)
}

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
