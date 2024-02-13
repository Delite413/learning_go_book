package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	withoutNames := Employee{
		"Edgar",
		"Pascua",
		0,
	}

	withNames := Employee{
		firstName: "Edgar",
		lastName:  "Pascua",
		id:        1,
	}

	var withVarDelcaration Employee
	withVarDelcaration.firstName = "Edgar"
	withVarDelcaration.lastName = "Pascua"
	withVarDelcaration.id = 2

	fmt.Println(withoutNames)
	fmt.Println(withNames)
	fmt.Println(withVarDelcaration)
}
