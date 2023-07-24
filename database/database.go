package database

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int64  `json:"age"`
}

type Message struct {
	Type   string   `json:"type"`
	Person []Person `json:"person"`
	Err    error    `json:"err"`
}

var IN_channel = make(chan []byte)
var OUT_channel = make(chan []byte)

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
		in_msg_raw := <-IN_channel
		var in_msg Message
		out_msg_raw := make([]byte,100)
		err := json.Unmarshal(in_msg_raw, &in_msg)
		if err != nil {
			fmt.Println(err)
			err_bytes := []byte("Database error unmarshalling message")
			OUT_channel <- err_bytes
		}
		fmt.Println("received message", in_msg)
		switch in_msg.Type {
		case "GET":
			persons, err := GetPersons()
			if err != nil {
				out_msg_raw, _ := json.Marshal(err)
				OUT_channel <- out_msg_raw
			}
			out_msg_raw, err = json.Marshal(persons)
			if err != nil {
				out_msg_raw, _ := json.Marshal(err)
				OUT_channel <- out_msg_raw
			}
			case "POST":
				err := PostPerson(in_msg.Person[0])
				if err != nil {
					out_msg_raw, _ := json.Marshal(err)
					OUT_channel <- out_msg_raw
				}
				out_msg_raw, err = json.Marshal(in_msg)
				if err != nil {
					out_msg_raw, _ := json.Marshal(err)
					OUT_channel <- out_msg_raw
				}		
		}
		if err == nil {
			OUT_channel <- out_msg_raw
		}
	}

}
