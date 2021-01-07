package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

func main() {
	str := "MD5testing"
	md5Str := md5V(str)
	fmt.Println(md5Str)
	fmt.Println(md5V2(str))
	fmt.Println(md5V3(str))
}
func md5V(str string) string {
	t1 := time.Now().UnixNano()
	h := md5.New()
	h.Write([]byte(str))
	t2 := time.Now().UnixNano()
	fmt.Println("md5v:", t2-t1)
	return hex.EncodeToString(h.Sum(nil))
}
func md5V2(str string) string {
	t1 := time.Now().UnixNano()
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	t2 := time.Now().UnixNano()
	fmt.Println("md5v2:", t2-t1)
	return md5str
}
func md5V3(str string) string {
	t1 := time.Now().UnixNano()
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	t2 := time.Now().UnixNano()
	fmt.Println("md5v3:", t2-t1)
	return md5str
}
