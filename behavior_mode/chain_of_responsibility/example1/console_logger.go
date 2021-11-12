package example1

import "fmt"

type ConsoleLogger struct {
	AbstractLogger
}

func NewConsole(level int) *ConsoleLogger {
	logger := new(ConsoleLogger)
	logger.level = level
	logger.logger = logger
	return logger
}

func (c *ConsoleLogger) write(message string) {
	fmt.Println("standard console logger", message)
}
