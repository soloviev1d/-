package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	FSM "github.com/soloviev1d/fsm-course/fsm"
)

func main() {
	scn := bufio.NewReader(os.Stdin)
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

	for fsm.Output() == "-" {
		fmt.Print("> ")
		l, err := scn.ReadString('\n')
		if err != nil {
			log.Fatalln("Failed to read string", err)
		}
		l = l[:len(l)-1]

		if _, ok := fsm.StateMap[fsm.Current()][FSM.Event(l)]; !ok {
			fmt.Println("Invalid input:", l, "State was not changed")
			continue
		}
		fsm.Transition(FSM.Event(l))
		fmt.Println("Inserted:", fsm.Current())
	}

	fmt.Println("Product purchased, change:", fsm.Output())
}
