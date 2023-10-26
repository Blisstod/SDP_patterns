package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logs []string
}

var instance *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
	once.Do(func() {
		instance = &Logger{logs: []string{}}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.logs = append(l.logs, message)
}

func (l *Logger) PrintLogs() {
	fmt.Println("Log Messages:")
	for _, log := range l.logs {
		fmt.Println(log)
	}
}

func main() {
	l := GetLoggerInstance()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		l := GetLoggerInstance()
		l.Log("Message from Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		l := GetLoggerInstance()
		l.Log("Message from Goroutine 2")
	}()

	wg.Wait()

	l.PrintLogs()
}
