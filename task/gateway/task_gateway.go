package gateway

import (
	"errors"
	"fmt"
	domain "github.com/hiroaki-sekine/taskforce/domain/task"
	component "github.com/hiroaki-sekine/taskforce/task"
)

type SqlTaskGateway struct {
	repository  component.Repository
	taskFactory domain.Factory
}

func NewSqlTaskGateway(repository component.Repository, taskFactory domain.Factory) component.Gateway {
	return &SqlTaskGateway{repository, taskFactory}
}

func (gw *SqlTaskGateway) Create(title, description string) (int, error) {
	query := fmt.Sprintf(
		`INSERT INTO TASKS (title, description, state) VALUES ("%s", "%s", "TODO")`, title, description)
	id, err := gw.repository.Create(query)
	return int(id), err
}

func (gw *SqlTaskGateway) Read(taskId int) (domain.Task, error) {
	query := fmt.Sprintf("SELECT id, title, description, state FROM TASKS WHERE id=%d", taskId)
	rows, err := gw.repository.ReadRow(query)
	if err != nil {
		return nil, err
	}
	var id int
	var title, description, state string
	err = rows.Scan(&id, &title, &description, &state)
	if err != nil {
		return nil, err
	}
	var s domain.State
	sf := gw.taskFactory.GetStateFactory()
	switch state {
	case "TODO":
		s = sf.Todo()
	case "DOING":
		s = sf.Doing()
	case "PAUSED":
		s = sf.Paused()
	case "COMPLETED":
		s = sf.Completed()
	case "CLOSED":
		s = sf.Closed()
	default:
		return nil, errors.New(fmt.Sprintf(`state "%s" is not supported"`, state))
	}
	task := gw.taskFactory.New(id, title, description, s)
	return task, nil
}

func (gw *SqlTaskGateway) Update(task domain.Task) error {
	query := fmt.Sprintf(`UPDATE TASKS SET title="%s", description="%s", state="%s" WHERE id=%d`, task.GetTitle(), task.GetDescription(), task.GetState().String(), task.GetId())
	return gw.repository.Update(query)
}

func (gw *SqlTaskGateway) Delete(taskId int) error {
	query := fmt.Sprintf(`DELETE FROM TASKS WHERE id=%d`, taskId)
	return gw.repository.Delete(query)
}
