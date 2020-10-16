package main

import (
	"fmt"
	"strconv"
)

type AddressInfo struct {
	Address       string `binding:"required"` // 用户地址
	Amount        uint64 `binding:"gt=0"`     // 回收的数量
	ShareProperty uint16 `binding:"gte=0"`    // 股权类别
}
type AddrAndAmount struct {
	Address string
	Amount  uint64
}

func main() {
	var a = 0
	b := string(a)
	fmt.Println("0" == b)
	fmt.Println("0" == strconv.Itoa(a))
	fmt.Println("000000000000000000000000000000000000000000")
	var info1 = AddressInfo{"aa", 10, 1}
	var info2 = AddressInfo{"aa", 5, 2}
	var info3 = AddressInfo{"b1b", 20, 1}
	var info4 = AddressInfo{"bb", 20, 1}
	var addressInfos []AddressInfo
	addressInfos = append(addressInfos, info1)
	addressInfos = append(addressInfos, info2)
	addressInfos = append(addressInfos, info3)
	addressInfos = append(addressInfos, info4)

	fmt.Printf("addressInfos:%+v\n", addressInfos)
	var addrAndAmountList []AddrAndAmount

	for _, info := range addressInfos {
		var hasExist = false
		for i, amount := range addrAndAmountList {
			if info.Address == amount.Address {
				amount.Amount += info.Amount
				addrAndAmountList[i] = amount
				hasExist = true
			}
		}
		if !hasExist {
			addrAndAmountList = append(addrAndAmountList, AddrAndAmount{info.Address, info.Amount})
		}
	}
	fmt.Printf("addrAndAmountList:%+v\n", addrAndAmountList)
	fmt.Println("=========================")
	var decreaseShareholderNum = 0
	fmt.Println("decreaseShareholderNum:", decreaseShareholderNum)
	fmt.Printf("decreaseShareholderNum Type:%T\n", decreaseShareholderNum)
	for i := 0; i < 3; i++ {
		decreaseShareholderNum--
	}
	fmt.Println("decreaseShareholderNum:", decreaseShareholderNum)

}
