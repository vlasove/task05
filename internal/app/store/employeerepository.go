package store

import "github.com/vlasove/test05/internal/app/model"

// EmployeeRepository ...
type EmployeeRepository struct {
	store *Store
}

// Create call employees.employee_add()
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	r.store.db.QueryRow("SELECT employees.employee_add()")
	return nil, nil
}

// Delete call employee_remove()
func (r *EmployeeRepository) Delete(id int) error {
	return nil
}

// Update call employees.employee_upd()
func (r *EmployeeRepository) Update(e *model.Employee) (*model.Employee, error) {
	return nil, nil
}

// GetAll call employees.employee_get_all()
func (r *EmployeeRepository) GetAll() ([]*model.Employee, error) {
	return nil, nil
}

// GetByID call employees.employee_get()
func (r *EmployeeRepository) GetByID(id int) (*model.Employee, error) {
	return nil, nil
}
