package apierrors

import "fmt"

type FileAlreadyExistsErr struct {
	err error
}

func NewFileAlreadyExistsErr(err error) *FileAlreadyExistsErr {
	return &FileAlreadyExistsErr{err: err}
}

func (e *FileAlreadyExistsErr) Error() string {
	return fmt.Sprintf("file already exists: %s", e.err)
}
