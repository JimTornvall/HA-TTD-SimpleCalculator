package calculator

import "log"

type Logger interface{ Log(int) }

type SimpleLogger struct{}

func (SimpleLogger) Log(num int) {

	// using log from std package to log the number
	log.Println("Number over 1000: ", num)
}

func NewSimpleLogger() SimpleLogger {
	return SimpleLogger{}
}
