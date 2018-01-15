package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gorestapi/example4/controllers"
)

func main() {
	router := gin.Default()
	// define a variable for person controller
	personController := controllers.PersonController{}

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/person/:id", personController.Get)
		v1.GET("/persons", personController.GetAll)
		v1.POST("/person", personController.Create)
		v1.PUT("/person", personController.Update)
		v1.DELETE("/person", personController.Delete)
	}
	// run on port 3000
	router.Run(":3000")
}
