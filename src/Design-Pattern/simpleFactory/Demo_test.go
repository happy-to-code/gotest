package simpleFactory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParse(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParse
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: JsonRuleConfigParse{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: YamlConfigParse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParse(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
