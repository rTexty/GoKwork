package errors

import "fmt"

// KworkError представляет общую ошибку Kwork API
type KworkError struct {
	Message string
}

func (e *KworkError) Error() string {
	return fmt.Sprintf("kwork error: %s", e.Message)
}

// NewKworkError создает новую ошибку Kwork
func NewKworkError(message string) *KworkError {
	return &KworkError{Message: message}
}

// KworkBotError представляет ошибку Kwork Bot
type KworkBotError struct {
	Message string
}

func (e *KworkBotError) Error() string {
	return fmt.Sprintf("kwork bot error: %s", e.Message)
}

// NewKworkBotError создает новую ошибку Kwork Bot
func NewKworkBotError(message string) *KworkBotError {
	return &KworkBotError{Message: message}
}
