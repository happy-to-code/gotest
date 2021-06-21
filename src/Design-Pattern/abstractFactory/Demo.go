package abstractFactory

import "fmt"

type IRuleConfigParse interface {
	Parse(data []byte)
}

type JsonRuleConfigParse struct{}

func (j JsonRuleConfigParse) Parse(d []byte) {
	fmt.Println("json Parse:", d)
}

type ISystemConfigParse interface {
	ParseSystem(data []byte)
}
type JsonSystemConfigParse struct{}

func (j JsonSystemConfigParse) ParseSystem(d []byte) {
	fmt.Println("json system parse:", d)
}

// IConfigParserFactory 工厂方法接口
type IConfigParseFactory interface {
	CreateRuleParse() IRuleConfigParse
	CreateSystemParse() ISystemConfigParse
}

type JsonConfigParseFactory struct{}

func (j JsonConfigParseFactory) CreateRuleParse() IRuleConfigParse {
	return JsonRuleConfigParse{}
}

func (j JsonConfigParseFactory) CreateSystemParse() ISystemConfigParse {
	return JsonSystemConfigParse{}
}
