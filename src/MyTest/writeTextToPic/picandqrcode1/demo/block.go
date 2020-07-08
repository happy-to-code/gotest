package main

import (
	"encoding/json"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
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
	FontDPI         float64
	Blocks          []struct {
		BlockType BlockType
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
	dc.DrawStringAnchored(tb.Content, x, y, 0.5, 0.5)
}

type DynamicTextBlock struct {
	TextBlock
	ContentList []struct {
		ShowValue, Field string
	}
}

func (dtb DynamicTextBlock) Draw(dc *gg.Context, params map[string]interface{}) {
	dc.LoadFontFace(dtb.FontFace, dtb.FontSize)
	dc.SetColor(color.RGBA{R: dtb.Color[0], G: dtb.Color[1], B: dtb.Color[2], A: dtb.Color[3]})
	x, y := dtb.Bounds()
	fmt.Println(x, y)
	for _, arg := range dtb.ContentList {
		var temp float64 = 0
		value, has := params[arg.Field]
		fmt.Println(value)
		if !has {
			continue
		}
		dc.DrawStringAnchored(arg.ShowValue, x, y+temp, 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("%v", value), x+1000, y+temp, 0, 0)
		temp += 300
	}
}

type QrcodeBlock struct {
}

func GenerateCert(tmplFilePath string, params map[string]interface{} /*, w io.Writer*/) error {
	tmpl, err := readJson(tmplFilePath)
	if err != nil {
		return err
	}
	blocks := UnmarshalBlocks(tmpl)
	fmt.Println("blocks:::::", blocks)
	im, err := gg.LoadImage(tmpl.ImagePath)
	if err != nil {
		return err
	}
	dc := gg.NewContextForImage(im)

	// 获取模板图片的尺寸
	width := im.Bounds().Size().X
	height := im.Bounds().Size().Y
	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)

	// 加载默认字体
	dc.LoadFontFace(tmpl.DefaultFontPath, tmpl.FontDPI)

	doGenerate(dc, blocks, params)
	// return dc.EncodePNG(w)
	dc.SavePNG("89.png")
	return err
}

func doGenerate(dc *gg.Context, blocks []Block, params map[string]interface{}) {
	for _, blk := range blocks {
		blk.Draw(dc, params)
	}
}

//	读取文件中的数据
func readJson(filePath string) (CertTemplate, error) {
	var tmpl CertTemplate
	file, err := os.Open(filePath)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return tmpl, err
	}
	err = json.NewDecoder(file).Decode(&tmpl)
	if err != nil {
		return tmpl, err
	}
	return tmpl, nil
}

func main() {
	var s = "E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode1\\conf\\delivery1.json"

	var storeInfoStr = `{
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
	paramMap, _ := JsonToMap(storeInfoStr, 1594137931634)

	GenerateCert(s, paramMap)
}
