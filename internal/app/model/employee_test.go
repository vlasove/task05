package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/test05/internal/app/model"
)

func TestEmployee_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		employee func() *model.Employee
		isValid  bool
	}{
		{
			name: "valid",
			employee: func() *model.Employee {
				return model.TestEmployee(t)
			},
			isValid: true,
		},
		{
			name: "invalid good jobs",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.GoodJobCount = -1
				return e
			},
			isValid: false,
		},

		{
			name: "invalid name",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.Name = "tt"
				return e
			},
			isValid: false,
		},
		{
			name: "invalid lastname",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.LastName = "tt"
				return e
			},
			isValid: false,
		},
		{
			name: "invalid patronymic",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.Patronymic = "tt"
				return e
			},
			isValid: false,
		},

		{
			name: "invalid phone",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.Phone = "01234"
				return e
			},
			isValid: false,
		},
		{
			name: "invalid position",
			employee: func() *model.Employee {
				e := model.TestEmployee(t)
				e.Phone = "pos"
				return e
			},
			isValid: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.NoError(t, tt.employee().Validate())
			} else {
				assert.Error(t, tt.employee().Validate())
			}
		})
	}
}
