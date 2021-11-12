package main

import (
	"fmt"
	"go-patterns/behavior_mode/observer"
)

func main() {
	subject := new(observer.Subject)
	observer.NewHexa(subject)
	observer.NewBinary(subject)
	observer.NewOctal(subject)

	fmt.Println("first change state 15")
	subject.SetState(15)
	fmt.Println("second change state 10")
	subject.SetState(10)

}
