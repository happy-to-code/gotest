package main

import (
	"log"
	"net/http"
	"os"

	"reflect"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {

	var (
		excel_file_path string                         = "./1.xlsx"
		file_result     map[int]map[int]map[int]string = make(map[int]map[int]map[int]string)
		sheet_result    map[int]map[int]string         = make(map[int]map[int]string)
	)
	// 打开一个excel文件资源
	f, err := xlsx.OpenFile(excel_file_path)
	if err != nil {
		log.Println(err.Error())
	}
	// 循环文件中所有工作表
	for sheet_key, sheet := range f.Sheets {
		// 循环对应工作表中行数
		for key, row := range sheet.Rows {
			row_result := make(map[int]string)
			// 循环工作表行数的每一列
			for k, cell := range row.Cells {
				row_result[k] = cell.Value
			}
			// 如果为空不添加对应值到 数组
			if !Empty(row_result) {
				sheet_result[key] = row_result
			}
		}
		// 如果为空不添加对应值到 数组
		if !Empty(sheet_result) {
			file_result[sheet_key] = sheet_result
		}

	}
	// 输出表格的结果
	for _, sheet := range file_result {
		for k, _ := range sheet {

			log.Printf("%d=%v\n", k, sheet[k])
		}

	}

}
func Empty(params interface{}) bool {
	// 初始化变量
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	// 获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	// 由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

// 统计数组长度
func Count(info []interface{}) int {
	return len(info)
}

// 设置等待时间
func Sleep(s time.Duration) {
	time.Sleep(s * time.Second)
}

// 终止程序
func Die(result interface{}) {
	log.Println(result.(string))
	os.Exit(1)
}

// 终止程序
func Exit(result interface{}) {
	log.Println(result.(string))
	os.Exit(1)

}

// 获取网页数据
func File_get_content(url string) string {
	var result string = ""
	if Empty(url) == true {
		return result
	}
	h, err := http.Get(url) // 获取url资源
	if err != nil {         // 如果获取失败返回空
		return result
	}
	if h.StatusCode != http.StatusOK { // 如果获取状态码不为 200 直接返回空
		return result
	}
	defer h.Body.Close()

	buf := make([]byte, 1024)
	for {
		num, _ := h.Body.Read(buf) // 读取http 内容
		if num == 0 {
			break
		}
		result += string(buf)
	}
	return result
}
