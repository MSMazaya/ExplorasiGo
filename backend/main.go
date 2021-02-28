package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	fmt.Println("Connecting Go to MySQL")
	db, err := sql.Open("mysql", "root:accessdenied@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('ELIOT')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Success")

}
