package main

import (
	"fmt"
	"strings"
)

func main() {

	var b = make(map[string]int64)

	b["aft_SmrQESKdMerLmnQYrFh5WJ6P4o4ix11fSwqUoZUsnKKw2Naz77V_T00043_0"] = 0
	b["aft_Sn2FPu5nj1D77CgYiT3CmAbNiBhtui61LwBmG5P1U7gBQ1F1waA_T00043_0"] = 350000000000
	b["alA_SmrQESKdMerLmnQYrFh5WJ6P4o4ix11fSwqUoZUsnKKw2Naz77V_T00043"] = 0
	b["alA_Sn2FPu5nj1D77CgYiT3CmAbNiBhtui61LwBmG5P1U7gBQ1F1waA_T00043"] = 350000000000
	b["alB_SmrQESKdMerLmnQYrFh5WJ6P4o4ix11fSwqUoZUsnKKw2Naz77V_T00043"] = 200000000000
	b["alB_Sn2FPu5nj1D77CgYiT3CmAbNiBhtui61LwBmG5P1U7gBQ1F1waA_T00043"] = 150000000000
	b["bal_SmrQESKdMerLmnQYrFh5WJ6P4o4ix11fSwqUoZUsnKKw2Naz77V_T00043_0"] = 200000000000
	b["bal_Sn2FPu5nj1D77CgYiT3CmAbNiBhtui61LwBmG5P1U7gBQ1F1waA_T00043_0"] = 150000000000
	b["com_T00043"] = 0
	b["unl_SmrQESKdMerLmnQYrFh5WJ6P4o4ix11fSwqUoZUsnKKw2Naz77V_T00043_0"] = 0
	b["unl_Sn2FPu5nj1D77CgYiT3CmAbNiBhtui61LwBmG5P1U7gBQ1F1waA_T00043_0"] = 0

	m := convertMap(b)
	fmt.Println("=========================================")
	fmt.Printf("%+v\n", m)

}

type BalanceAndAmount struct {
	AmountWithType      int64 // 总数量==》转让前的总数量(区分股权性质)
	AfterAmountWithType int64 // 转让后的数量(区分股权性质)
	FzAmount            int64 // 冻结的数量
	TotalAmountAfter    int64 // 转让后的数量(不区分 股权性质)
	TotalAmountBefore   int64 // 转让后的数量(不区分 股权性质)
}

// 将map转换成  地址：BalanceAndAmount 形式
func convertMap(b map[string]int64) map[string]BalanceAndAmount {
	fmt.Printf("======：：：%v\n", b)
	if len(b) <= 0 {
		return nil
	}
	bkMap := make(map[string]BalanceAndAmount)
	for k, v := range b {
		var hasExist = false
		// 字符串开始的3个字母
		beginStr := k[:3]
		if beginStr != "aft" && beginStr != "bal" && beginStr != "unl" && beginStr != "alA" && beginStr != "alB" {
			continue
		}
		preStr, addr := getAddr(k)
		for address, balanceAndAmount := range bkMap {
			if addr == address {
				hasExist = true
				if preStr == "bal" {
					balanceAndAmount.AmountWithType += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "unl" {
					balanceAndAmount.FzAmount += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "aft" {
					balanceAndAmount.AfterAmountWithType += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "alA" {
					balanceAndAmount.TotalAmountAfter += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else if preStr == "alB" {
					balanceAndAmount.TotalAmountBefore += v / 1000000
					bkMap[addr] = balanceAndAmount
				} else {
					continue
				}

			}

		}
		if !hasExist {
			var b BalanceAndAmount
			if preStr == "bal" {
				b.AmountWithType = v / 1000000
			} else if preStr == "unl" {
				b.FzAmount = v / 1000000
			} else if preStr == "aft" {
				b.AfterAmountWithType = v / 1000000
			} else if preStr == "alA" {
				b.TotalAmountAfter = v / 1000000
			} else if preStr == "alB" {
				b.TotalAmountBefore = v / 1000000
			}
			bkMap[addr] = b
		}

	}
	return bkMap
}
func getAddr(s string) (preStr, addr string) {
	// fmt.Println("========>", s)
	preStr = s[:3]
	index0 := strings.Index(s, "_")

	s2 := s[index0+1:]

	index2 := strings.Index(s2, "_")

	s3 := s2[:index2]
	return preStr, s3
}
