package memes

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID        uint64       `db:"id"`
	Name      string       `db:"name"`
	Count     uint64       `db:"count"`
	CreatedAt time.Time    `db:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type Meme struct {
	ID          uint64       `db:"id"`
	Title       string       `db:"title"`
	StorageKey  string       `db:"storage_key"`
	Type        MemeType     `db:"type"`
	CreatorID   uint64       `db:"creator_id"`
	Status      MemeStatus   `db:"status"`
	LikeCount   uint64       `db:"like_count"`
	UnLikeCount uint64       `db:"unlike_count"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type MemeTag struct {
	MemeID      uint64       `db:"meme_id"`
	TagID       uint64       `db:"tag_id"`
	LikeCount   uint64       `db:"like_count"`
	UnLikeCount uint64       `db:"unlike_count"`
	CreatedAt   time.Time    `db:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}
