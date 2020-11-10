package main

type B struct {
	Name string
	Age  int
}

func (b B) getName() string {
	return "xiaoming"
}
