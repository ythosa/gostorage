package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type FileCredentials struct {
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

func (m *FileCredentials) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Name, validation.Required, is.Alphanumeric),
		validation.Field(&m.Ext, validation.Required, is.Alphanumeric),
	)
}
