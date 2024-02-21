package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var ErrInvalidId = errors.New("invalid id")

type EmptyFieldError struct {
	FieldName string
}

func (fe EmptyFieldError) Error() string {
	return fe.FieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}

		err = ValidateEmployee(emp)
		var fieldErr EmptyFieldError
		if err != nil {
			if errors.Is(err, ErrInvalidId) {
				fmt.Printf("record %d: %+v error: invalid Id: %s\n", count, emp, emp.Id)
			} else if errors.As(err, &fieldErr) {
				fmt.Printf("record %d: %+v error: empty field %s\n", count, emp, fieldErr.FieldName)
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			}

			continue
		}

		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Doe",
	"title": "Manager"
}
{
	"id": "",
	"first_name": "John",
	"last_name": "Doe",
	"title": "Employee"
}`

var validId = regexp.MustCompile(`\w{4}-\d{3}`)

func ValidateEmployee(e Employee) error {
	if len(e.Id) == 0 {
		return EmptyFieldError{FieldName: "Id"}
	}

	if !validId.MatchString(e.Id) {
		return ErrInvalidId
	}

	return nil
}
