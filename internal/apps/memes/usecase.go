package memes

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"strconv"
)
import "we-tools/internal/common/storage"

type CreateMemeInputDto struct {
	Title          string
	FileContent    []byte
	Type           MemeType
	FilenameSuffix string
}

type CreateMemeOutputDto struct {
	ID  uint64
	Url string
}

type Usecase interface {
	CreateMeme(input *CreateMemeInputDto) (*CreateMemeOutputDto, error)
}

type UsecaseImpl struct {
	repo        Repo
	idGenerator *snowflake.Node
	storage     storage.Storage
}

var _ Usecase = (*UsecaseImpl)(nil)

// NewUsecase creates a new Usecase
func NewUsecase(repo Repo, idGenerator *snowflake.Node, storage storage.Storage) Usecase {
	return &UsecaseImpl{
		repo:        repo,
		idGenerator: idGenerator,
		storage:     storage,
	}
}

// CreateMeme creates a new meme
func (u *UsecaseImpl) CreateMeme(input *CreateMemeInputDto) (*CreateMemeOutputDto, error) {
	id := uint64(u.idGenerator.Generate())
	key := "meme/" + strconv.FormatUint(id, 10) + "." + input.FilenameSuffix
	url, err := u.storage.Upload(key, input.FileContent)
	if err != nil {
		return nil, err
	}
	meme := Meme{
		ID:         id,
		Title:      input.Title,
		StorageKey: key,
		Type:       input.Type,
		CreatorID:  0,
		Status:     MemeStatusPending,
		LikeCount:  0,
	}

	err = u.repo.WithUnitOfWork(func(tx *sqlx.Tx) error {
		err := u.repo.CreateMeme(tx, &meme)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &CreateMemeOutputDto{
		ID:  id,
		Url: url,
	}, nil
}
