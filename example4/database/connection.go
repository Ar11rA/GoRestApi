package database

import (
	"database/sql"
	"fmt"
	// mysql connection driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect for db
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return db, nil
}
