package main

import (
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  简单继承
 * @Date 2020/7/15 17:51
 * @Param
 * @return
 */
type Father struct {
	MingZi string
}

func (this *Father) Say() string {
	return "大家好，我叫 " + this.MingZi
}

type Child struct {
	Father
}

func main() {
	c := new(Child)
	c.MingZi = "小明"
	fmt.Println(c.Say())
}
