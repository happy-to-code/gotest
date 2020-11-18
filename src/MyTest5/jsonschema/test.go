package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader("E:\\20.06.16Project\\GoTest\\src\\MyTest5\\jsonschema\\file\\数据对照表v2.0.0-Alpha2.xlsx")
	documentLoader := gojsonschema.NewReferenceLoader("E:\\20.06.16Project\\GoTest\\src\\MyTest5\\jsonschema\\file\\数据对照表v2.0.0-Alpha2.xlsx")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
