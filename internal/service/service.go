package service

import (
	"github.com/Ythosa/gostorage/internal/dto"
	"github.com/sirupsen/logrus"
)

type File interface {
	UploadFile(payload dto.File) (dto.FileURL, error)
	UpdateFile(payload dto.File) (dto.FileURL, error)
	DeleteFile(payload dto.FileCredentials) error
}

type Service struct {
	File
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{
		File: NewFileService(logger),
	}
}
