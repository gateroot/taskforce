package domain

type Task struct {
	Id          string
	Title       string
	Description string
	Assignee    string
	State       State
}

func NewTask(title, description string) Task {
	return Task{
		// TODO: assign unique ID
		Id:          "1000",
		Title:       title,
		Description: description,
		Assignee:    "Bob",
		State:       StateTodo{},
	}
}

func (t Task) Assign(user string) Task {
	t.Assignee = user
	t.State = StateAssigned{}
	return t
}

func (t Task) UnAssign() Task {
	t.Assignee = ""
	t.State = StateTodo{}
	return t
}

func (t Task) Close() Task {
	t.State = StateClosed{}
	return t
}

func (t Task) CurrentState() string {
	return t.State.CurrentState()
}

type State interface {
	CurrentState() string
}

type StateTodo struct{}

func (s StateTodo) CurrentState() string {
	return "TODO"
}

type StateAssigned struct{}

func (s StateAssigned) CurrentState() string {
	return "ASSIGNED"
}

type StateDoing struct{}

func (s StateDoing) CurrentState() string {
	return "DOING"
}

type StatePaused struct{}

func (s StatePaused) CurrentState() string {
	return "PAUSED"
}

type StateCompleted struct{}

func (s StateCompleted) CurrentState() string {
	return "COMPLETED"
}

type StateClosed struct{}

func (s StateClosed) CurrentState() string {
	return "CLOSED"
}
