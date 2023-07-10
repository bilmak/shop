package core

import "errors"

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p Product) Validate() error {
	if p.Name == "" || p.Price == 0 {
		return errors.New("requared field is empty")
	}
	return nil

}
