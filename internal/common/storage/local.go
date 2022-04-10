package storage

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"we-tools/internal/common"
)

type LocalStorage struct {
	path      string
	urlPrefix string
}

func NewLocalStorage(path string, urlPrefix string) (*LocalStorage, error) {
	if path == "" {
		return nil, errors.New("path is empty")
	}
	if urlPrefix == "" {
		return nil, errors.New("urlPrefix is empty")
	}
	if !strings.HasSuffix(urlPrefix, "/") {
		urlPrefix += "/"
	}
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}
	return &LocalStorage{path: path, urlPrefix: urlPrefix}, nil
}

var _ Storage = (*LocalStorage)(nil) // LocalStorage implements Storage interface

func (l *LocalStorage) Exists(key string) bool {
	if key == "" {
		return false
	}
	_, err := os.Stat(l.path + "/" + key)
	return err == nil
}

func (l *LocalStorage) GetUrl(key string) (string, error) {
	if !l.Exists(key) {
		return "", ErrNotFound
	}
	return l.urlPrefix + key, nil
}

func (l *LocalStorage) Upload(key string, data []byte) (string, error) {

	if strings.Contains(key, "/") {
		err := os.MkdirAll(l.path+"/"+strings.Split(key, "/")[0], 0755)
		if err != nil {
			return "", err
		}
	}

	err := os.WriteFile(l.path+"/"+key, data, 0644)
	if err != nil {
		return "", err
	}
	url, err := l.GetUrl(key)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (l *LocalStorage) Get(key string) ([]byte, error) {
	return os.ReadFile(l.path + "/" + key)
}

func (l *LocalStorage) Handle(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		common.Fail(c, "key is empty")
		return
	}
	if !l.Exists(key) {
		common.Fail(c, "key not found")
		return
	}
	data, err := l.Get(key)
	if err != nil {
		common.Fail(c, "get file error")
		return
	}
	c.Data(200, "image/jpeg", data)
}
