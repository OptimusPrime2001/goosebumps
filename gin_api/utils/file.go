package utils

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func CheckFileSize(file *multipart.FileHeader, maxSize int64) error {
	if file.Size > maxSize {
		return errors.New("file size exceeds the maximum allowed size")
	}
	return nil
}

var allowedExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func CheckFileName(file *multipart.FileHeader) error {
	fileType := strings.ToLower(filepath.Ext(file.Filename))
	if _, ok := allowedExts[fileType]; ok {
		return nil
	}
	return errors.New("file type not allowed")
}
func CheckFileType(file *multipart.FileHeader) error {
	return CheckFileName(file)
}
