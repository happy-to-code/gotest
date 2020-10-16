package main

import (
	"fmt"
)

type BalanceAndAmount struct {
	AmountWithType      int64 // 总数量==》转让前的总数量(区分股权性质)
	AfterAmountWithType int64 // 转让后的数量(区分股权性质)
	FzAmount            int64 // 冻结的数量
	TotalAmountAfter    int64 // 转让后的数量(不区分 股权性质)
	TotalAmountBefore   int64 // 转让后的数量(不区分 股权性质)
}

func main() {

	var b = BalanceAndAmount{
		AmountWithType:      20,
		AfterAmountWithType: 30,
		FzAmount:            0,
		TotalAmountAfter:    0,
		TotalAmountBefore:   0,
	}

	var b2 = BalanceAndAmount{
		AmountWithType:      380,
		AfterAmountWithType: 370,
		FzAmount:            0,
		TotalAmountAfter:    0,
		TotalAmountBefore:   0,
	}

	m := make(map[string]BalanceAndAmount, 0)
	m["SnBLwhWVVQHiSJ1CkfNUet1aos1vUVCGHFEADjHL2VQpUotSmJg"] = b
	m["SocyiHxxPT4Sp1zfk4S1vtdBitkSXeoYtudDc2V8uRN873W2bnf"] = b2

	fmt.Println(m)
	fmt.Println("---------------------")
	// 转出方地址
	var countFromAddr string
	// 接收方地址
	var countToAddr string
	for k, balanceAndAmount := range m {
		if (balanceAndAmount.AmountWithType - balanceAndAmount.AfterAmountWithType) > 0 {
			countFromAddr = k
		} else {
			countToAddr = k
		}
	}
	fmt.Printf("转出方地址:%s,接收方地址:%s\n", countFromAddr, countToAddr)

}
