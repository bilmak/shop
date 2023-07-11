package core

import "errors"

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p Customer) ValidateCustomer() error {
	if p.Name == "" {
		return errors.New("requared field is empty")
	}
	return nil

}
