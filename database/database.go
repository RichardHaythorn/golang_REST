package database

import "fmt"

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int64  `json:"age"`
}

type Message struct {
	Type   string
	Person []Person
	Err error
}

var IN_channel = make(chan Message)
var OUT_channel = make(chan Message)

var Persons = []Person{
	{ID: 0, FirstName: "John", LastName: "Smith", Age: 25},
	{ID: 1, FirstName: "Steven", LastName: "Gerrard", Age: 46},
}

func GetPersons() ([]Person, error) {
	return Persons, nil
}

func PostPerson(newPerson Person) error {
	Persons = append(Persons, newPerson)
	return nil
}

func Main() {
	fmt.Println("Starting Database")
	for {
		select {
		case in_msg := <-IN_channel:
			fmt.Println("received message", in_msg)
			switch in_msg.Type {
			case "GET":
				persons, err := GetPersons()
				out_msg := Message{Type: "GET", Person: persons, Err: err}
				OUT_channel <- out_msg
			case "POST":
				err := PostPerson(in_msg.Person[0])
				out_msg := Message{Type: "POST", Person: nil, Err: err}
				OUT_channel <- out_msg
			}
		default:
		}
	}

}
