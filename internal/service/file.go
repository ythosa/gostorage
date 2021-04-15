package service

import (
	"fmt"
	"github.com/Ythosa/gostorage/internal/apierrors"
	"github.com/Ythosa/gostorage/internal/config"
	"github.com/Ythosa/gostorage/internal/dto"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type FileService struct {
	logger *logrus.Logger
}

func NewFileService(logger *logrus.Logger) *FileService {
	return &FileService{
		logger: logger,
	}
}

func (s *FileService) UploadFile(payload dto.File) (dto.FileURL, error) {
	filePath := fmt.Sprintf("%s%s.%s", config.Get().Files.Path, payload.Name, payload.Ext)

	if _, err := os.Stat(filePath); err == nil {
		return dto.FileURL{}, apierrors.NewFileAlreadyExistsErr(fmt.Errorf("path: %s", filePath))
	}

	err := ioutil.WriteFile(fmt.Sprintf(".%s", filePath), []byte(payload.Data), 0600)
	if err != nil {
		return dto.FileURL{}, apierrors.NewWritingErr(err)
	}

	return dto.FileURL{URL: generateHref(filePath)}, nil
}

func (s *FileService) UpdateFile(payload dto.File) (dto.FileURL, error) {
	filePath := fmt.Sprintf("%s%s.%s", config.Get().Files.Path, payload.Name, payload.Ext)
	fp := fmt.Sprintf(".%s", filePath)

	if _, err := os.Stat(fp); err != nil && os.IsNotExist(err) {
		return dto.FileURL{}, apierrors.NewFileIsNotExistsErr(err)
	}

	err := ioutil.WriteFile(fp, []byte(payload.Data), 0600)
	if err != nil {
		return dto.FileURL{}, apierrors.NewWritingErr(err)
	}

	return dto.FileURL{URL: generateHref(filePath)}, nil
}

func (s *FileService) DeleteFile(payload dto.FileCredentials) error {
	fp := fmt.Sprintf(".%s%s.%s", config.Get().Files.Path, payload.Name, payload.Ext)

	if err := os.Remove(fp); err != nil {
		return apierrors.NewFileIsNotExistsErr(err)
	}

	return nil
}
