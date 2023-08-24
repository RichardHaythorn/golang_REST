package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RichardHaythorn/golang_REST/database"
)

var d = database.PersonsDB{
	Data: []database.Person{
	{ID: 0, FirstName: "John", LastName: "Smith", Age: 25},
	{ID: 1, FirstName: "Steven", LastName: "Gerrard", Age: 46},
},
}

func GetPersons(c *gin.Context) {
	persons := d.GetEntity(-1)
	c.JSON(http.StatusOK, persons)
}

// func GetPersonByFirstName(c *gin.Context, d database.PersonsDB) {
// 	id_str := c.Param("id")
// 	ID, err := strconv.ParseInt(id_str,0,64)
// 	if err != nil{}

// 	persons := d.GetEntity(strconv.ParseInt(ID))
// 	c.JSON(http.StatusOK, persons)

// 	c.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
// }

// func PostPerson(c *gin.Context) {
// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	msg_str := fmt.Sprintf("{\"type\": \"%s\", \"Person\" : [%s]}", c.Request.Method, string(body))
// 	msg_bytes := []byte(msg_str)
// 	database.IN_channel <- msg_bytes
// 	return_msg_raw := <-database.OUT_channel
// 	c.String(http.StatusOK, string(return_msg_raw))
// }

// func PatchPerson(c *gin.Context) {

// 	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 	var updatedPerson database.Person

// 	if err := c.BindJSON(&updatedPerson); err != nil {
// 		return
// 	}
// 	if updatedPerson.ID != 0 {
// 		c.JSON(http.StatusNotAcceptable, "Cannot update ID")
// 	}

// 	newpersons := []database.Person{}
// 	for _, person := range database.Persons {
// 		if person.ID == id {
// 			newpersons = append(newpersons, updatedPerson)
// 		} else {
// 			newpersons = append(newpersons, person)
// 		}
// 	}
// 	database.Persons = newpersons
// 	c.JSON(http.StatusCreated, updatedPerson)
// }

func Main() {
	fmt.Println("Starting Server")

	router := gin.Default()
	router.GET("/persons", GetPersons)
	// router.POST("/persons", PostPerson)
	// router.GET("/persons/:firstname", GetPersonByFirstName)
	// router.PATCH("/persons/:id", PatchPerson)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server Running")
}
