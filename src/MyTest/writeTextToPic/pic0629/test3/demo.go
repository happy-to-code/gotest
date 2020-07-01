//图片合成
package main

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func main() {
	HandleUserImage("pic1")

}

func HandleUserImage(fileName string) (string, error) {
	m, err := imaging.Open("test3.png")
	if err != nil {
		fmt.Printf("open file failed")
	}

	bm, err := imaging.Open("target.png")
	if err != nil {
		fmt.Printf("open file failed")
	}

	// 图片按比例缩放
	dst := imaging.Resize(m, 300, 500, imaging.Lanczos)
	// 将图片粘贴到背景图的固定位置
	result := imaging.Overlay(bm, dst, image.Pt(120, 140), 1)

	fileName = fmt.Sprintf("%d.jpg", fileName)
	err = imaging.Save(result, fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
