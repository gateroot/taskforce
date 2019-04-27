package task

type State interface {
	String() string
}

type stateTodo struct{}

func (s stateTodo) String() string {
	return "TODO"
}

type stateDoing struct{}

func (s stateDoing) String() string {
	return "DOING"
}

type statePaused struct{}

func (s statePaused) String() string {
	return "PAUSED"
}

type stateCompleted struct{}

func (s stateCompleted) String() string {
	return "COMPLETED"
}

type stateClosed struct{}

func (s stateClosed) String() string {
	return "CLOSED"
}
