package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
)

func main() {
	im, err := gg.LoadImage("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo4\\bg.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>", im.Bounds())

}
