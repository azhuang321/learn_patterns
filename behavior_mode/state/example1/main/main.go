package main

import "go-patterns/behavior_mode/state"

func main() {
	context := state.New()

	startState := new(state.StartState)
	startState.DoAction(*context)

	stopState := new(state.StopState)
	stopState.DoAction(*context)
}
