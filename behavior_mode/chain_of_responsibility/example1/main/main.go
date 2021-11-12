package main

import (
	"fmt"

	"go-patterns/behavior_mode/chain_of_responsibility/example1"
)

func main() {
	errLogger := example1.NewError(example1.ERROR)
	fileLogger := example1.NewFile(example1.DEBUG)
	consoleLogger := example1.NewConsole(example1.INFO)

	errLogger.SetNextLogger(fileLogger)
	fileLogger.SetNextLogger(consoleLogger)

	errLogger.LogMessage(example1.INFO, "this is an info")
	fmt.Println()
	errLogger.LogMessage(example1.DEBUG, "this is a debug")
	fmt.Println()
	errLogger.LogMessage(example1.ERROR, "this is an error")

}
