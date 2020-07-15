package main

import (
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  接口继承
 * @Date 2020/7/15 17:55
 * @Param
 * @return
 */
type Action interface {
	Sing()
}

type Cat struct {
}
type Dog struct {
}

func (*Cat) Sing() {
	fmt.Println("Cat is singing")
}
func (*Dog) Sing() {
	fmt.Println("Dog is singing")
}
func Asing(a Action) {
	a.Sing()
}
func main() {
	cat := new(Cat)
	dog := new(Dog)
	Asing(cat)
	Asing(dog)
}
