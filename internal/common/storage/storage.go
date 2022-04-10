package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	Upload(key string, data []byte) (string, error)
	//Delete(key string) error
	GetUrl(key string) (string, error)
	Get(key string) ([]byte, error)
}
