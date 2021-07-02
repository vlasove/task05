package store

// Store ...
type Store interface {
	Employee() EmployeeRepository
}
