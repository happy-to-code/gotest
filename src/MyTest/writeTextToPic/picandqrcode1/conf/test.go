package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//	读取文件中的数据
func readJson(filePath string) (result string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}
	return result
}

func main() {
	var data struct {
		ImagePath       string
		DefaultFontPath string
		FontDPI         float64
		Blocks          []struct {
			Color       []uint8
			FontFace    string
			FontSize    float64
			Point       []float64
			BlockType   string
			Content     string
			ContentList []struct {
				Field     string
				ShowValue string
			}
		}
	}
	result := readJson("E:/20.06.16Project/GoTest/src/MyTest/writeTextToPic/picandqrcode1/conf/delivery.json")
	err := json.Unmarshal([]byte(result), &data)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("data:", data)
	fmt.Println("===:", data.Blocks[0].FontSize)
}
