package main

import "fmt"

func main() {
	var boy1 = Boy{
		Name: "xiaoming",
		Age:  12,
	}
	var boy2 = Boy{
		Name: "zhanghua",
		Age:  18,
	}

	var cls = Class{
		ClassName: "高一3班",
		Boys:      []Boy{boy1, boy2},
	}

	fmt.Printf("cls:%+v\n", cls)
	fmt.Println("------------------------")
	fmt.Printf("cls.Boys:%+v\n", cls.Boys)
	for i, v := range cls.Boys {
		fmt.Println(i)
		fmt.Println(cls.Boys[i])
		fmt.Println("v:::::", v)
	}

}

type Boy struct {
	Name string
	Age  int64
}
type Class struct {
	ClassName string
	Boys      []Boy
}
