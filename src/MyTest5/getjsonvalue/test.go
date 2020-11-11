package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Subject struct {
	Name string
	Age  int
	Role []string
}

func main() {
	subject := Subject{
		Name: "AA",
		Age:  16,
		Role: []string{"r1", "r2"},
	}

	fmt.Printf("%+v\n", subject)
	subjectMarshal, _ := json.Marshal(subject)
	fmt.Println(string(subjectMarshal))
	replaceStr := strings.Replace(string(subjectMarshal), "r1", "r3", 1)
	fmt.Println(replaceStr)
	fmt.Println("json串转map")
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(replaceStr), &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	fmt.Println("m:", m)

}
