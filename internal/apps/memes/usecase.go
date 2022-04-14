package memes

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"log"
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

type GetTagsOutputDto struct {
	Tags []TagDto
}

type Usecase interface {
	CreateMeme(input *CreateMemeInputDto) (*CreateMemeOutputDto, error)
	GetTags() (*GetTagsOutputDto, error)
	ListMemes(page int, prePage int) ([]MemeDto, error)
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

// ListMemes returns all memes
func (u *UsecaseImpl) ListMemes(page int, prePage int) ([]MemeDto, error) {
	memes, err := u.repo.ListMemes(page, prePage)
	if err != nil {
		return nil, err
	}
	memeDtoList := make([]MemeDto, 0)
	for _, meme := range *memes {
		url, err := u.storage.GetUrl(meme.StorageKey)
		if err != nil {
			log.Printf("meme error getting url. id: %v, err: %v", meme.ID, err)
			continue
		}
		memeDtoList = append(memeDtoList, MemeDto{
			ID:    strconv.FormatUint(meme.ID, 10),
			Url:   url,
			Title: meme.Title,
		})
	}
	return memeDtoList, nil
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

// GetTags returns all tags
func (u *UsecaseImpl) GetTags() (*GetTagsOutputDto, error) {
	tags, err := u.repo.GetTags()
	if err != nil {
		return nil, err
	}

	tagDtos := make([]TagDto, len(*tags))
	for i, tag := range *tags {
		tagDtos[i] = TagDto{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}
	return &GetTagsOutputDto{
		Tags: tagDtos,
	}, nil
}
