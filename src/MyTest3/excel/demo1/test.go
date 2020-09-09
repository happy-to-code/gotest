package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

var (
	inFile = "E:\\20.06.16Project\\GoTest\\src\\MyTest3\\excel\\demo\\1.xlsx"
)

func main() {
	Import()
}

func Import() {
	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		// 遍历行读取
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			// 遍历每行的列读取
			for _, cell := range row.Cells {

				text := cell.String()
				fmt.Printf("%20s", text)
			}
			fmt.Print("\n")
		}
	}
	fmt.Println("\n\nimport success")
}
