// be-module/database.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

const (
	// OracleDBConnectionString is the connection string for Oracle DB
	OracleDBConnectionString = "your_oracle_db_connection_string"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("godror", OracleDBConnectionString)
	if err != nil {
		fmt.Println("Error connecting to Oracle DB:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging Oracle DB:", err)
		return
	}

	fmt.Println("Connected to Oracle DB")
}

func closeDB() {
	db.Close()
}

func main() {
	initDB()
	defer closeDB()
}