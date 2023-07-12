package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPOST(t *testing.T) {
	url := "http://localhost:8080/persons"
	payload := person{ID: 2, FirstName: "Steve", LastName: "Smith", Age: 35}
	// encode payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Errorf("Failed POST")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("Resp status not 201, Status: %s", resp.Status)
	}

	defer resp.Body.Close()
}

func TestGET(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/persons")
	if err != nil {
		t.Errorf("Failed GET")
	}
	if resp.StatusCode != 200 {
		t.Errorf("Resp status not OK, Status: %s", resp.Status)
	}
	defer resp.Body.Close()
}

func TestGETByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/persons/John")
	if err != nil {
		t.Errorf("Failed GET")
	}
	if resp.StatusCode != 200 {
		t.Errorf("Resp status not OK, Status: %s", resp.Status)
	}
	defer resp.Body.Close()
}

func TestPATCH(t *testing.T) {
	url := "http://localhost:8080/persons/0"
	payload := map[string]string{"firstname": "Steve", "lastname": "Jones"}

	// encode payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Errorf("Failed PATCH")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("Resp status not 201 Created, Status: %s", resp.Status)
	}

	defer resp.Body.Close()
}

func setup(router *gin.Engine) {
	router.GET("/persons", getPersons)
	router.POST("/persons", postPersons)
	router.GET("/persons/:firstname", getPersonByFirstName)
	router.PATCH("/persons/:id", patchPerson)

	err := router.Run("localhost:8080")
    if err != nil {
        panic(err)
    }
}

func TestMain(m *testing.M) {
	router := gin.Default()
	go setup(router)
	code := m.Run()
	os.Exit(code)
}
