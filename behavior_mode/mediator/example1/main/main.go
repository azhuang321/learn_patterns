package main

import "go-patterns/behavior_mode/mediator"

func main() {
	robert := mediator.NewUser("robert")
	john := mediator.NewUser("john")

	robert.SendMessage("hi this is robert")
	john.SendMessage("hi this is john")
}
