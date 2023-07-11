package core

import "errors"

type Employees struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c Employees) ValidateEmployees() error {
	if c.Name == "" {
		return errors.New("requared field is empty")
	}
	return nil

}
