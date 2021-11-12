package main

import (
	"fmt"
	"go-patterns/behavior_mode/memento"
)

func main() {
	originator := &memento.Originator{}
	careTaker := memento.CareTaker{}

	originator.SetState("1")
	originator.SetState("2")

	careTaker.Add(originator.SaveStateToMemento())

	originator.SetState("3")
	careTaker.Add(originator.SaveStateToMemento())
	originator.SetState("4")

	fmt.Println("current state", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(0))

	fmt.Println("First saved state", originator.GetState())
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("Second saved state", originator.GetState())
}
