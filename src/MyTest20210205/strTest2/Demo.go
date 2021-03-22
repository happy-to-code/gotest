package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	for i := 0; i < 5; i++ {
		if i == 0 {
			s += fmt.Sprintf("%d:%d", i, i+100)
		} else {
			s += fmt.Sprintf("%s%d:%d", "#FG#", i, i+100)
		}
	}
	fmt.Println(s)

	ledgerId := "PPPPP"
	sql := fmt.Sprintf("replace into subcli_%s%s%s%s", ledgerId, "(ID,Date,Height)values(?,?,IFNULL((select Height from subcli_", ledgerId, " where ID =?),1))")
	fmt.Println(sql)

	fmt.Println("______________________________________")
	var str = "p3foukma37:0#FG#p3fghkma37:12"
	splitStr := strings.Split(str, "#FG#aa")
	fmt.Println("len:", len(splitStr))
	fmt.Println("splitStr:", splitStr)
	for _, s2 := range splitStr {
		split := strings.Split(s2, ":")
		fmt.Println(split)

	}

}
