package task

import (
	domain "taskforce/domain/task"
)

type Gateway interface {
	Create(title, description string) (int, error)
	Read(taskId int) (*domain.Task, error)
	Update(task domain.Task) error
	Delete(taskId int) error
}
