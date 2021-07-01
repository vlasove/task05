package store

import (
	"log"

	"github.com/vlasove/test05/internal/app/model"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	store *Store
}

// getLastID helper for Create method
func (r *EmployeeRepository) getLastID() (int, error) {
	var id int
	if err := r.store.db.QueryRow(
		"SELECT CURRVAL('employees.employees_id_seq')",
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

// Create call employees.employee_add()
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if _, err := r.store.db.Exec(
		"SELECT employees.employee_add($1, $2, $3, $4, $5, $6)",
		e.Name,
		e.LastName,
		e.Patronymic,
		e.Phone,
		e.Position,
		e.GoodJobCount,
	); err != nil {
		return nil, err
	}
	id, err := r.getLastID()
	if err != nil {
		return nil, err
	}
	e.ID = id
	return nil, nil
}

// Delete call employees.employee_remove()
func (r *EmployeeRepository) Delete(id int) error {
	if _, err := r.store.db.Exec("SELECT employees.employee_remove($1)", id); err != nil {
		return err
	}
	return nil
}

// Update call employees.employee_upd()
func (r *EmployeeRepository) Update(e *model.Employee) (*model.Employee, error) {
	if _, err := r.store.db.Exec(
		"SELECT employees.employee_update($1, $2, $3, $4, $5, $6, $7)",
		e.ID,
		e.Name,
		e.LastName,
		e.Patronymic,
		e.Phone,
		e.Position,
		e.GoodJobCount,
	); err != nil {
		return nil, err
	}
	return e, nil
}

// GetAll call employees.employee_get_all()
func (r *EmployeeRepository) GetAll() ([]*model.Employee, error) {
	empls := []*model.Employee{}
	rows, err := r.store.db.Query("SELECT employees.employees_get_all()")
	if err != nil {
		return empls, err
	}
	defer rows.Close()

	for rows.Next() {
		e := new(model.Employee)
		if err := rows.Scan(&e.Name); err != nil {
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
func (r *EmployeeRepository) GetByID(id int) (*model.Employee, error) {
	e := new(model.Employee)
	if err := r.store.db.QueryRow("SELECT employees.employee_get($1)", id).Scan(
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
