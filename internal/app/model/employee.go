package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Employee ...
type Employee struct {
	ID           int
	Name         string
	LastName     string
	Patronymic   string
	Phone        string
	Position     string
	GoodJobCount int
}

// Validate ...
func (e *Employee) Validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(&e.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.LastName, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.Patronymic, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.Phone, validation.Required, validation.Length(7, 100)),
		validation.Field(&e.Position, validation.Required, validation.Length(3, 100)),
		validation.Field(&e.GoodJobCount, validation.Required, validation.Min(0)),
	)
}
