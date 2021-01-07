package main

import "fmt"

func main() {
	var refKeyMap map[string]bool = map[string]bool{
		"register_subject_account_ref": true,
		"register_transaction_ref":     false,
	}

	// var refKeyMap map[string]bool
	fmt.Println("refKeyMap1:", refKeyMap)

	delete(refKeyMap, "register_subject_account_ref")
	fmt.Println("refKeyMap2:", refKeyMap)

	// isCheckHash := refKeyMap["register_subject_account_ref"]
	// fmt.Println("isCheckHash:", !isCheckHash)

}
