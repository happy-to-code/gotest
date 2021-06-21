package main

import (
	"fmt"
	"github.com/pborman/uuid"
)

type ID_Version struct {
	ObjectId string
	Version  int
}

func main() {
	var idVersions []ID_Version

	id_versions := forFun(idVersions)
	fmt.Printf("id_versions:%+v\n", id_versions)

}

func forFun(ivs []ID_Version) []ID_Version {
	for i := 0; i < 5; i++ {
		ivs = append(ivs, ID_Version{ObjectId: uuid.NewUUID().String(), Version: i})
	}
	return ivs
}
