package task

type StateFactory interface {
	Todo() State
	Doing() State
	Paused() State
	Completed() State
	Closed() State
}

type stateFactory struct{}

func NewStateFactory() StateFactory {
	return &stateFactory{}
}

func (sf *stateFactory) Todo() State {
	return stateTodo{}
}

func (sf *stateFactory) Doing() State {
	return stateDoing{}
}

func (sf *stateFactory) Paused() State {
	return statePaused{}
}

func (sf *stateFactory) Completed() State {
	return stateCompleted{}
}

func (sf *stateFactory) Closed() State {
	return stateClosed{}
}
