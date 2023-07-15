package api_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/RichardHaythorn/golang_REST/api"
	"github.com/gin-gonic/gin"
)

func TestPOST(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/persons")
	if err != nil {
		t.Errorf("Failed GET")
	}
	if resp.Status != "200 OK" {
		t.Errorf("Resp status not OK")
	}
	defer resp.Body.Close()
}

func TestGET(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/persons")
	if err != nil {
		t.Errorf("Failed GET")
	}
	if resp.Status != "200 OK" {
		t.Errorf("Resp status not OK, Status: %s", resp.Status)
	}
	defer resp.Body.Close()
}

func TestGETByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/persons/John")
	if err != nil {
		t.Errorf("Failed GET")
	}
	if resp.Status != "200 OK" {
		t.Errorf("Resp status not OK, Status: %s", resp.Status)
	}
	defer resp.Body.Close()
}

func TestPATCH(t *testing.T) {

}

func setup(router *gin.Engine) {
	router.GET("/persons", api.GetPersons)
	router.POST("/persons", api.PostPerson)
	router.GET("/persons/:firstname", api.GetPersonByFirstName)
	router.PATCH("/persons/:id", api.PatchPerson)

	router.Run("localhost:8080")
}

func TestMain(m *testing.M) {
	router := gin.Default()
	go setup(router)
	code := m.Run()
	os.Exit(code)
}
