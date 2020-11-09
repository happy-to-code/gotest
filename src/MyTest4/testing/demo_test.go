package main

import "testing"

func Test_jsonTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
