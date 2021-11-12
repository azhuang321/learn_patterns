package example1

type Logger interface {
	SetNextLogger(logger Logger)
	LogMessage(level int, message string)
	write(message string)
}

const (
	INFO = iota
	DEBUG
	ERROR
)

type AbstractLogger struct {
	level      int
	nextLogger Logger
	logger     Logger
}

func (l *AbstractLogger) SetNextLogger(logger Logger) {
	l.nextLogger = logger
}
func (l *AbstractLogger) write(message string) {}

func (l *AbstractLogger) LogMessage(level int, message string) {
	if l.level <= level {
		l.logger.write(message)
	}
	if l.nextLogger != nil {
		l.nextLogger.LogMessage(level, message)
	}
}
