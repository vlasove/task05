package teststore

import (
	"context"

	"github.com/vlasove/test05/internal/app/model"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	store     *Store
	employees map[int]*model.Employee
}

// Create ...
func (r *EmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	if err := e.Validate(); err != nil {
		return err
	}
	e.ID = len(r.employees) + 1
	r.employees[e.ID] = e
	return nil
}

// Delete ...
func (r *EmployeeRepository) Delete(ctx context.Context, id int) error {
	_, ok := r.employees[id]
	if !ok {
		return ErrRecordNotFound
	}
	delete(r.employees, id)
	return nil
}

// Update ...
func (r *EmployeeRepository) Update(ctx context.Context, e *model.Employee) error {
	_, ok := r.employees[e.ID]
	if !ok {
		return ErrRecordNotFound
	}
	r.employees[e.ID] = e
	return nil
}

// GetAll ...
func (r *EmployeeRepository) GetAll(ctx context.Context) ([]*model.Employee, error) {
	ans := []*model.Employee{}
	for _, e := range r.employees {
		ans = append(ans, e)
	}
	return ans, nil
}

// GetByID ...
func (r *EmployeeRepository) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	e, ok := r.employees[id]
	if !ok {
		return nil, ErrRecordNotFound
	}
	return e, nil
}
