package gateway

import (
	"errors"
	"fmt"
	"log"
	task2 "taskforce/domain/task"
	"taskforce/task"
)

type ViewTaskGateway struct {
	repository task.Repository
	factory    task2.Factory
}

func NewViewTaskGateway(repository task.Repository, factory task2.Factory) task.ViewGateway {
	return &ViewTaskGateway{repository, factory}
}

func (gw *ViewTaskGateway) Get(taskId int) (task2.Task, error) {
	query := fmt.Sprintf("SELECT id, title, description, state FROM TASKS WHERE id=%d", taskId)
	row, err := gw.repository.ReadRow(query)
	if err != nil {
		return nil, err
	}
	var id int
	var title, description, state string
	err = row.Scan(&id, &title, &description, &state)
	if err != nil {
		return nil, err
	}
	var s task2.State
	sf := gw.factory.GetStateFactory()
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
	task := gw.factory.New(id, title, description, s)
	return task, nil
}

func (gw *ViewTaskGateway) List() ([]task2.Task, error) {
	query := fmt.Sprint("SELECT id, title, description, state FROM TASKS")
	rows, err := gw.repository.ReadRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]task2.Task, 0)
	for rows.Next() {
		var id int
		var title string
		var description string
		var state string
		if err := rows.Scan(&id, &title, &description, &state); err != nil {
			return nil, err
		}
		var s task2.State
		sf := gw.factory.GetStateFactory()
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
			log.Fatal(err)
		}
		t := gw.factory.New(id, title, description, s)
		tasks = append(tasks, t)
	}
	return tasks, nil
}
