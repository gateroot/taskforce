package task

type Task interface {
	GetId() int
	GetTitle() string
	GetDescription() string
	GetState() State
	setState(state State)
}

type task struct {
	id          int
	title       string
	description string
	state       State
}

func (t *task) GetId() int {
	return t.id
}

func (t *task) GetTitle() string {
	return t.title
}

func (t *task) GetDescription() string {
	return t.description
}

func (t *task) GetState() State {
	return t.state
}

func (t *task) setState(state State) {
	t.state = state
}
