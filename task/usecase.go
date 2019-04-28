package task

type Usecase interface {
	New(title, description string) error
	Start(taskId int) error
	Stop(taskId int) error
	Complete(taskId int) error
	Close(taskId int) error
}
