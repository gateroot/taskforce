package task

type MockTask struct {
	Id          int
	Title       string
	Description string
	State       State
}

func (t *MockTask) GetId() int {
	return t.Id
}

func (t *MockTask) GetTitle() string {
	return t.Title
}

func (t *MockTask) GetDescription() string {
	return t.Description
}

func (t *MockTask) GetState() State {
	return t.State
}

func (t *MockTask) setState(state State) {
	t.State = state
}
