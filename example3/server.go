package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gorestapi/example3/controllers"
)

func main() {
	router := gin.Default()
	// define a variable for person controller
	personController := controllers.PersonController{}

	// GET a person detail
	router.GET("/person/:id", personController.Get)

	// GET all persons
	router.GET("/persons", personController.GetAll)

	// POST new person details
	router.POST("/person", personController.Create)

	// PUT - update a person details
	router.PUT("/person", personController.Update)

	// Delete resources
	router.DELETE("/person", personController.Delete)

	// run on port 3000
	router.Run(":3000")
}
