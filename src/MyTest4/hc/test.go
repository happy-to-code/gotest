package main

import "fmt"

var sysCache = make(map[string][]byte, 110)

func main() {
	sysCache["equityCode_1"] = []byte("516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85")
	sysCache["contractAddress_1"] = []byte("123asdasd")
	sysCache["userAddress_1"] = []byte("456qqxsad")

	fmt.Printf("sysCache:%s\n", sysCache)
	fmt.Printf("========================\n")
	fmt.Println("===>", sysCache["equityCode"])
	delete(sysCache, "userAddress")
	fmt.Printf("sysCache:%s\n", sysCache)
	fmt.Println("===>", sysCache["equityCode111"])
	fmt.Println("===>", string(sysCache["equityCode111"]) == "")
	fmt.Println("===>", string(sysCache["equityCode111"]) == " ")
	fmt.Println("===>", len(sysCache["equityCode111"]) <= 0)
	fmt.Println("--------------------------------------")
	cache := getValueFromSysCache(1, "11")
	fmt.Println("cache:", string(cache))
	fmt.Println("cache:", len(cache))
}

func getValueFromSysCache(tp int, value string) []byte {
	var byBack []byte
	if tp == 1 {
		byBack = sysCache["equityCode_"+value]
		return byBack
	}
	if tp == 2 {
		byBack = sysCache["constantAddr_"+value]
		return byBack
	}
	if tp == 3 {
		byBack = sysCache["userAddr_"+value]
		return byBack
	}
	return nil
}
