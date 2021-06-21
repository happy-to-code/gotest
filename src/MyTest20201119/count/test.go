package main

import (
	"fmt"
	"unicode"
)

// HanCounter to count the number of chinese character.
func HanCounter(s string) float64 {
	var hanCount float64 // 0
	var letterCount int  // 0
	var numberCount int  // 0
	var spaceCount int   // 0
	var PunctCount int   // 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			hanCount++
		} else if unicode.Is(unicode.Letter, c) {
			letterCount++
		} else if unicode.Is(unicode.Number, c) {
			numberCount++
		} else if unicode.Is(unicode.Punct, c) {
			PunctCount++
		} else if unicode.Is(unicode.Space, c) {
			spaceCount++
		}
	}
	fmt.Println(letterCount)
	fmt.Println("numberCount:", numberCount)
	fmt.Println("spaceCount:", spaceCount)
	fmt.Println("PunctCount:", PunctCount)
	return hanCount
}

func main() {
	s := "Hello,“” 世界1234%$#^"
	count := HanCounter(s)
	fmt.Println(count)
}
