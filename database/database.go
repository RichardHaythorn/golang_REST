package database

type Database interface {
	GetEntity(ID int64) []Person
	PostEntity() string
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

func (d PersonsDB) GetEntity(ID int64) []Person {
	if ID == -1 {
		return d.Data
	} else {
		for i := 0; i < len(d.Data); i++ {
			if d.Data[i].ID == ID {
				return []Person{d.Data[i]}
			}
		}
	}
	return nil

}

func (d PersonsDB) PostEntity() string {
	return ""
}

// func PostPerson(newPerson Person) error {
// 	Persons = append(Persons, newPerson)
// 	return nil
// }
