package domain

import "errors"

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewCustomer(customer *Customer) (*Customer, error) {
	if customer.FirstName == "" || customer.LastName == "" {
		return customer, errors.New("missing customers values")
	}

	return customer, nil
}
