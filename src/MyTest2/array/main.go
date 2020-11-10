package main

import (
	"fmt"
	"strconv"
)

type UTXO struct {
	Amount    int64
	TokenType string
	SubT      string
	Address   string
}

func main() {

	var u = UTXO{10, "10005", "0", "0001"}
	if u.TokenType == "10005" {
		u.TokenType = "xddd"
	}
	fmt.Println("------", u)

	// strconv.Itoa(bb.json)
	a, _ := strconv.Atoi("bb.json")
	fmt.Printf("%T,%d\n:", a, a)

	var utxos []UTXO

	utxos = append(utxos, UTXO{10, "10005", "0", "0001"})
	utxos = append(utxos, UTXO{20, "30005", "bb.json", "0021"})
	utxos = append(utxos, UTXO{30, "10005", "3", "0101"})
	utxos = append(utxos, UTXO{40, "30005", "bb.json", "0801"})

	fmt.Printf("utxos:::%v\n", utxos)
	fmt.Println("---------------------------")

	var newUtxos []UTXO
	for _, u := range utxos {
		isDump := false
		for j, nu := range newUtxos {
			if u.TokenType == nu.TokenType && u.SubT == nu.SubT {
				nu.Amount += u.Amount
				newUtxos[j] = nu
				fmt.Println(nu.Amount)
				isDump = true
			}
		}
		if !isDump {
			newUtxos = append(newUtxos, u)
		}

	}
	fmt.Printf("newUtxos:%v\n", newUtxos)
}
