package model

import "testing"

// TestEmployee ...
func TestEmployee(t *testing.T) *Employee {
	t.Helper()
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
