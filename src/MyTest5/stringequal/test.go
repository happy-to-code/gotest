package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("HELLO", "hello"))
	fmt.Println(strings.EqualFold("Subject_create_time", "subject_create_time"))
	fmt.Println(strings.EqualFold("Subject_create_time", "Subject_Create_time"))
	// fmt.Println(strings.EqualFold("ÑOÑO", "ñoño"))
}
