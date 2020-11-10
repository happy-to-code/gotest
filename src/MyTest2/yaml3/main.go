package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MyFile struct {
	Dicts []Dict `yaml:"dicts"`
}
type Dict struct {
	Item  string `yaml:"item"`
	Label Label  `yaml:"label"`
}
type Label struct {
	Name     string `yaml: "name"`
	Code     int32  `yaml: "code"`
	Describe string `yaml: "describe"`
}

var ds map[string][]Label

func main() {

	ds = getProfileMap()
	fmt.Printf("---%v\n", ds)

	fmt.Println("=======", ds["CQ_GFXZ"])
	fmt.Println("=======", getLabel("CQ_GFXZ", 1).Describe)
	// fmt.Println("=======", getLabel("BZ", 20).Describe)

	// dir,_ := os.Getwd()
	// fmt.Println("当前路径：",dir)

}

func getLabel(item string, code int32) *Label {
	for k, v := range ds {
		if k == item {
			for _, label := range v {
				if label.Code == code {
					return &label
				}
			}
		}
	}
	return nil
}

func getProfileMap() map[string][]Label {
	var c MyFile
	// 读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf1()

	var item2Label = make(map[string][]Label)
	for _, d := range conf.Dicts {
		var lab []Label
		lab = append(lab, d.Label)

		var hasExist = false
		for k, v := range item2Label {
			if d.Item == k {
				v = append(v, d.Label)
				item2Label[k] = v
				hasExist = true
			}
		}
		if !hasExist {
			item2Label[d.Item] = lab
		}

	}

	return item2Label
}

func (myFile *MyFile) getConf1() *MyFile {
	// 应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("E:\\20.06.16Project\\GoTest\\src\\MyTest2\\yaml3\\dict.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	// err = yaml.Unmarshal(yamlFile, myFile)
	err = yaml.UnmarshalStrict(yamlFile, myFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	return myFile
}
