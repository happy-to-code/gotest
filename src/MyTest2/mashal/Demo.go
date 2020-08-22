package main

import (
	"encoding/json"
	"fmt"
)

// 企业信息和合约地址结构体
type EnterpriseInfoAndContractAddress struct {
	ContractAddress    string `binding:"required"` // 合约地址
	EquityCode         string `binding:"required"` // 股权代码
	EquityAbbreviation string `binding:"required"` // 股权简称
}

func main() {
	var infoAndContractAddress EnterpriseInfoAndContractAddress
	infoAndContractAddress.ContractAddress = "0x0001"
	infoAndContractAddress.EquityCode = "10005"
	infoAndContractAddress.EquityAbbreviation = "xyz"

	fmt.Println("info:", infoAndContractAddress)

	bytes, err := json.Marshal(infoAndContractAddress)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	fmt.Printf("Bytes:%+v\n", bytes)

	fmt.Printf("String(bytes)11111::::==>%s\n", bytes)
	fmt.Println("String(bytes)2222::::==>", string(bytes))

}
