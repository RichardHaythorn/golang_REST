package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RichardHaythorn/golang_REST/database"
	"github.com/gin-gonic/gin"
	//"github.com/apache/arrow/go/v12/arrow"
)

func GetPersons(c *gin.Context) {
	msg := database.Message{Type: "GET", Person: nil}
	database.IN_channel <- msg
	return_msg := <- database.OUT_channel
	if return_msg.Err != nil{
		c.JSON(http.StatusInternalServerError, msg)
	}
	c.JSON(http.StatusOK, return_msg.Person)
}

func PostPerson(c *gin.Context) {
	var newPerson database.Person

	//TODO generate new IDs
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}	
	msg := database.Message{Type: "POST", Person: []database.Person{newPerson}}
	database.IN_channel <- msg
	return_msg := <- database.OUT_channel
	if return_msg.Err != nil{
		c.JSON(http.StatusInternalServerError, msg)
	}
	c.JSON(http.StatusCreated, newPerson)
}

func PatchPerson(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var updatedPerson database.Person

	if err := c.BindJSON(&updatedPerson); err != nil {
		return
	}
	if updatedPerson.ID != 0 {
		c.JSON(http.StatusNotAcceptable, "Cannot update ID")
	}

	newpersons := []database.Person{}
	for _, person := range database.Persons {
		if person.ID == id {
			newpersons = append(newpersons, updatedPerson)
		} else {
			newpersons = append(newpersons, person)
		}
	}
	database.Persons = newpersons
	c.JSON(http.StatusCreated, updatedPerson)
}

func GetPersonByFirstName(c *gin.Context) {
	firstname := c.Param("firstname")

	for _, person := range database.Persons {
		if person.FirstName == firstname {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func Main() {
	fmt.Println("Starting Server")

	router := gin.Default()
	router.GET("/persons", GetPersons)
	router.POST("/persons", PostPerson)
	router.GET("/persons/:firstname", GetPersonByFirstName)
	router.PATCH("/persons/:id", PatchPerson)

	err := router.Run("localhost:8080")
	if err != nil{
		panic(err)
	}

	fmt.Println("Server Running")
}
