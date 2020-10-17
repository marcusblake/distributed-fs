package internal

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
