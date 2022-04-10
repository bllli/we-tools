package memes

import (
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	//GetTags(page uint8, prePage uint8) ([]Tag, error)
	//GetMemes(tagId, page, prePage uint8) ([]Meme, error)
	GetMemeById(id uint64) (*Meme, error)

	//CreateTag(tag Tag) error
	CreateMeme(meme *Meme) error
}

type RepoImpl struct {
	db *sqlx.DB
}

var _ Repo = (*RepoImpl)(nil)

func NewRepo(db *sqlx.DB) Repo {
	return &RepoImpl{db: db}
}

//func (r *RepoImpl) GetTags(page uint8, prePage uint8) ([]Tag, error) {
//	var tags []Tag
//	err := r.db.Select(&tags, "SELECT * FROM tags LIMIT ?, ?", page, prePage)
//	return tags, err
//}

//// GetMemes returns memes by tag id
//func (r *RepoImpl) GetMemes(tagId, page, prePage uint8) ([]Meme, error) {
//	var memes []Meme
//	if tagId == 0 {
//		err := r.db.Select(&memes, "SELECT * FROM memes ORDER BY like desc LIMIT ?, ?", page, prePage)
//		return memes, err
//	} else {
//		stmt := "SELECT * FROM memes " +
//			"JOIN meme_tag ON memes.id = meme_tag.meme_id" +
//			" WHERE meme_tag.tag_id = ? ORDER BY meme_tag.like desc LIMIT ?, ?"
//		err := r.db.Select(&memes, stmt, tagId, page, prePage)
//		return memes, err
//	}
//}

// GetMemeById returns a meme by id
func (r *RepoImpl) GetMemeById(id uint64) (*Meme, error) {
	var meme Meme
	err := r.db.Get(&meme, "SELECT * FROM memes WHERE id = ?", id)
	return &meme, err
}

// CreateMeme creates a new meme
func (r *RepoImpl) CreateMeme(meme *Meme) error {
	stmt := "INSERT INTO memes (id, title, storage_key, type, creator_id, status, like_count) " +
		"VALUES (:id, :title, :storage_key, :type, :creator_id, :status, :like_count)"
	_, err := r.db.NamedExec(stmt, meme)
	return err
}
