package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type People struct {
	Name   string
	Age    int
	Class  int
	Gender int
	Tools  []string
}

func main() {
	var tools []string = []string{"zs", "dn"}
	var p = People{"xiaoming", 15, 5, 1, tools}

	toMap := StructToMap(p)
	fmt.Printf("structToMap:%+v\n", toMap)
	// MapToStruct()
	var people People
	err := mapstructure.Decode(toMap, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("people:%+v\n", people)
	fmt.Println("=================")
	toStruct := MapToStruct(toMap)
	fmt.Printf("people2:%+v\n", toStruct)
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
func MapToStruct(mapInstance map[string]interface{}) (people People) {
	var p People
	err := mapstructure.Decode(mapInstance, &p)
	if err != nil {
		fmt.Println(err)
	}
	return p
}
