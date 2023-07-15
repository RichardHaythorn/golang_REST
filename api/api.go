package api

import (
	"net/http"
	"strconv"

	"github.com/RichardHaythorn/golang_REST/database"
	"github.com/gin-gonic/gin"
	//"github.com/apache/arrow/go/v12/arrow"
)

func GetPersons(c *gin.Context) {
	persons, err := database.GetPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, persons)
	} else {
		c.JSON(http.StatusOK, persons)
	}
}

func PostPerson(c *gin.Context) {
	var newPerson database.Person

	//TODO generate new IDs
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	err := database.PostPerson(newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newPerson)
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
