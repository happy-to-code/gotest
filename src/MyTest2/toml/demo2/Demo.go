package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type item struct {
	ItemName   string
	ItemDesc   string
	Attributes attributes
	Part       []part
}

type part struct {
	PartName string
	PartLink string
}

type attributes struct {
	Material string
	Light    string
	TransPC  string
	Displace string
}

type items struct {
	Item []item
}

func main() {

	var allitems items

	if _, err := toml.DecodeFile("E:\\20.06.16Project\\GoTest\\src\\MyTest2\\toml\\demo2\\123.toml", &allitems); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n")

	for _, items := range allitems.Item {
		fmt.Printf("    Item Name: %s \n", items.ItemName)
		fmt.Printf("  Description: %s \n\n", items.ItemDesc)

		fmt.Printf("     Material: %s \n", items.Attributes.Material)
		fmt.Printf("     Lightmap: %s \n", items.Attributes.Light)
		fmt.Printf(" TL Precision: %s \n", items.Attributes.TransPC)
		fmt.Printf("   DP Channel: %s \n", items.Attributes.Displace)

		fmt.Printf("    Part Name: %s \n", items.Part[0].PartName)
		fmt.Printf("    Part Link: %s \n", items.Part[0].PartLink)

		fmt.Printf("\n────────────────────────────────────────────────────┤\n\n")

	}

	fmt.Printf("\n")

}
