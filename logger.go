package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

//Logger describes how we should do logging in this application.
type Logger interface {
	Printf(format string, args ...interface{})
}

func valuesPrettyPrint(indentCount int, values map[string][]string) string {
	var out string
	for k, v := range values {
		out += fmt.Sprintf("%s %s => %v\n", strings.Repeat(" ", indentCount), k, v)
	}
	return out
}
//NewLogger creates a new logger.
func NewLogger(logger string) Logger{
	switch (logger) {
	case "logrus":
		return NewLogrusLogger()
	default:
		return NewLogrusLogger()
	}
}
//NewLogrusLogger creates new loggrus based logger. 
func NewLogrusLogger() Logger {
	return logrus.New()
}