package database_test

import (
	"github.com/RichardHaythorn/golang_REST/database"
	"testing"
)

func TestGetPersons(t *testing.T) {
	_, err := database.GetPersons()
	if err != nil {
		t.Errorf("Failed Get")
	}
}
