package main

import "fmt"

func main() {
	var temp = "aaa"

	var sqlStart = "CREATE TABLE IF NOT EXISTS subcli_"
	sprintf := fmt.Sprintf("%s%s%s", sqlStart, temp, "(")
	fmt.Println("sprintf====>:", sprintf)
	fmt.Println("--------------")
	fmt.Printf("%T\n", sprintf)

	sqlTable := `
		ID VARCHAR(256) PRIMARY KEY,
		Date DATE NULL,
		Height INTEGER DEFAULT 0
	);
	`
	fmt.Println(fmt.Sprintf("%s%s", sprintf, sqlTable))
	fmt.Println("=======================================\n")
	s := appendSQL("xcvId")
	fmt.Println(s)

}

func appendSQL(ledgerId string) string {
	// subcli SQL语句
	var subcliSQL = "CREATE TABLE IF NOT EXISTS subcli_"
	var subcliSQL2 = `
		ID VARCHAR(256) PRIMARY KEY,
		Date DATE NULL,
		Height INTEGER DEFAULT 0
	);
	`
	subcliSQL = fmt.Sprintf("%s%s%s", subcliSQL, ledgerId, "(")
	subcliSQL = fmt.Sprintf("%s%s", subcliSQL, subcliSQL2)

	// msg SQL语句
	var msgSQL = "CREATE TABLE IF NOT EXISTS msg_"
	var msgSQL2 = `
		UID INTEGER PRIMARY KEY AUTOINCREMENT,
		Topic VARCHAR(64) NULL,
		SubTopic VARCHAR(64) NULL,
		Hash BINARY UNIQUE,
		Date DATE NULL,
		Height INTEGER NULL,
		EventMsg BLOB
	);
	-- CREATE INDEX IF NOT EXISTS msgIndex ON msg(Height);
	`
	msgSQL = fmt.Sprintf("%s%s%s", msgSQL, ledgerId, "(")
	msgSQL = fmt.Sprintf("%s%s", msgSQL, msgSQL2)

	return fmt.Sprintf("%s%s", subcliSQL, msgSQL)
}
