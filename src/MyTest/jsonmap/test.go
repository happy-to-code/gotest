// Test for json and map converting
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{
    "check": {
        "digest": "64313237313338663739393632663038643965356663363639353965366562373133346436323432363030376432626230323866383830373265343731306339",
        "hashAlgo": "sha256",
        "sign": "3045022100dabee3a740d23f9ab4672d272bde2795894da1cd3c054856865b4add0d1f1ac10220391171c73e4a6b6872ce671fcbb3dcf4ed29bb87d774625ba1217717d86b1a05"
    },
    "data": {
        "address": "13821324323",
        "signoffTime": "2020-06-11 10:32:19",
        "method": "短信",
        "sender": "相城法院",
        "idcard": "3456784567845678",
        "people": "张三",
        "sendTime": "2020-06-10 15:32:23"
    },
    "header": {
        "bizSystemId": "通达海-电子送达系统",
        "caseId": "XC1001-001",
        "category": "电子送达",
        "courtId": 123,
        "courtName": "相城法院",
        "subCategory": "送达签收",
        "timestamp": 1593772650696
    }
}`

	// test json string to map
	m, err := JsonToMap(jsonStr)
	if err != nil {
		fmt.Printf("Convert json to map failed with error: %+v\n", err)
	}

	fmt.Printf("Converted to map result: %+v\n", m)

	// test map to json string
	jsonRes, err := MapToJson(m)
	if err != nil {
		fmt.Printf("Convert json to map failed with error: %+v\n", err)
	}

	fmt.Printf("Convert to json string result: %+v\n", jsonRes)

}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}
