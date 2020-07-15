package main

import (
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  多重继承
 * @Date 2020/7/15 17:55
 * @Param
 * @return
 */
type Father struct {
	MingZi string
}

func (this *Father) Say() string {
	return "大家好，我叫 " + this.MingZi
}

type Mother struct {
	Name string
}

func (this *Mother) Say() string {
	return "Hello, my name is " + this.Name
}

type Child struct {
	Father
	Mother
}

func main() {
	c := new(Child)
	c.MingZi = "小明"
	c.Name = "xiaoming"
	fmt.Println(c.Father.Say())
	fmt.Println(c.Mother.Say())
}
