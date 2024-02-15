package main

type Person struct {
	name string
	age  int
}

func main() {
	people := make([]Person, 0, 10_000_000)

	for range 10_000_000 {
		people = append(people, Person{"John", 69})
	}
}
