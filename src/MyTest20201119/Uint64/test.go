package main

import "fmt"

func main() {
	// var a uint64 = -1
	// fmt.Println(a)

	// msg SQL语句
	var msgSQL = "CREATE TABLE IF NOT EXISTS msg_"
	var msgSQL2 = `
		UID INTEGER PRIMARY KEY AUTOINCREMENT,
		Topic VARCHAR(64) NULL,
		SubTopic VARCHAR(64) NULL,
		Hash VARCHAR(128) NULL,
		Date DATE NULL,
		Height INTEGER NULL,
		EventMsg BLOB
	);
	-- CREATE INDEX IF NOT EXISTS msgIndex ON msg_`
	msgSQL = fmt.Sprintf("%s%s%s", msgSQL, "ledgerId", "(")
	msgSQL = fmt.Sprintf("%s%s%s%s", msgSQL, msgSQL2, "ledgerId", "(Height);")

	fmt.Println(msgSQL)
}
