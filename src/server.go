package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

type person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

var persons = []person{
	{FirstName: "John", LastName: "Smith", Age: 25},
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

func main() {

    router := gin.Default()
    router.GET("/persons", getPersons)
    router.POST("/persons", postPersons)

    router.Run("localhost:8080")
}
