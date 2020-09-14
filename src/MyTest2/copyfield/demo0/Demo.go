package main

import (
	"errors"
	"fmt"
	"reflect"
)

// 企业信息结构体
type EnterpriseInfo struct {
	EquityCode         string `binding:"required"` // 股权代码
	EquityAbbreviation string `binding:"required"` // 股权简称
	Name               string `binding:"required"`
}

// 企业信息和合约地址结构体
type EnterpriseInfoAndContractAddress struct {
	ContractAddress    string `binding:"required"` // 合约地址
	EquityCode         string `binding:"required"` // 股权代码
	EquityAbbreviation string `binding:"required"` // 股权简称
}

type SharesLock struct {
	Amount        uint64 `binding:"gt=0"`  // 冻结的数量
	ShareProperty uint16 `binding:"gte=0"` // 股份类型
	Name          map[string]interface{}
}
type SharesLockForContract struct {
	Amount        uint64 // 冻结的数量
	ShareProperty uint16 // 股份类型
	Name          string
}

func main() {
	var m = make(map[string]interface{})
	m["12112"] = "BBBB"
	m["12114"] = 888
	sharesLock := SharesLock{
		Amount:        15,
		ShareProperty: 16,
		Name:          m,
	}
	fmt.Printf("sharesLock:%+v\n", sharesLock)

	var sharesLockForContract SharesLockForContract
	err := SimpleCopyProperties(&sharesLockForContract, sharesLock)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("sharesLockForContract:%+v\n", sharesLockForContract)

}

func SimpleCopyProperties(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}
