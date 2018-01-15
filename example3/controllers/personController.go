package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"gorestapi/example3/database"
	"gorestapi/example3/models"

	"github.com/gin-gonic/gin"
)

// PersonController for person
type PersonController struct{}

// GetAll person detail
func (P *PersonController) GetAll(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	var (
		person  models.Person
		persons []models.Person
	)
	rows, err := db.Query("select id, first_name, last_name from person;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		persons = append(persons, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": persons,
		"count":  len(persons),
	})
	defer db.Close()
}

// Get single person detail
func (P *PersonController) Get(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	row := db.QueryRow("select id, firstname, lastname from person where id = ?;", id)
	err = row.Scan(&person.ID, &person.FirstName, &person.LastName)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
	defer db.Close()
}

// Create person
func (P *PersonController) Create(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	var buffer bytes.Buffer
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	stmt, err := db.Prepare("insert into person (first_name, last_name) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(firstname, lastname)

	if err != nil {
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(firstname)
	buffer.WriteString(" ")
	buffer.WriteString(lastname)
	defer stmt.Close()
	name := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name),
	})
	defer db.Close()
}

// Update person
func (P *PersonController) Update(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	var buffer bytes.Buffer
	id := c.Query("id")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	stmt, err := db.Prepare("update person set first_name= ?, last_name= ? where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(firstname, lastname, id)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(firstname)
	buffer.WriteString(" ")
	buffer.WriteString(lastname)
	defer stmt.Close()
	name := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated to %s", name),
	})
	defer db.Close()
}

// Delete Person
func (P *PersonController) Delete(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	id := c.Query("id")
	stmt, err := db.Prepare("delete from person where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})
	defer db.Close()
}
