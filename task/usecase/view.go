package usecase

import (
	"errors"
	task2 "github.com/hiroaki-sekine/taskforce/domain/task"
	"github.com/hiroaki-sekine/taskforce/task"
)

type ViewTaskUsecase struct {
	gateway task.ViewGateway
}

func NewViewTaskUsecase(gateway task.ViewGateway) task.ViewTaskUsecase {
	return &ViewTaskUsecase{gateway}
}

func (uc *ViewTaskUsecase) Get(taskId int) (task2.Task, error) {
	t, err := uc.gateway.Get(taskId)
	if err != nil {
		return nil, errors.New("get task failed.\n")
	}
	return t, nil
}

func (uc *ViewTaskUsecase) List() ([]task2.Task, error) {
	tasks, err := uc.gateway.List()
	if err != nil {
		return nil, errors.New("list task failed.\n")
	}
	return tasks, nil
}
