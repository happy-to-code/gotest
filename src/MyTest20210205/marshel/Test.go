package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var ossDatas []OSSData

	// 元数据
	var metaObject = MetaObject{
		Data_type: "supervision",
	}
	for i := 0; i < 5; i++ {
		var ossDataReturn = OSSData{
			Digest:      fmt.Sprintf("%s+%d", "Digest", i),
			Digest_algo: "sha-256",
			Uri:         fmt.Sprintf("%s+%d", "Uri", i),
			Meta:        metaObject,
		}

		ossDatas = append(ossDatas, ossDataReturn)
	}
	fmt.Printf("%+v\n", ossDatas)
	fmt.Printf("%+v\n", "========================")
	ossDataBytes, _ := json.Marshal(ossDatas)
	fmt.Println("--->", string(ossDataBytes))
	var datas []OSSData
	err := json.Unmarshal(ossDataBytes, &datas)
	if err != nil {
		panic(err)
	}
	fmt.Println("datas:", datas)

	// var tempOssDatas = []OSSData{
	// 	{
	// 		Digest:      fmt.Sprintf("%s+%d", "Digest", 888),
	// 		Digest_algo: "sha-256",
	// 		Uri:         fmt.Sprintf("%s+%d", "Uri", 888),
	// 		Meta:        metaObject,
	// 	},
	// 	{
	// 		Digest:      fmt.Sprintf("%s+%d", "Digest", 999),
	// 		Digest_algo: "sha-256",
	// 		Uri:         fmt.Sprintf("%s+%d", "Uri", 999),
	// 		Meta:        metaObject,
	// 	},
	// }
	// fmt.Printf("%+v\n", tempOssDatas)
	// fmt.Println("-----------------------------")
	// ossDatas = append(ossDatas, tempOssDatas...)
	// fmt.Printf("ossdatas:%+v%+v\n", ossDatas, len(ossDatas))
}

// OSS
type OSSData struct {
	Digest      string      `json:"digest"`
	Digest_algo string      `json:"digest_algo"`
	Uri         string      `json:"uri"`
	Meta        interface{} `json:"meta"`
}

// OSS元数据
type MetaObject struct {
	Data_type string `json:"data_type"`
}
