package main

import (

	// mysql connection library
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gorestapi/example5/database"
	"gorestapi/example5/models"
)

// Connect is for establishing mysql connectivity
func migrate() {
	//open a db connection
	db, _ := database.Connect()
	//Migrate the schema
	db.AutoMigrate(&models.User{})
}
