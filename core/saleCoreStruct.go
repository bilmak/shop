package core

import (
	"errors"
	"time"
)

type Sales struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	CustomerID  int       `json:"customerid"`
	EmployeesID int       `json:"employeesid`
	ProductID   int       `json:"productid"`
	UnitPrice   int       `json:"unitprice"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"totalprice"`
}

func (s Sales) ValidateSale() error {
	if s.Date.IsZero() || s.CustomerID == 0 ||
		s.EmployeesID == 0 || s.ProductID == 0 || s.UnitPrice == 0 || s.Quantity == 0 || s.TotalPrice == 0 {
		return errors.New("sales validate")

	}
	return nil
}
