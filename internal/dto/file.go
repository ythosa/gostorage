package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type File struct {
	Name string `json:"name"`
	Ext  string `json:"ext"`
	Data string `json:"data"`
}

func (m *File) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Data, validation.Required),
		validation.Field(&m.Name, validation.Required, is.Alphanumeric),
		validation.Field(&m.Ext, validation.Required, is.Alphanumeric),
	)
}
