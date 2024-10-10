package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	FSM "github.com/soloviev1d/fsm-course/fsm"
)

func main() {
	// create a scanner for stdin
	scn := bufio.NewReader(os.Stdin)

	// define a Mealy finite state machine with initial state 0 | -
	// and state changes described by FSM.StateMap
	fsm := FSM.NewFSM(FSM.StateOutputTuple{"0", "-"}, FSM.StateMap{
		"0": FSM.StateTransition{
			"50":  FSM.StateOutputTuple{"50", "-"},
			"100": FSM.StateOutputTuple{"100", "-"},
			"200": FSM.StateOutputTuple{"200", "-"},
		},
		"50": FSM.StateTransition{
			"50":  FSM.StateOutputTuple{"100", "-"},
			"100": FSM.StateOutputTuple{"150", "-"},
			"200": FSM.StateOutputTuple{"0", "0"},
		},
		"100": FSM.StateTransition{
			"50":  FSM.StateOutputTuple{"150", "-"},
			"100": FSM.StateOutputTuple{"200", "-"},
			"200": FSM.StateOutputTuple{"0", "50"},
		},
		"150": FSM.StateTransition{
			"50":  FSM.StateOutputTuple{"200", "-"},
			"100": FSM.StateOutputTuple{"0", "0"},
			"200": FSM.StateOutputTuple{"0", "100"},
		},
		"200": FSM.StateTransition{
			"50":  FSM.StateOutputTuple{"0", "0"},
			"100": FSM.StateOutputTuple{"0", "50"},
			"200": FSM.StateOutputTuple{"0", "150"},
		},
	})

	// transition state until output signal is equal to something other than "-"
	for fsm.Output() == "-" {
		// print prompt and read user input
		fmt.Print("> ")
		l, err := scn.ReadString('\n')
		if err != nil {
			log.Fatalln("Failed to read string", err)
		}

		// trim '\n'
		l = l[:len(l)-1]

		// check if user entered a valid value according to possible transitions
		if _, ok := fsm.StateMap[fsm.Current()][FSM.Event(l)]; !ok {
			fmt.Println("Invalid input:", l, "State was not changed")
			continue
		}

		// transition state
		fsm.Transition(FSM.Event(l))
		fmt.Println("Inserted:", fsm.Current())
	}

	fmt.Println("Product purchased, change:", fsm.Output())
}
