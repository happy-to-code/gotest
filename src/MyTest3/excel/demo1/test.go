package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	inFile  = "E:\\20.06.16Project\\GoTest\\src\\MyTest3\\excel\\demo\\1.xlsx"
	inFile2 = "C:\\Users\\user\\Desktop\\上海股权托管交易中心区块链系统\\excel上传主体信息.xlsx"
)
var ztzz = make(map[int]string)

func main() {
	Import()
	fmt.Println("===================")
	fmt.Println(ztzz)
}

func Import() {
	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)

		var ztzzmap = make(map[string]interface{})
		// 一个sheet 遍历行读取
		for i, row := range sheet.Rows {
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				if i == 0 {
					ztzz[i] = cell.String()
				} else {
					ztzzmap[ztzz[i]] = cell.String()
				}

				// text := cell.String()
				// fmt.Printf("%20s", text)
			}
			fmt.Print("\n")
		}
	}
	fmt.Println("\n\nimport success")
}
