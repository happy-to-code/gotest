package main

import "fmt"

type PolicyType int32

const (
	Policy_MIN PolicyType = 0
	Policy_MAX PolicyType = 1
	Policy_MID PolicyType = 2
	Policy_AVG PolicyType = 3
)

func foo(p PolicyType) {
	fmt.Printf("enum value: %v\n", p)
}
func (p PolicyType) Str() string {
	switch p {
	case Policy_MIN:
		return "MIN"
	case Policy_MAX:
		return "MAX"
	case Policy_MID:
		return "MID"
	case Policy_AVG:
		return "AVG"
	default:
		return "UNKNOWN"
	}
}

func main() {
	foo(Policy_MAX)

	var policyType = Policy_AVG
	str := policyType.Str()
	fmt.Printf("str: %s\n", str)
}
