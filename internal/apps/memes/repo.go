package memes

import (
	"github.com/jmoiron/sqlx"
	"we-tools/internal/common/persistence"
)

type Repo interface {
	WithUnitOfWork(f persistence.UnitOfWorkFunc) error

	GetTags() (*[]Tag, error)
	GetTagsByNames(names []string) (*[]Tag, error)
	GetMemes(tagId, page, prePage uint8) (*[]Meme, error)
	GetMemeById(id uint64) (*Meme, error)

	CreateTag(name string) error
	CreateMeme(tx *sqlx.Tx, meme *Meme) error
	CreateMemeTag(tx *sqlx.Tx, memeId uint64, tagIds []uint64) error
}

type RepoImpl struct {
	db  *sqlx.DB
	uow *persistence.UnitOfWork
}

var _ Repo = (*RepoImpl)(nil)

func NewRepo(db *sqlx.DB) Repo {
	uow := persistence.NewUnitOfWork(db)
	return &RepoImpl{
		db:  db,
		uow: uow,
	}
}

// WithUnitOfWork executes the given function in a UnitOfWork.
func (r *RepoImpl) WithUnitOfWork(f persistence.UnitOfWorkFunc) error {
	return r.uow.Execute(f)
}

func (r *RepoImpl) GetTags() (*[]Tag, error) {
	var tags []Tag
	err := r.db.Select(&tags, "SELECT * FROM tags where deleted_at is null")
	if err != nil {
		return nil, err
	}
	return &tags, err
}

// GetTagsByNames returns tags by names
func (r *RepoImpl) GetTagsByNames(names []string) (*[]Tag, error) {
	var tags []Tag
	query, args, err := sqlx.In("SELECT * FROM tags WHERE name IN (?)", names)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&tags, query, args...)
	if err != nil {
		return nil, err
	}
	return &tags, err
}

// GetMemes returns memes by tag id
func (r *RepoImpl) GetMemes(tagId, page, prePage uint8) (*[]Meme, error) {
	var memes []Meme
	if tagId == 0 {
		err := r.db.Select(&memes, "SELECT * FROM memes ORDER BY like desc LIMIT ?, ?", page, prePage)
		return &memes, err
	} else {
		stmt := "SELECT * FROM memes " +
			"JOIN meme_tag ON memes.id = meme_tag.meme_id" +
			" WHERE meme_tag.tag_id = ? ORDER BY meme_tag.like desc LIMIT ?, ?"
		err := r.db.Select(&memes, stmt, tagId, page, prePage)
		return &memes, err
	}
}

// GetMemeById returns a meme by id
func (r *RepoImpl) GetMemeById(id uint64) (*Meme, error) {
	var meme Meme
	err := r.db.Get(&meme, "SELECT * FROM memes WHERE id = ?", id)
	return &meme, err
}

// CreateMeme creates a new meme
func (r *RepoImpl) CreateMeme(tx *sqlx.Tx, meme *Meme) error {
	stmt := "INSERT INTO memes (id, title, storage_key, type, creator_id, status, like_count) " +
		"VALUES (:id, :title, :storage_key, :type, :creator_id, :status, :like_count)"
	_, err := tx.NamedExec(stmt, meme)
	return err
}

// CreateTag creates a new tag
func (r *RepoImpl) CreateTag(name string) error {
	stmt := "INSERT INTO tags (name) VALUES (?) ON DUPLICATE KEY UPDATE deleted_at = NULL"
	_, err := r.db.Exec(stmt, name)
	return err
}

// CreateMemeTag creates a new meme tag
func (r *RepoImpl) CreateMemeTag(tx *sqlx.Tx, memeId uint64, tagIds []uint64) error {
	for _, tagId := range tagIds {
		stmt := "INSERT INTO meme_tag (meme_id, tag_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE deleted_at = NULL"
		_, err := tx.Exec(stmt, memeId, tagId)
		if err != nil {
			return err
		}
	}
	return nil
}
