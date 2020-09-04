package main

import (
	"fmt"
)

type Boy struct {
	Name string
	Age  int32
	X    string
}

func main() {
	// var mymap map[string][]Boy
	mymap := make(map[string][]Boy, 2)

	var b1 []Boy
	b1 = append(b1, Boy{"xiaoMing", 17, ""})

	var b2 []Boy
	b2 = append(b2, Boy{"Hong", 18, ""})
	b2 = append(b2, Boy{"Hua", 16, ""})
	mymap["aaa"] = b1
	mymap["ccc"] = b2

	for k, v := range mymap {
		fmt.Println(k, ":\t", v)
		for i, boy := range v {
			boy.X = k
			v[i] = boy
		}
	}

	fmt.Println("-------------")
	fmt.Printf("%v", mymap)

	fmt.Println("===========|")
	var a []string
	// a = append(a, "aa")
	// a = append(a, "bb")
	// a = append(a, "cc")

	fmt.Printf("%v\n", a)

	fmt.Print(contain("a", a))
	fmt.Print(contain("aa", a))
	fmt.Print(contain("33", a))

}
func contain(a string, b []string) bool {
	for _, s := range b {
		if s == a {
			return true
		}
	}
	return false
}
