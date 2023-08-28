package database

import (
	"errors"
)

type Database interface {
	GetEntity(ID int64) []Person
	GetAllEntities()
	PostEntity(Person) error
}

type PersonsDB struct {
	Data []Person
}

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int64  `json:"age"`
}

var Person_DB = PersonsDB{
	Data: []Person{
		{ID: 0, FirstName: "John", LastName: "Smith", Age: 25},
		{ID: 1, FirstName: "Steven", LastName: "Gerrard", Age: 46},
	},
}

func (d *PersonsDB) GetAllEntities() []Person {
	return d.Data
}

func (d *PersonsDB) GetEntity(ID int64) (Person, error) {
	for i := 0; i < len(d.Data); i++ {
		if d.Data[i].ID == ID {
			return d.Data[i], nil
		}
	}
	return Person{}, errors.New("ID not found")
}

func (d *PersonsDB) PostEntity(p Person) error {
	d.Data = append(d.Data,p)
	return nil
}

