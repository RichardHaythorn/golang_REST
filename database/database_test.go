package database_test

import (
	"github.com/RichardHaythorn/golang_REST/database"
	"testing"
)

func TestGetAllEntities(t *testing.T) {
	_ = database.Person_DB.GetAllEntities()
}

func TestGetEntity(t *testing.T) {
	_, err := database.Person_DB.GetEntity(7)
	if err != nil {
		t.Errorf("Could not get entity")
	}
}

func TestPostEntity(t *testing.T) {
	newPerson := database.Person{ID: 3, FirstName: "Geoff", LastName: "Hurst", Age: 35}
	err := database.Person_DB.PostEntity(newPerson)
	if err != nil {
		t.Errorf("Could not post entity")
	}
}

