package server

import "log"

type Logger struct {
}

func (l *Logger) Info(msg string, args ...any) {
	log.Printf(msg, args...)
}
