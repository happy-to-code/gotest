package main

import (
	"encoding/binary"
)

func main() {
	// 保存 int16 数据
	i := uint16(233)

	// 将 int16 转换为 byte 数据，并输出
	b := Int16ToBytes(i)
	println(b)

	// 输出 byte 转换后 int16 数据
	println(BytesToInt16(b))
}

func Int16ToBytes(i uint16) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToInt16(buf []byte) uint16 {
	return uint16(binary.BigEndian.Uint16(buf))
}
