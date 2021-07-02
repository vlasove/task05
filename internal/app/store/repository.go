package store

import "github.com/vlasove/test05/internal/app/model"

// EmployeeRepository ...
type EmployeeRepository interface {
	Create(*model.Employee) error
	Delete(int) error
	Update(*model.Employee) error
	GetAll() ([]*model.Employee, error)
	GetByID(int) (*model.Employee, error)
}
