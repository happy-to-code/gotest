package main

import (
	"fmt"
	"strings"
)

const LEDGERID_PREFIX = "ledgerId_index_"

func main() {

	var s = LEDGERID_PREFIX + "jjdjfgj"
	fmt.Println("s:", s)
	fmt.Println("s2:", strings.Split(s, LEDGERID_PREFIX)[1])
	fmt.Println("==========================")
	sql := fmt.Sprintf("%s%s%s", "update subcli_", "ent.LedgerID", " set Height = ? where ID = ? ")

	fmt.Println(sql)
}
