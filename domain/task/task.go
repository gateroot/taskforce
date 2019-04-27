package task

type Task interface {
	GetTitle() string
	GetDescription() string
	GetState() State
	setState(state State) Task
}

type task struct {
	id          int
	title       string
	description string
	state       State
}

func (t task) GetTitle() string {
	return t.title
}

func (t task) GetDescription() string {
	return t.description
}

func (t task) GetState() State {
	return t.state
}

func (t task) setState(state State) Task {
	t.state = state
	return t
}

