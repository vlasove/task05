package model

import "testing"

func TestEmployee(t *testing.T) *Employee {
	return &Employee{
		ID:           7,
		Name:         "TestName",
		LastName:     "TestLastName",
		Patronymic:   "TestPatronymic",
		Phone:        "+7-999-999-99-99",
		Position:     "TestPosition",
		GoodJobCount: 7,
	}
}
