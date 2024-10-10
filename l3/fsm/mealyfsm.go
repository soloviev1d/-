package fsm

type DeterministicFiniteStateMachine struct {
	Initial  StateOutputTuple
	current  StateOutputTuple
	StateMap StateMap
}

func NewFSM(initial StateOutputTuple, stateMap StateMap) *DeterministicFiniteStateMachine {
	return &DeterministicFiniteStateMachine{
		Initial:  initial,
		current:  initial,
		StateMap: stateMap,
	}
}

func (fsm *DeterministicFiniteStateMachine) Current() State {
	return fsm.current.S
}

func (fsm *DeterministicFiniteStateMachine) Output() string {
	return fsm.current.O
}

func (fsm *DeterministicFiniteStateMachine) Transition(event Event) error {
	fsm.current.O = fsm.StateMap[fsm.Current()][event].O
	fsm.current.S = fsm.StateMap[fsm.Current()][event].S
	return nil
}

func (fsm *DeterministicFiniteStateMachine) Reset() {
	fsm.current = fsm.Initial
}
