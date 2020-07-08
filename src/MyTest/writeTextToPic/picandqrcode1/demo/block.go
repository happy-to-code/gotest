package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"io"
	"os"
)

type PicTemplate struct {
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

type CertTemplate struct {
	ImagePath       string
	DefaultFontPath string
	FontDPI         int
	Blocks          []struct {
		BlockType BlockType `json:"type"`
		BlockObj  json.RawMessage
	}
}

func UnmarshalBlocks(tmpl CertTemplate) []Block {
	blocks := make([]Block, 0, len(tmpl.Blocks))
	for _, blk := range tmpl.Blocks {
		switch blk.BlockType {
		case BlockType_Text:
			var text TextBlock
			json.Unmarshal(blk.BlockObj, &text)
			blocks = append(blocks, text)
		case BlockType_DynamicText:
			var text DynamicTextBlock
			json.Unmarshal(blk.BlockObj, &text)
			blocks = append(blocks, text)
		}
	}
	return blocks
}

type BlockType string

const (
	BlockType_Text        BlockType = "text"
	BlockType_DynamicText BlockType = "dynamic"
)

type Block interface {
	Bounds() (x, y float64)
	Draw(dc *gg.Context, params map[string]interface{})
	// Type() BlockType
}

type TextBlock struct {
	Color    [4]uint8
	FontFace string
	FontSize float64
	Point    [2]float64
	Content  string
}

func (tb TextBlock) Bounds() (x, y float64) {
	return tb.Point[0], tb.Point[1]
}

func (tb TextBlock) Draw(dc *gg.Context, _ map[string]interface{}) {
	dc.LoadFontFace(tb.FontFace, tb.FontSize)
	dc.SetColor(color.RGBA{R: tb.Color[0], G: tb.Color[1], B: tb.Color[2], A: tb.Color[3]})
	x, y := tb.Bounds()
	dc.DrawStringAnchored(tb.Content, x, y, 0, 0)
}

type DynamicTextBlock struct {
	TextBlock
	ContentList []struct {
		ShowValue, Field string
	}
}

func (dtb DynamicTextBlock) Draw(dc *gg.Context, params map[string]interface{}) {
	dc.LoadFontFace(dtb.FontFace, dtb.FontSize)
	dc.SetColor(color.RGBA{dtb.Color[0], dtb.Color[1], dtb.Color[2], dtb.Color[3]})
	x, y := dtb.Bounds()
	fmt.Println(x, y)
	for _, arg := range dtb.ContentList {
		value, has := params[arg.Field]
		fmt.Println(value)
		if !has {
			continue
		}
		dc.DrawStringAnchored(arg.ShowValue, x, y, 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("%v", value), x, y, 0, 0)
	}
}

type QrcodeBlock struct {
}

func GenerateCert(dc *gg.Context, blocks []Block, params map[string]interface{}) {
	for _, blk := range blocks {
		blk.Draw(dc, params)
	}
}

//	读取文件中的数据
func ReadJson(filePath string) (result string) {
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

func GetPicTemplate(templateName string) (template PicTemplate) {
	// var filePath = GetCurrentPath() + "/" + templateName
	var filePath = "D:/200707Go/gotest/src/MyTest/writeTextToPic/picandqrcode1/conf/" + templateName

	var data PicTemplate
	result := ReadJson(filePath)
	err := json.Unmarshal([]byte(result), &data)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	return data
}

func main1() {
	template := GetPicTemplate("delivery.json")
	fmt.Println("template::::", template)

}
