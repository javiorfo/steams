package data

type Person struct {
	Name string
	Age  int
	Pets []Pet
}

type Pet struct {
	Name string
	Type string
	Age  int
}

const DOG = "DOG"
const CAT = "CAT"

var PeopleWithPets = []Person{
	{Name: "Carl", Age: 34, Pets: []Pet{}},
	{Name: "John", Age: 20, Pets: []Pet{{Name: "Bobby", Type: DOG, Age: 2}, {Name: "Mike", Type: DOG, Age: 12}}},
	{Name: "Grace", Age: 40, Pets: []Pet{{Name: "Pepe", Type: DOG, Age: 4}, {Name: "Snowball", Type: CAT, Age: 8}}},
	{Name: "Robert", Age: 40, Pets: []Pet{{Name: "Ronny", Type: CAT, Age: 3}}},
}
