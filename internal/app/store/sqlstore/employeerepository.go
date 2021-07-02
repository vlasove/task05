package sqlstore

import (
	"context"
	"log"

	"github.com/vlasove/test05/internal/app/model"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	store *Store
}

// Create call employees.employee_add()
func (r *EmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	if err := e.Validate(); err != nil {
		return err
	}
	_, err := r.store.db.ExecContext(
		ctx,
		"SELECT employees.employee_add($1, $2, $3, $4, $5, $6)",
		e.Name,
		e.LastName,
		e.Patronymic,
		e.Phone,
		e.Position,
		e.GoodJobCount,
	)
	return err
}

// Delete call employees.employee_remove()
func (r *EmployeeRepository) Delete(ctx context.Context, id int) error {
	_, err := r.store.db.ExecContext(ctx, "SELECT employees.employee_remove($1)", id)
	return err
}

// Update call employees.employee_upd()
func (r *EmployeeRepository) Update(ctx context.Context, e *model.Employee) error {
	if err := e.Validate(); err != nil {
		return err
	}
	_, err := r.store.db.ExecContext(
		ctx,
		"SELECT employees.employee_upd($1, $2, $3, $4, $5, $6, $7)",
		e.ID,
		e.Name,
		e.LastName,
		e.Patronymic,
		e.Phone,
		e.Position,
		e.GoodJobCount,
	)
	return err
}

// GetAll call employees.employee_get_all()
func (r *EmployeeRepository) GetAll(ctx context.Context) ([]*model.Employee, error) {
	empls := []*model.Employee{}
	rows, err := r.store.db.QueryContext(ctx, "SELECT * FROM employees.employees_get_all()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := new(model.Employee)
		if err := rows.Scan(
			&e.Name,
			&e.LastName,
			&e.ID,
			&e.Patronymic,
			&e.Phone,
			&e.Position,
			&e.GoodJobCount); err != nil {
			log.Println(err)
			continue
		}
		empls = append(empls, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return empls, nil
}

// GetByID call employees.employee_get()
func (r *EmployeeRepository) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	e := new(model.Employee)
	if err := r.store.db.QueryRowContext(ctx, "SELECT * FROM employees.employee_get($1)", id).Scan(
		&e.Name,
		&e.LastName,
		&e.ID,
		&e.Patronymic,
		&e.Phone,
		&e.Position,
		&e.GoodJobCount,
	); err != nil {

		return nil, err
	}
	return e, nil
}
