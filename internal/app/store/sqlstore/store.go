package sqlstore

import (
	"database/sql"

	"github.com/vlasove/test05/internal/app/store"
)

// Store ...
type Store struct {
	// config             *Config TO DELETE
	db                 *sql.DB
	employeeRepository *EmployeeRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Employee public method for managment repository
func (s *Store) Employee() store.EmployeeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}
	s.employeeRepository = &EmployeeRepository{store: s}
	return s.employeeRepository
}
