package database

type Person struct {
    ID        int64    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int64    `json:"age"`
}

var Persons = []Person{
	{ID: 0, FirstName: "John", LastName: "Smith", Age: 25},
	{ID: 1, FirstName: "Steven", LastName: "Gerrard", Age: 46},
}

func GetDBPersons() []Person {
	return Persons
}
