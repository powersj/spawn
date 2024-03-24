package main

import (
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar = logger.Sugar()
}

type Logger struct {
	Name string
}

func NewLogger(family string, name string, alias string) *Logger {
	return &Logger{
		Name: createLogName(family, name, alias),
	}
}

// Debug logs a debug message, patterned after log.Print.
func (l *Logger) Debug(args ...interface{}) {
	sugar.Debug(append([]interface{}{"D! [" + l.Name + "] "}, args...)...)
}

// Debugf logs a debug message, patterned after log.Printf.
func (l *Logger) Debugf(format string, args ...interface{}) {
	sugar.Debugf("D! ["+l.Name+"] "+format, args...)
}

// Error logs a error message, patterned after log.Print.
func (l *Logger) Error(args ...interface{}) {
	sugar.Error(append([]interface{}{"W! [" + l.Name + "] "}, args...)...)
}

// Errorf logs a errorf message, patterned after log.Printf.
func (l *Logger) Errorf(format string, args ...interface{}) {
	sugar.Errorf("W! ["+l.Name+"] "+format, args...)
}

// Info logs an information message, patterned after log.Print.
func (l *Logger) Info(args ...interface{}) {
	sugar.Info(append([]interface{}{"I! [" + l.Name + "] "}, args...)...)
}

// Infof logs an information message, patterned after log.Printf.
func (l *Logger) Infof(format string, args ...interface{}) {
	sugar.Infof("I! ["+l.Name+"] "+format, args...)
}

// Warn logs a warning message, patterned after log.Print.
func (l *Logger) Warn(args ...interface{}) {
	sugar.Warn(append([]interface{}{"W! [" + l.Name + "] "}, args...)...)
}

// Warnf logs a warning message, patterned after log.Printf.
func (l *Logger) Warnf(format string, args ...interface{}) {
	sugar.Warnf("W! ["+l.Name+"] "+format, args...)
}

// Create a log name using a specific format.
func createLogName(family string, name string, alias string) string {
	switch {
	case family != "" && name != "" && alias != "":
		return family + "." + name + "." + alias
	case family != "" && name != "":
		return family + "." + name
	case family != "":
		return family
	default:
		return ""
	}
}
