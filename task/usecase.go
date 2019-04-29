package task

import "github.com/hiroaki-sekine/taskforce/domain/task"

type Usecase interface {
	New(title, description string) error
	Start(taskId int) error
	Stop(taskId int) error
	Pause(taskId int) error
	Complete(taskId int) error
	Close(taskId int) error
	Delete(taskId int) error
}

type ViewTaskUsecase interface {
	Get(taskId int) (task.Task, error)
	List() ([]task.Task, error)
}
