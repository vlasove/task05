package store

import (
	"context"

	"github.com/vlasove/test05/internal/app/model"
)

// EmployeeRepository ...
type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	Delete(context.Context, int) error
	Update(context.Context, *model.Employee) error
	GetAll(context.Context) ([]*model.Employee, error)
	GetByID(context.Context, int) (*model.Employee, error)
}
