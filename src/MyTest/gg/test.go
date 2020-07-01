package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"unicode/utf8"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	s := "18932156666"

	w, h := dc.MeasureString(s)
	fmt.Println("w:", w)
	fmt.Println("h:", h)
	fmt.Println("len():", len(s))

	count := utf8.RuneCountInString(s)
	fmt.Println("count:", count)

	dc.SavePNG("out.png")
}
