package main

import (
	
	"github.com/gin-gonic/gin"    	
	"github.com/RichardHaythorn/golang_REST/api"
)

func main() {

    router := gin.Default()
    router.GET("/persons", api.GetPersonsAPI)
    router.POST("/persons", api.PostPersons)
    router.GET("/persons/:firstname", api.GetPersonByFirstName)
    router.PATCH("/persons/:id", api.PatchPerson)

    router.Run("localhost:8080")
}
