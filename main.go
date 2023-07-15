package main

import (
	"github.com/RichardHaythorn/golang_REST/api"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/persons", api.GetPersons)
	router.POST("/persons", api.PostPerson)
	router.GET("/persons/:firstname", api.GetPersonByFirstName)
	router.PATCH("/persons/:id", api.PatchPerson)

	router.Run("localhost:8080")
}
