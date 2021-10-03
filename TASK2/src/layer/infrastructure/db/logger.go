package db

import "fmt"

type LoggerRepository interface {
	Log(string)
}

type loggerRepository struct{}

func NewLoggerRepository() LoggerRepository {
	return &loggerRepository{}
}

var (
	mysqlClient = "Mock mysqlClient instance"
)

func (r *loggerRepository) Log(message string) {
	// TODO: save message to database
	fmt.Println("Save message log: " + message)
}
