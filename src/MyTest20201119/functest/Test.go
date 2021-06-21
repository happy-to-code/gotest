package main

import (
	"fmt"
	"time"
)

func main() {
	paramerBody := getParamerBody("a", "t", "12", "3", "ssss", "23423")
	fmt.Println(paramerBody)
	fmt.Println(time.Now().Unix())
}

var body = `
	{
		"appid":"%s",
		"token":"%s",
		"hash":"%s",
		"projectName":"%s",
		"storeType":"%s",
		"picUrl":"%s"
	}
	`

func getParamerBody(appid, token, hash, projectName, storeType, picUrl string) string {
	vbody := fmt.Sprintf(body, appid, token, hash, projectName, storeType, picUrl)
	return vbody
}
