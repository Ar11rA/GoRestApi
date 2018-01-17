package database

import (
	"github.com/jinzhu/gorm"
	// mysql connection library
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect is for establishing mysql connectivity
func Connect() (*gorm.DB, error) {
	//open a db connection
	db, err := gorm.Open("mysql", "root:pass@/gotest_orm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
