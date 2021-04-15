package apierrors

import "fmt"

type FileIsNotExistErr struct {
	err error
}

func NewFileIsNotExistsErr(err error) *FileIsNotExistErr {
	return &FileIsNotExistErr{err: err}
}

func (e *FileIsNotExistErr) Error() string {
	return fmt.Sprintf("file is not exists: %s", e.err)
}
