package domain

import (
	"errors"
	"regexp"
)

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func NewCustomer(customer *Customer) (*Customer, error) {
	if customer.Name == "" || customer.PhoneNumber == "" {
		return customer, errors.New("missing customer values")
	}

	isValid := CheckPhoneNumber(customer.PhoneNumber)
	if !isValid {
		return customer, errors.New("wrong phone number input")
	}

	return customer, nil
}

func CheckPhoneNumber(phoneNumber string) bool {
	pattern := `^\+2547\d{8}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Use the compiled regex to match the phone number
	return regex.MatchString(phoneNumber)
}
