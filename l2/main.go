package main

import (
	"fmt"

	FSM "github.com/soloviev1d/fsm-course/fsm"
)

func main() {
	fsm := FSM.NewFSM("S0", FSM.StateMap{
		"S0": FSM.StateTransition{
			"0": "S0",
			"1": "S1",
		},
		"S1": FSM.StateTransition{
			"0": "S0",
			"1": "S2",
		},
		"S2": FSM.StateTransition{
			"0": "S0",
			"1": "S2",
		},
	})

	events1 := []FSM.Event{"1", "0", "1", "1", "0", "0"}
	events2 := []FSM.Event{"1", "0", "1", "1", "0", "1"}
	events3 := []FSM.Event{"1", "0", "1", "1", "1", "1"}

	state1, _ := FSM.FeedThrough(fsm, events1)
	fsm.Reset()
	state2, _ := FSM.FeedThrough(fsm, events2)
	fsm.Reset()
	state3, _ := FSM.FeedThrough(fsm, events3)
	fsm.Reset()

	fmt.Printf("Actual: %s, expected: !S2\n", state1)
	fmt.Printf("Actual: %s, expected: !S2\n", state2)
	fmt.Printf("Actual: %s, expected: S2\n", state3)

}
