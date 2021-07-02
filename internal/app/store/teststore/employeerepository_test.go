package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/test05/internal/app/model"
	"github.com/vlasove/test05/internal/app/store/teststore"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	assert.NoError(t, s.Employee().Create(e))
}

func TestEmployeeRepository_GetByID(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	id := 1
	e.ID = id
	t.Run("not existing employee", func(t *testing.T) {
		empl, err := s.Employee().GetByID(id)
		assert.Error(t, err, teststore.ErrRecordNotFound.Error())
		assert.Nil(t, empl)
	})

	t.Run("existing user", func(t *testing.T) {
		err := s.Employee().Create(e)
		assert.NoError(t, err)

		empl, err := s.Employee().GetByID(id)
		assert.NoError(t, err)
		assert.NotNil(t, empl)
	})
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	s := teststore.New()
	e1 := model.TestEmployee(t)
	e2 := model.TestEmployee(t)

	t.Run("no employee in database", func(t *testing.T) {
		empls, err := s.Employee().GetAll()
		assert.NoError(t, err)
		assert.Equal(t, len(empls), 0)

	})

	t.Run("two employees in database", func(t *testing.T) {
		err := s.Employee().Create(e1)
		assert.NoError(t, err)
		err = s.Employee().Create(e2)
		assert.NoError(t, err)

		empls, err := s.Employee().GetAll()
		assert.NoError(t, err)
		assert.Equal(t, len(empls), 2)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	assert.NoError(t, s.Employee().Create(e))
	assert.NoError(t, s.Employee().Delete(1))

}

func TestEmployeeRepository_Update(t *testing.T) {
	s := teststore.New()
	originEmployee := model.TestEmployee(t)
	updatedEmployee := model.TestEmployee(t)
	newName := "Updated Name"
	updatedEmployee.Name = newName
	updatedEmployee.ID = 1

	t.Run("update not existing user", func(t *testing.T) {
		err := s.Employee().Update(updatedEmployee)
		assert.Error(t, err, teststore.ErrRecordNotFound.Error())
	})

	t.Run("update existing user", func(t *testing.T) {
		err := s.Employee().Create(originEmployee)
		assert.NoError(t, err)

		err = s.Employee().Update(updatedEmployee)
		assert.NoError(t, err)

		emplInDB, err := s.Employee().GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, emplInDB.Name, newName)
	})

}
