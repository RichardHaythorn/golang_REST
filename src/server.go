package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type person struct {
    ID        int64    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int64    `json:"age"`
}

var persons = []person{
	{ID: 0, FirstName: "John", LastName: "Smith", Age: 25},
	{ID: 1, FirstName: "Steven", LastName: "Gerrard", Age: 46},
}

type PersonServer struct {
	Server http.Server
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func postPersons(c *gin.Context) {
    var newPerson person

    if err := c.BindJSON(&newPerson); err != nil {
        return
    }

    persons = append(persons, newPerson)
    c.IndentedJSON(http.StatusCreated, newPerson)
}

func patchPerson(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var updatedPerson person

    if err := c.BindJSON(&updatedPerson); err != nil {
        return
    }

    newpersons := []person{}
    for _, person := range persons {
        if person.ID == id {
            newpersons = append(newpersons, updatedPerson)
        } else {            
            newpersons = append(newpersons, person)
        }
    }
    persons = newpersons
    c.IndentedJSON(http.StatusCreated, updatedPerson)
}

func getPersonByFirstName(c *gin.Context) {
    firstname := c.Param("firstname")

    for _, person := range persons {
        if person.FirstName == firstname {
            c.IndentedJSON(http.StatusOK, person)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func main() {

    router := gin.Default()
    router.GET("/persons", getPersons)
    router.POST("/persons", postPersons)
    router.GET("/persons/:firstname", getPersonByFirstName)
    router.PATCH("/persons", patchPerson)

    router.Run("localhost:8080")
}
