package main

import (
	"database/sql"
	_ "dm"
	"fmt"
)

var db *sql.DB

func main() {
	x, err := sql.Open("dm", "dm://SYSDBA:SYSDBA@10.1.3.150:5236?autoCommit=true")
	db = x
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
	query()
	//update()
}

func insert() {
	sql := fmt.Sprintf("insert into bc_tt values(?,?,?)")
	res, err := db.Exec(sql, 1, "2", 3)
	fmt.Println(res, err)
}

func query() {
	sql := fmt.Sprintf("select state from bc_tt where id = ? and txid = ? and STATE =?")
	var state string
	res := db.QueryRow(sql, 1,234,7).Scan(&state)
	if res != nil {
		fmt.Println(res.Error())
	}
	fmt.Println(state)
	//for rows.Next() {
	//	if err = rows.Scan(&state); err != nil {
	//		fmt.Println(err)
	//	}
	//	// blobLen, _ := photo.GetLength()
	//	fmt.Printf("%v\t %v\n", state,state)
	//}
	//fmt.Println(err)
	//fmt.Println("state",state)
}

func update() {
	sql := fmt.Sprintf("update  bc_tt set STATE=?,TXID = ? where ID = ? ")
	res, err := db.Exec(sql, 7,"234")
	row, err := res.RowsAffected()
	fmt.Println(err, row)
}
