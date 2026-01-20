package storage

import (
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const (
	FILE_STORAGE_PATH = "media/"
)

type FileStorage interface {
	// 파일 스트림, 메타데이터
	Save(file []byte, metadata map[string]string) (string, error)
	Delete(path string) error
}

type fileStorage struct {
}

func NewFileStorage() FileStorage {
	return &fileStorage{}
}

func (s *fileStorage) Save(file []byte, metadata map[string]string) (string, error) {
	now := time.Now()
	dir := filepath.Join(
		FILE_STORAGE_PATH,
		now.Format("2006"),
		now.Format("01"),
		now.Format("02"),
	)

	ext := filepath.Ext(metadata["filename"])
	filename := uuid.New().String() + ext
	fullPath := filepath.Join(dir, filename)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	if err := os.WriteFile(fullPath, file, 0644); err != nil {
		return "", err
	}

	return fullPath, nil
}

func (s *fileStorage) Delete(path string) error {
	if path == "" {
		return nil
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
