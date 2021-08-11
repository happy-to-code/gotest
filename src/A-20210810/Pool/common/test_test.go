package main

import (
	"fmt"
	"testing"
)

func Test_createConnection(t *testing.T) {

	got, err := createConnection()
	if err != nil {
		t.Errorf("createConnection() error = %v", err)
		return
	}
	fmt.Println(got)

}
