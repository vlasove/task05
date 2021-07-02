package teststore

import (
	"errors"

	"github.com/vlasove/test05/internal/app/model"
)

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found in database")
)

// Store ...
type Store struct {
	employeeRepository *EmployeeRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Employee public method for managment repository
func (s *Store) Employee() *EmployeeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}
	s.employeeRepository = &EmployeeRepository{
		store:     s,
		employees: make(map[int]*model.Employee),
	}
	return s.employeeRepository
}
