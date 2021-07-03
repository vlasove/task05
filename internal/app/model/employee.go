package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Employee ...
type Employee struct {
	ID           int    `json:"id" xml:"id,attr"`
	Name         string `json:"name" xml:"name,attr"`
	LastName     string `json:"last_name" xml:"last_name,attr"`
	Patronymic   string `json:"patronymic" xml:"patronymic,attr"`
	Phone        string `json:"phone" xml:"phone,attr"`
	Position     string `json:"position" xml:"position,attr"`
	GoodJobCount int    `json:"good_job_count" xml:"good_job_count,attr"`
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
