package controllers

import (
	"net/http"

	"gorestapi/example5/database"
	"gorestapi/example5/models"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct{}

// Create add a new user
func (U *UserController) Create(c *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "message": "Db connect failed"})
		return
	}
	user := models.User{Name: c.PostForm("name"), Email: c.PostForm("email")}
	db.Create(&user)
	if db.NewRecord(user) {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Db conflict"})
		return
	}
	db.Save(user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "user item created successfully!", "resourceId": user.ID})
}

// Read fetch all users
func (U *UserController) Read(c *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "message": "Db connect failed"})
		return
	}
	var users []models.User
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// Index fetch a single user
func (U *UserController) Index(c *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "message": "Db connect failed"})
		return
	}
	var user models.User
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// Update update a user
func (U *UserController) Update(c *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "message": "Db connect failed"})
		return
	}
	var user models.User
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Model(&user).Update("email", c.PostForm("email"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "user updated successfully!"})
}

// Delete remove a user
func (U *UserController) Delete(c *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": http.StatusServiceUnavailable, "message": "Db connect failed"})
		return
	}
	var user models.User
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "user deleted successfully!"})
}
