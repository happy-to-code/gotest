package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	s := fmt.Sprintf("%s+%s", "kkkk", "LLLLL")
	fmt.Println(s)

	sendContentDetail := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%d", "网络列表中,名称为:",
		"name", ",\nid为:", "peer.Id", ",\n内网地址为:", "peer.InAddr", ",\n外网地址为:", "peer.OutAddr", "的节点\n已经离线,请尽快处理!",1)
	fmt.Println(sendContentDetail)
}
