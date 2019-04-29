package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/hiroaki-sekine/taskforce/domain/task"
	"github.com/hiroaki-sekine/taskforce/task/mocks"
	"testing"
)

func TestTaskUsecase_New(t *testing.T) {
	gw := &mocks.Gateway{}
	uc := TaskUsecase{gw, task.NewTaskFactory(task.NewStateFactory())}
	gw.On("Create", "title", "description").Return(1, nil)
	err := uc.New("title", "description")
	assert.NoError(t, err)
}

// TODO: write error case
func TestTaskUsecase_Start(t *testing.T) {
	gw := &mocks.Gateway{}
	uc := TaskUsecase{gw, task.NewTaskFactory(task.NewStateFactory())}
	sf := task.NewStateFactory()

	gw.On("ReadRow", 1).Return(&task.MockTask{
		Id:          1,
		Title:       "title",
		Description: "description",
		State:       sf.Todo(),
	}, nil)

	gw.On("Update",
		mock.MatchedBy(func(t task.Task) bool {
			if t.GetId() == 1 &&
				t.GetTitle() == "title" &&
				t.GetDescription() == "description" &&
				t.GetState().String() == "DOING" {
				return true
			}
			return false
		})).
		Return(nil)

	err := uc.Start(1)
	assert.NoError(t, err)
}

// TODO: write error case
func TestTaskUsecase_Stop(t *testing.T) {
	gw := &mocks.Gateway{}
	uc := TaskUsecase{gw, task.NewTaskFactory(task.NewStateFactory())}
	sf := task.NewStateFactory()

	gw.On("ReadRow", 1).Return(&task.MockTask{
		Id:          1,
		Title:       "title",
		Description: "description",
		State:       sf.Doing(),
	}, nil)

	gw.On("Update",
		mock.MatchedBy(func(t task.Task) bool {
			if t.GetId() == 1 &&
				t.GetTitle() == "title" &&
				t.GetDescription() == "description" &&
				t.GetState().String() == "TODO" {
				return true
			}
			return false
		})).
		Return(nil)

	err := uc.Stop(1)
	assert.NoError(t, err)
}

// TODO: write error case
func TestTaskUsecase_Complete(t *testing.T) {
	gw := &mocks.Gateway{}
	uc := TaskUsecase{gw, task.NewTaskFactory(task.NewStateFactory())}
	sf := task.NewStateFactory()

	gw.On("ReadRow", 1).Return(&task.MockTask{
		Id:          1,
		Title:       "title",
		Description: "description",
		State:       sf.Doing(),
	}, nil)

	gw.On("Update",
		mock.MatchedBy(func(t task.Task) bool {
			if t.GetId() == 1 &&
				t.GetTitle() == "title" &&
				t.GetDescription() == "description" &&
				t.GetState().String() == "COMPLETED" {
				return true
			}
			return false
		})).
		Return(nil)

	err := uc.Complete(1)
	assert.NoError(t, err)
}

// TODO: write error case
func TestTaskUsecase_Close(t *testing.T) {
	gw := &mocks.Gateway{}
	uc := TaskUsecase{gw, task.NewTaskFactory(task.NewStateFactory())}
	sf := task.NewStateFactory()

	gw.On("ReadRow", 1).Return(&task.MockTask{
		Id:          1,
		Title:       "title",
		Description: "description",
		State:       sf.Completed(),
	}, nil)

	gw.On("Update",
		mock.MatchedBy(func(t task.Task) bool {
			if t.GetId() == 1 &&
				t.GetTitle() == "title" &&
				t.GetDescription() == "description" &&
				t.GetState().String() == "CLOSED" {
				return true
			}
			return false
		})).
		Return(nil)

	err := uc.Close(1)
	assert.NoError(t, err)
}
