package main

import (
	"fmt"
	"time"

	"go-patterns/synchronization/semaphore"
)

func main() {
	ticket, timeout := 1, 2*time.Second
	sem := semaphore.New(ticket, timeout)

	if err := sem.Acquire(); err != nil {
		fmt.Println(err)
	}

	if err := sem.Release(); err != nil {
		fmt.Println(err)
	}
	if err := sem.Release(); err != nil {
		fmt.Println(err)
	}
}
