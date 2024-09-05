package storage

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"path"
)

const (
	DocumentsFile string = ".cliga_cache.json"
	TokenFile     string = ".cliga_token.json"

	DocsStr  string = "docs"
	TokenStr string = "token"
)

type CacheData map[string]interface{}
type TokenData map[string]string

type FileStorage interface {
	Read(key string) ([]byte, error)
	Write(key string, data []byte) error
}

type TmpFileStorage struct {
	tmpDocumentsFile string
	tmpTokenFile     string
}

func NewFileStorage() FileStorage {
	tmpDir := os.TempDir()
	return &TmpFileStorage{
		tmpDocumentsFile: path.Join(tmpDir, DocumentsFile),
		tmpTokenFile:     path.Join(tmpDir, TokenFile),
	}
}

func ReadFile(filepath string) ([]byte, error) {
	var emptyBinary []byte

	log := slog.With("file", filepath)

	if _, err := os.Stat(filepath); err != nil {
		log.Debug("no file found")
		return emptyBinary, err
	}
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0660)
	if err != nil {
		return emptyBinary, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Debug("error get file stat")
		return emptyBinary, err
	}

	if fileInfo.Size() == 0 {
		log.Debug("size file is 0")
		return emptyBinary, err
	}

	dataBinary, err := io.ReadAll(file)
	if err != nil {
		log.Debug("error on read data")
		return emptyBinary, err
	}
	return dataBinary, nil
}

// Rewrite tmp file
func WriteFile(filepath string, data []byte) error {
	log := slog.With("file", filepath)

	file, err := os.OpenFile(filepath, os.O_WRONLY, 0660)
	if err != nil {
		log.Debug("error: open file")
		return err
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		log.Debug("error: write file")
		return err
	}
	return nil
}

func (s *TmpFileStorage) Read(key string) ([]byte, error) {
	switch key {
	case DocsStr:
		ReadFile(s.tmpDocumentsFile)
	case TokenStr:
		ReadFile(s.tmpTokenFile)
	}
	return []byte{}, errors.New("not valid key")
}

func (s *TmpFileStorage) Write(key string, data []byte) error {
	switch key {
	case DocsStr:
		WriteFile(s.tmpDocumentsFile, data)
	case TokenStr:
		WriteFile(s.tmpTokenFile, data)
	}
	return errors.New("not valid key")
}

func (s *TmpFileStorage) CheckCacheExist(key string) bool {
	switch key {
	case DocsStr:
		return checkFileExist(s.tmpDocumentsFile)
	case TokenStr:
		return checkFileExist(s.tmpTokenFile)
	}
	return false
}

func checkFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
