package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Employee ...
type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Patronymic   string `json:"patronymic"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	GoodJobCount int    `json:"good_job_count"`
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
