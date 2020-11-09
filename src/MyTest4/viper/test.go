package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	jsonStr := `{
			"check": {
				"digest": "39663063336134306338333237373065303036663633623337333132323836303066613630336132383631633438313461343237353631623832653666373761",
				"hashAlgo": "sha256",
				"sign": "3045022100dc4314ac926c4ba578a78cbd50fcb67153a2687e92b6537187e2505da66028a4022050570d4a19d40f54826e6c32dd29795e998c69cca7202b4cb09665ef9cfe5235"
			},
			"data": {
				"address": "13821324323",
				"signoffTime": "2020年07月11日 10:32:19",
				"method": "短信",
				"sender": "相城法院",
				"idcard": "3456784567845678",
				"people": "张三",
				"sendTime": "2020-07-10 15:32:23"
			},
			"header": {
				"bizSystemId": "通达海-电子送达系统",
				"caseId": "XC1001-001",
				"category": "电子送达",
				"courtId": 123,
				"courtName": "相城法院",
				"subCategory": "送达签收",
				"timestamp": 1594006098964
			}
		}`
	v := viperDemo(jsonStr)
	fmt.Println(v.Get("check.digest"))
	fmt.Println(v.Get("check.hashAlgo"))
	fmt.Println(v.Get("picUrl"))
	fmt.Println(v.Get("header.bizSystemId"))
	fmt.Println(v.Get("header.subCategory"))
}

func viperDemo(jsonStr string) *viper.Viper {

	vi := viper.New()
	vi.Set("picUrl", "123/345.jpg")
	vi.SetConfigType("JSON")
	if err := vi.ReadConfig(strings.NewReader(jsonStr)); err != nil {
		panic(err)
	}

	return vi

}
