package log

import (
	"log"
	"os"
)

type level bool

func (l level) println(level string, v ...interface{}) {
	if l {
		newv := append([]interface{}{"[" + level + "]"}, v...)
		log.Println(newv...)
	}
}

func (l level) printf(level, format string, args ...interface{}) {
	if l {
		format = "[" + level + "] " + format
		log.Printf(format, args...)
	}
}

var (
	DebugLogger, InfoLogger, ErrorLogger, FatalLogger level = true, true, true, true
)

func Debug(v ...interface{}) {
	DebugLogger.println("DEBUG", v...)
}
func DebugF(format string, args ...interface{}) {
	DebugLogger.printf("DEBUG", format, args...)
}

func Info(v ...interface{}) {
	InfoLogger.println("INFO", v...)
}
func InfoF(format string, args ...interface{}) {
	InfoLogger.printf("INFO", format, args...)
}

func Error(v ...interface{}) {
	ErrorLogger.println("ERROR", v...)
}
func ErrorF(format string, args ...interface{}) {
	ErrorLogger.printf("ERROR", format, args...)
}

func Fatal(v ...interface{}) {
	FatalLogger.println("FATAL", v...)
	os.Exit(1)
}
func FatalF(format string, args ...interface{}) {
	FatalLogger.printf("FATAL", format, args...)
	os.Exit(1)
}
