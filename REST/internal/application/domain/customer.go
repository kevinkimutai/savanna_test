package domain

import "errors"

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func NewCustomer(customer *Customer) (*Customer, error) {
	if customer.Name == "" {
		return customer, errors.New("missing customers values")
	}

	return customer, nil
}
