package domain

import "errors"

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewCustomer(customer *Customer) (*Customer, error) {
	if customer.Name == "" {
		return customer, errors.New("missing customers values")
	}

	return customer, nil
}
