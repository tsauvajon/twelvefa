package main

// ErrIntOverflow : integer overflow or underflow
type ErrIntOverflow struct {
	message string
}

func newErrIntOverflow(message string) *ErrIntOverflow {
	return &ErrIntOverflow{
		message: message,
	}
}

// Error formats the error as a string
func (e *ErrIntOverflow) Error() string {
	return e.message
}

// ErrZeroDivision : trying to divide by 0
type ErrZeroDivision struct {
	message string
}

func newErrZeroDivision(message string) *ErrZeroDivision {
	return &ErrZeroDivision{
		message: message,
	}
}

// Error formats the error as a string
func (e *ErrZeroDivision) Error() string {
	return e.message
}
