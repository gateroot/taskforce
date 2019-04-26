package task

type State interface {
	CurrentState() string
}

type stateTodo struct{}

func (s stateTodo) CurrentState() string {
	return "TODO"
}

type stateDoing struct{}

func (s stateDoing) CurrentState() string {
	return "DOING"
}

type statePaused struct{}

func (s statePaused) CurrentState() string {
	return "PAUSED"
}

type stateCompleted struct{}

func (s stateCompleted) CurrentState() string {
	return "COMPLETED"
}

type stateClosed struct{}

func (s stateClosed) CurrentState() string {
	return "CLOSED"
}
