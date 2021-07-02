package teststore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/test05/internal/app/model"
	"github.com/vlasove/test05/internal/app/store/teststore"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	assert.NoError(t, s.Employee().Create(context.Background(), e))
}

func TestEmployeeRepository_GetByID(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	id := 1
	e.ID = id
	t.Run("not existing employee", func(t *testing.T) {
		empl, err := s.Employee().GetByID(context.Background(), id)
		assert.Error(t, err, teststore.ErrRecordNotFound.Error())
		assert.Nil(t, empl)
	})

	t.Run("existing user", func(t *testing.T) {
		err := s.Employee().Create(context.Background(), e)
		assert.NoError(t, err)

		empl, err := s.Employee().GetByID(context.Background(), id)
		assert.NoError(t, err)
		assert.NotNil(t, empl)
	})
}

func TestEmployeeRepository_GetAll(t *testing.T) {
	s := teststore.New()
	e1 := model.TestEmployee(t)
	e2 := model.TestEmployee(t)

	t.Run("no employee in database", func(t *testing.T) {
		empls, err := s.Employee().GetAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(empls), 0)

	})

	t.Run("two employees in database", func(t *testing.T) {
		err := s.Employee().Create(context.Background(), e1)
		assert.NoError(t, err)
		err = s.Employee().Create(context.Background(), e2)
		assert.NoError(t, err)

		empls, err := s.Employee().GetAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(empls), 2)
	})
}

func TestEmployeeRepository_Delete(t *testing.T) {
	s := teststore.New()
	e := model.TestEmployee(t)
	assert.NoError(t, s.Employee().Create(context.Background(), e))
	assert.NoError(t, s.Employee().Delete(context.Background(), 1))

}

func TestEmployeeRepository_Update(t *testing.T) {
	s := teststore.New()
	originEmployee := model.TestEmployee(t)
	updatedEmployee := model.TestEmployee(t)
	newName := "Updated Name"
	updatedEmployee.Name = newName
	updatedEmployee.ID = 1

	t.Run("update not existing user", func(t *testing.T) {
		err := s.Employee().Update(context.Background(), updatedEmployee)
		assert.Error(t, err, teststore.ErrRecordNotFound.Error())
	})

	t.Run("update existing user", func(t *testing.T) {
		err := s.Employee().Create(context.Background(), originEmployee)
		assert.NoError(t, err)

		err = s.Employee().Update(context.Background(), updatedEmployee)
		assert.NoError(t, err)

		emplInDB, err := s.Employee().GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, emplInDB.Name, newName)
	})

}
