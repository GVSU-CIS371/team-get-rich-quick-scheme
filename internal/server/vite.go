package server

import (
	"errors"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func getTargetFile(base, filePath string) (*os.File, error) {
	targetFile, err := filepath.Abs(path.Join(base, filePath))
	if err != nil {
		return nil, err
	}
	absBase, _ := filepath.Abs(base)
	if strings.HasPrefix(targetFile, absBase) {
		info, err := os.Stat(targetFile)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			return nil, errors.New("target file is a directory")
		}
		return os.OpenFile(targetFile, os.O_RDONLY, 0666)
	}
	return nil, errors.New("path traversal error")
}

func getFileContentType(extension string, data []byte) string {
	cType := mime.TypeByExtension(extension)
	if cType != "" {
		return cType
	}
	return http.DetectContentType(data)
}
