// Package Util implements general-purpose functions
// that can be accessed from anywhere in the application.
package util

import (
	"log"
	"os"
	"time"
	"velocity-limits/config"
)

// OpenFile tries to open an input file from the given path.
func OpenFile(config *config.Configuration, path string) (*os.File, error) {
	input, err := os.Open(path + config.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

// CreateFile tries to create an output file from the given path.
func CreateFile(config *config.Configuration, path string) *os.File {
	output, err := os.Create(path + config.OutputFile)
	if err != nil {
		log.Fatal("Error - Unable to open file: ", err)
	}
	return output
}

// Returns beginning of the day in UTC format.
func GetBeginningOfTheDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
}

// Returns beginning of the week in UTC format. Week starts from Monday.
func GetBeginningOfTheWeek(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()+int(time.Monday-d.Weekday()), 0, 0, 0, 0, time.UTC)
}
