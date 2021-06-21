package main

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServiceWithoutAuth(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
	}{
		{name: "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
