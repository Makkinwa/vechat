package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)
var id = 1
func mysql_handle(buf []byte){
	db, err := sql.Open("mysql", "admin:pass@tcp(127.0.0.1:3306)/data")
	if err != nil {
		print("MySQL connection error")
		os.Exit(2)
	} else {
		db.Exec("CREATE TABLE IF NOT EXISTS admin (id INT, message varchar(255));")
		var sql = "INSERT INTO admin VALUES (" + strconv.Itoa(id) + ", '" + string(buf) +"')"
		db.Exec(sql)
		id = id + 1
	}

}
