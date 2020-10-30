package logger

import (
	"github.com/fatih/color"
)

// Warning prints a warning to the screen
func Warning(message string) {
	color.Yellow(message)
}

// Failure prints a failure to the screen
func Failure(message string) {
	color.Red(message)
}

// Success prints a success to the screen
func Success(message string) {
	color.Green(message)
}

// Info prints information to the screen
func Info(message string) {
	color.Cyan(message)
}

// Warningf prints a warning to the screen
func Warningf(format string, a ...interface{}) {
	printer := color.New(color.FgYellow)
	printer.Printf(format, a...)
}

// Failuref prints a failure to the screen
func Failuref(format string, a ...interface{}) {
	printer := color.New(color.FgRed)
	printer.Printf(format, a...)
}

// Successf prints a success to the screen
func Successf(format string, a ...interface{}) {
	printer := color.New(color.FgGreen)
	printer.Printf(format, a...)
}

// Infof prints information to the screen
func Infof(format string, a ...interface{}) {
	printer := color.New(color.FgCyan)
	printer.Printf(format, a...)
}
