// Реализация паттерна "цепочка вызовов" на примере абстрактного логгера
package pattern

import (
	"fmt"
)

// Уровни логгирования
const (
	LevelInfo = iota
	LevelDebug
	LevelError
)

// Logger - интерфейс для всех логгеров
type Logger interface {
	Log(int, string)
	AppendNext(Logger)
}

// AbstractLogger - структура с полями, общими для всех логгеров
type AbstractLogger struct {
	nextLogger Logger
	level int
}

// AppendNext - добавление нового логгера в цепочку
func (logger *AbstractLogger) AppendNext(next Logger) {
	if logger.nextLogger != nil {
		logger.nextLogger.AppendNext(next)
		return
	}
	logger.nextLogger = next
} 





// ConsoleLogger - логгер, пишущий в консоль
type ConsoleLogger struct {
	AbstractLogger
}

// NewConsoleLogger - конуструктор ConsoleLogger
func NewConsoleLogger(level int) *ConsoleLogger {
	return &ConsoleLogger{
		AbstractLogger{
			level: level,
		},
	}
} 

// Log симулирует обработку сообщения
func (logger *ConsoleLogger) Log(level int, s string) {
	if level >= logger.level {
		fmt.Println("Console logger: " + s)
	}
	if logger.nextLogger != nil {
		logger.nextLogger.Log(level, s)
	}
	
}




// ErrorLogger - логгер, пишущий в stderr
type ErrorLogger struct {
	AbstractLogger
}

// NewErrorLogger - конструктор ErrorLogger
func NewErrorLogger(level int) *ErrorLogger {
	return &ErrorLogger{
		AbstractLogger{
			level: level,
		},
	}
} 

// Log симулирует обработку сообщения
func (logger *ErrorLogger) Log(level int, s string) {
	if level >= logger.level {
		fmt.Println("Error logger: " + s)
	}
	if logger.nextLogger != nil {
		logger.nextLogger.Log(level, s)
	}
}





// FileLogger - логгер, пишущий в файл
type FileLogger struct {
	AbstractLogger
}

// NewFileLogger - конструктор FileLogger
func NewFileLogger(level int) *FileLogger {
	return &FileLogger{
		AbstractLogger{
			level: level,
		},
	}
} 

// Log симулирует обработку сообщения
func (logger *FileLogger) Log(level int, s string) {
	if level >= logger.level {
		fmt.Println("File logger: " + s)
	}
	if logger.nextLogger != nil {
		logger.nextLogger.Log(level, s)
	}
}



// Пример использования
/*func main() {
	var logger = NewConsoleLogger(LevelInfo)
	logger.AppendNext(NewFileLogger(LevelDebug))
	logger.AppendNext(NewErrorLogger(LevelError))

	logger.Log(LevelInfo, "<Info message>")
	logger.Log(LevelDebug, "<Debug message>")
	logger.Log(LevelError, "<Error message>")
}*/