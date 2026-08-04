package databasepackage database

import "fmt"

type DatabaseError struct {
	Code    string
	Message string
	Err     error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
}

var (
	ErrDatabaseConnection = &DatabaseError{
		Code:    "DB_CONNECTION",
		Message: "Erro ao conectar no banco",
		Err:     nil,
	}
)

func NewDatabaseError(code string, message string, err error) *DatabaseError {
	return &DatabaseError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}