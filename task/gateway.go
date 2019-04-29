package task

import (
	domain "github.com/hiroaki-sekine/taskforce/domain/task"
)

type Gateway interface {
	Create(title, description string) (int, error)
	Read(taskId int) (domain.Task, error)
	Update(task domain.Task) error
	Delete(taskId int) error
}

type ViewGateway interface {
	Get(taskId int) (domain.Task, error)
	List() ([]domain.Task, error)
}
