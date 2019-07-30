package calc

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

// ErrOutOfRange : input is not in the defined range
type ErrOutOfRange struct {
	message string
}

func newErrOutOfRange(message string) *ErrOutOfRange {
	return &ErrOutOfRange{
		message: message,
	}
}

// Error formats the error as a string
func (e *ErrOutOfRange) Error() string {
	return e.message
}
