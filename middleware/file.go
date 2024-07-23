package middleware

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const uploadPath = "./uploads"

func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", err
	}

	filename := filepath.Base(file.Filename)
	if ext := filepath.Ext(filename); ext != ".pdf" && ext != ".docx" {
		return "", errors.New("file format not allowed")
	}

	filePath := filepath.Join(uploadPath, filename)

	if err := SaveFile(file, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

func SaveFile(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
