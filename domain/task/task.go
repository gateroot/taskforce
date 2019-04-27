package task

type Task struct {
	id          int
	Title       string
	Description string
	state       State
}

func NewTask(title, description string) *Task {
	return &Task{
		// TODO: assign unique ID
		id:          1,
		Title:       title,
		Description: description,
		state:       stateTodo{},
	}
}

func (t *Task) Stop() *Task {
	t.state = stateTodo{}
	return t
}

func (t *Task) Start() *Task {
	t.state = stateDoing{}
	return t
}

func (t *Task) Pause() *Task {
	t.state = statePaused{}
	return t
}

func (t *Task) Complete() *Task {
	t.state = stateCompleted{}
	return t
}

func (t *Task) Close() *Task {
	t.state = stateClosed{}
	return t
}

func (t *Task) CurrentState() string {
	return t.state.CurrentState()
}
