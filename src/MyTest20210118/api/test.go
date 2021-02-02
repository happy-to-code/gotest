package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("20210129.log.gz", ".gz"))

	fn := filepath.Join("/var/logs/", "20210129.log")
	fmt.Println(fn)
}
