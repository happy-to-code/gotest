package main

import (
	"fmt"
	"reflect"
	"sync"
)

var sysCache sync.Map

func main() {
	sysCache.Store("equityCode_SH1005", []byte("516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85"))
	sysCache.Store("equityCode_SH1006", []byte("1006"))

	fmt.Printf("sysCache:%+v\n", sysCache)
	fmt.Println("[]byte()::::", []byte("516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85"))
	fmt.Println("================================")
	fromSysCache := GetValueFromSysCache("equityCode_SH1005558")
	fmt.Printf("==>%+v%v\n", fromSysCache, len(fromSysCache))

	sysCache.Delete("equityCode_SH1005")
	fromSysCache2 := GetValueFromSysCache("equityCode_SH100555")
	fmt.Printf("2==>%+v\n", fromSysCache2)

}
func GetValueFromSysCache(key string) []byte {
	value, ok := sysCache.Load(key)
	if ok {
		fmt.Println("value:", value, reflect.TypeOf(value))
		fmt.Println("value:", value.([]byte), reflect.TypeOf(value.([]byte)))
		return value.([]byte)
	}
	return nil
}
