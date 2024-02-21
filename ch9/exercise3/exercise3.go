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
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}

				message = message + " allErrors: " + strings.Join(messages, ", ")
			default:
				message = message + " error: " + processError(err, emp)
			}
		}

		fmt.Println(message)
	}
}

func processError(err error, emp Employee) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidId) {
		return fmt.Sprintf("invalid Id: %s", emp.Id)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
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
	var allErrors []error
	if len(e.Id) == 0 {
		allErrors = append(allErrors, EmptyFieldError{FieldName: "Id"})
	}

	if !validId.MatchString(e.Id) {
		allErrors = append(allErrors, ErrInvalidId)
	}

	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}
