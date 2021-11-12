package example1

import "fmt"

type ErrorLogger struct {
	AbstractLogger
}

func NewError(level int) *ErrorLogger {
	logger := new(ErrorLogger)
	logger.level = level
	logger.logger = logger
	return logger
}

func (c *ErrorLogger) write(message string) {
	fmt.Println("standard error logger", message)
}
