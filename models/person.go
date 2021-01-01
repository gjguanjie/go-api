package models

var (
	personList map[string]*Person
)

func init()  {
	personList = make(map[string]*Person)
	var person = Person{"lishi",100}
	personList["1"] = &person
}

type Person struct {
	Name string `json:name`
	Age int `json:age`
}

func GetAllPerson()map[string]*Person  {
	return personList
}
