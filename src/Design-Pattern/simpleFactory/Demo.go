package simpleFactory

import "fmt"

type IRuleConfigParse interface {
	Parse(date []byte)
}

type JsonRuleConfigParse struct {
}

type YamlConfigParse struct {
}

func (j JsonRuleConfigParse) Parse(d []byte) {
	fmt.Println("jsonParse")
}

func (y YamlConfigParse) Parse(d []byte) {
	fmt.Println("yamlParse")
}

func NewIRuleConfigParse(t string) IRuleConfigParse {
	switch t {
	case "json":
		return JsonRuleConfigParse{}
	case "yaml":
		return YamlConfigParse{}
	}
	return nil

}
