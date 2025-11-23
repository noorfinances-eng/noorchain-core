package app

// AppError is a placeholder error type for NOORCHAIN.
// In Phase 2 it remains minimal until real error handling is implemented.
type AppError struct {
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

// NewError creates a new placeholder AppError instance.
func NewError(msg string) AppError {
	return AppError{Message: msg}
}
