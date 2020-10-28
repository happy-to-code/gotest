package main

import "testing"

func Test_isNil(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{""}, want: true},
		{name: "2", args: args{" "}, want: true},
		{name: "3", args: args{"aas12 "}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNil(tt.args.s); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
