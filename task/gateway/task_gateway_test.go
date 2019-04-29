package gateway

import (
	domain "github.com/hiroaki-sekine/taskforce/domain/task"
	"github.com/hiroaki-sekine/taskforce/task/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSqlTaskGateway_Create(t *testing.T) {
	taskRepo := mocks.Repository{}
	taskRepo.On("Create",
		`INSERT INTO TASKS (title, description, state) VALUES ("title", "description", "TODO")`,
	).Return(int64(1), nil)
	gw := SqlTaskGateway{&taskRepo, domain.NewTaskFactory(domain.NewStateFactory())}
	id, err := gw.Create("title", "description")
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestSqlTaskGateway_Read(t *testing.T) {
	rows := &mocks.Scanner{}
	rows.On(
		"Scan",
		mock.AnythingOfType("*int"),
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*string")).
		Return(nil).
		Run(func(args mock.Arguments) {
			id := args.Get(0).(*int)
			*id = 1
			title := args.Get(1).(*string)
			*title = "title"
			description := args.Get(2).(*string)
			*description = "description"
			state := args.Get(3).(*string)
			*state = "TODO"
		})

	taskRepo := mocks.Repository{}
	taskRepo.On("ReadRow", `SELECT id, title, description, state FROM TASKS WHERE id=1`).
		Return(rows, nil)

	gw := SqlTaskGateway{&taskRepo, domain.NewTaskFactory(domain.NewStateFactory())}

	ret, err := gw.Read(1)
	assert.NoError(t, err)
	assert.NotNil(t, ret)
	task := *ret
	assert.Equal(t, 1, task.GetId())
	assert.Equal(t, "title", task.GetTitle())
	assert.Equal(t, "description", task.GetDescription())
	assert.Equal(t, "TODO", task.GetState().String())
}

func TestSqlTaskGateway_Update(t *testing.T) {
	taskRepo := mocks.Repository{}
	taskRepo.On("Update", `UPDATE TASKS SET title="test", description="hello_world", state="DOING" WHERE id=1`).
		Return(nil)
	gw := SqlTaskGateway{&taskRepo, domain.NewTaskFactory(domain.NewStateFactory())}
	task := gw.taskFactory.New(1, "test", "hello_world", gw.taskFactory.GetStateFactory().Doing())
	err := gw.Update(task)
	assert.NoError(t, err)
}

func TestSqlTaskGateway_Delete(t *testing.T) {
	taskRepo := mocks.Repository{}
	taskRepo.On("Delete", `DELETE FROM TASKS WHERE id=1`).
		Return(nil)
	gw := SqlTaskGateway{&taskRepo, domain.NewTaskFactory(domain.NewStateFactory())}
	err := gw.Delete(1)
	assert.NoError(t, err)
}
