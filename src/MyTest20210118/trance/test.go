package main

import (
	"os"
	"runtime/trace"
)

// go run main.go 2>trace.go
// go tool trace trace.out
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "YIDA"
	}()

	<-ch
}
