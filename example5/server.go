package main

import (
	"gorestapi/example5/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	userController := controllers.UserController{}
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/users", userController.Create)
		v1.GET("/users", userController.Read)
		v1.GET("/users/:id", userController.Index)
		v1.PUT("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)
	}
	router.Run(":3000")
}
