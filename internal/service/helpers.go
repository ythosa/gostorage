package service

import "github.com/Ythosa/gostorage/internal/config"

func generateHref(filePath string) string {
	return config.Get().BindHost + config.Get().BindAddr + filePath
}
