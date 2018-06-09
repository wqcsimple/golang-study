package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:7@/yeneiren")

	defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}

	rows, err := db.Query("SELECT name FROM `user` WHERE weight >= 0")
	if err != nil {
		panic(err)
		return
	}

	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
			return
		}

		fmt.Println(name)
	}
	

}