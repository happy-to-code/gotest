package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s := "G46bHt+UIlEXVqpCFsonYgpKd9+KmiYhS7bkq4wQkXM="
	encodeToString := hex.EncodeToString([]byte(s))
	fmt.Printf("%s", encodeToString)
}
