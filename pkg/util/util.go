package util

import (
	"log"
	"os"
	"time"
	"velocity-limits/config"
)

func OpenFile(config *config.Configuration, path string) (*os.File, error) {
	input, err := os.Open(path + config.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func CreateFile(config *config.Configuration, path string) *os.File {
	output, err := os.Create(path + config.OutputFile)
	if err != nil {
		log.Fatal("Error - Unable to open file: ", err)
	}
	return output
}

func GetBeginningOfTheDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
}

func GetBeginningOfTheWeek(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()+int(time.Monday-d.Weekday()), 0, 0, 0, 0, time.UTC)
}
