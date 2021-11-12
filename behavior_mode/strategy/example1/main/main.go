package main

import (
	"fmt"
	"go-patterns/behavior_mode/strategy"
)

func main() {
	context1 := strategy.New(new(strategy.OperationAdd))
	fmt.Println("10 + 5 = ", context1.ExecStrategy(10, 5))

	context2 := strategy.New(new(strategy.OperationSub))
	fmt.Println("10 - 5 = ", context2.ExecStrategy(10, 5))

	context3 := strategy.New(new(strategy.OperationMul))
	fmt.Println("10 * 5 = ", context3.ExecStrategy(10, 5))
}
