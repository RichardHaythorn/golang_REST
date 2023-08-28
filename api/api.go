package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/RichardHaythorn/golang_REST/database"
)



func GetPersons(c *gin.Context) {
	persons := database.Person_DB.GetAllEntities()
	c.JSON(http.StatusOK, persons)
}

func GetPersonByID(c *gin.Context) {
 	ID, err := strconv.ParseInt(c.Param("id"),0,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "Bad ID value")
		return
	}
	persons, err := database.Person_DB.GetEntity(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, "ID not found")
		return
	}
	c.JSON(http.StatusOK, persons)
}

func PostPerson(c *gin.Context) {
    var newPerson database.Person

    if err := c.BindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, "Bad ID")
        return
    }
	database.Person_DB.PostEntity(newPerson)
    c.JSON(http.StatusCreated, newPerson)
}

func Main() {
	fmt.Println("Starting Server")

	router := gin.Default()
	router.GET("/persons", GetPersons)
	router.GET("/persons/:id", GetPersonByID)
	router.POST("/persons", PostPerson)
	// router.PATCH("/persons/:id", PatchPerson)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Running")
}
