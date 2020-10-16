package main

import (
	"fmt"
	"strings"
)

type BalanceAndAmount struct {
	Amount      int64 // 总数量==》转让前的总数量
	FzAmount    int64 // 冻结的数量
	AfterAmount int64 // 转让后的数量
}

func main() {
	var a int64 = 0
	fmt.Println(a)
	var b uint64 = 5
	fmt.Println(a - int64(b))

	m := make(map[string]int64)
	m["aft_SmzBcqpKCtyZoJLUdXGS9jh2goo96Yc9p7nifsh6mVsACgemEzk_3688_0"] = 920000000
	m["bal_SmzBcqpKCtyZoJLUdXGS9jh2goo96Yc9p7nifsh6mVsACgemEzk_3688_0"] = 910000000
	m["unl_SmzBcqpKCtyZoJLUdXGS9jh2goo96Yc9p7nifsh6mVsACgemEzk_3688_0"] = 0

	m["aft_SoCiYcW1rXUnHxBPxCqv4mCwatnCenpn3T4b3hUkZXGtmuWHsq7_3688_0"] = 980000000
	m["bal_SoCiYcW1rXUnHxBPxCqv4mCwatnCenpn3T4b3hUkZXGtmuWHsq7_3688_0"] = 990000000
	m["unl_SoCiYcW1rXUnHxBPxCqv4mCwatnCenpn3T4b3hUkZXGtmuWHsq7_3688_0"] = 100000000

	bkMap := make(map[string]BalanceAndAmount)
	for k, v := range m {
		var hasExist = false
		preStr, addr := getAddr(k)
		for address, balanceAndAmount := range bkMap {
			if addr == address {
				hasExist = true
				if preStr == "bal" {
					balanceAndAmount.Amount += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "unl" {
					balanceAndAmount.FzAmount += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "aft" {
					balanceAndAmount.AfterAmount += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else {
					continue
				}

			}

		}
		if !hasExist {
			var b BalanceAndAmount
			if preStr == "bal" {
				b.Amount = v / 1000000
			} else if preStr == "unl" {
				b.FzAmount = v / 1000000
			} else if preStr == "aft" {
				b.AfterAmount = v / 1000000
			}
			bkMap[addr] = b
		}

	}
	fmt.Println("----------------------")
	fmt.Printf("bkMap:%+v\n", bkMap)

}
func getAddr(s string) (preStr, addr string) {
	preStr = s[:3]

	index0 := strings.Index(s, "_")
	s2 := s[index0+1:]
	index2 := strings.Index(s2, "_")
	s3 := s2[:index2]
	return preStr, s3
}
