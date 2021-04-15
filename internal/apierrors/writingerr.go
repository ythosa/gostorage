package apierrors

import "fmt"

type WritingErr struct {
	err error
}

func NewWritingErr(err error) *WritingErr {
	return &WritingErr{err: err}
}

func (e *WritingErr) Error() string {
	return fmt.Sprintf("error occurred while writing into file: %s", e.err)
}
