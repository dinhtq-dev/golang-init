package logger

import (
	"log"
	"runtime"

	"github.com/natefinch/lumberjack"
)

// Logger struct to manage loggers
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

const logFilePath = "logs/app.log"

// NewLogger initializes the logger and opens the log file with log rotation
func NewLogger() (*Logger, error) {
	// Using lumberjack to automatically rotate log files
	logFile := &lumberjack.Logger{
		Filename:   logFilePath, // Path to the log file
		MaxSize:    10,          // Max size of the log file in MB
		MaxBackups: 5,           // Number of backup log files to keep
		MaxAge:     30,          // Retain logs for 30 days
		Compress:   true,        // Compress old log files
	}

	// Create loggers for info, warning, and error levels
	return &Logger{
		infoLogger:    log.New(logFile, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(logFile, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(logFile, "ERROR: ", log.Ldate|log.Ltime),
	}, nil
}

// Info logs informational messages with file and line number
func (l *Logger) Info(message string) {
	_, file, line, _ := runtime.Caller(1)
	l.infoLogger.Printf("%s:%d: %s", file, line, message)
}

// Warning logs warning messages with file and line number
func (l *Logger) Warning(message string) {
	_, file, line, _ := runtime.Caller(1)
	l.warningLogger.Printf("%s:%d: %s", file, line, message)
}

// Error logs error messages with file and line number
func (l *Logger) Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	l.errorLogger.Printf("%s:%d: %s", file, line, message)
}
