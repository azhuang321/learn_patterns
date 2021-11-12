package example1

import "fmt"

type FileLogger struct {
	AbstractLogger
}

func NewFile(level int) *FileLogger {
	logger := new(FileLogger)
	logger.level = level
	logger.logger = logger
	return logger
}

func (c *FileLogger) write(message string) {
	fmt.Println("standard file logger", message)
}
