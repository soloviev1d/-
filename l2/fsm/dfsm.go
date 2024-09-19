package fsm

type DeterministicFiniteStateMachine struct {
	Initial  State
	current  State
	StateMap StateMap
}

func NewFSM(initial State, stateMap StateMap) *DeterministicFiniteStateMachine {
	return &DeterministicFiniteStateMachine{
		Initial:  initial,
		current:  initial,
		StateMap: stateMap,
	}
}

func (fsm *DeterministicFiniteStateMachine) Current() State {
	return fsm.current
}

func (fsm *DeterministicFiniteStateMachine) Transition(event Event) error {
	fsm.current = fsm.StateMap[fsm.Current()][event]
	return nil
}

func (fsm *DeterministicFiniteStateMachine) Reset() {
	fsm.current = fsm.Initial
}
