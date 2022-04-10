package memes

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var createTableStmt = `
CREATE TABLE IF NOT EXISTS memes
(
    id           BIGINT unsigned primary key,
    title        VARCHAR(254)                                        NOT NULL,
    storage_key  VARCHAR(254)                                        NOT NULL,
    type         ENUM ('image', 'gif')                               NOT NULL,
    creator_id   BIGINT                                              NOT NULL,
    status       ENUM ('pending', 'approved', 'rejected', 'deleted') NOT NULL,
    like_count   BIGINT                                              NOT NULL default 0
);
`

func TestRepo(t *testing.T) {

	node, err := snowflake.NewNode(1)
	assert.NoError(t, err)

	db, err := sqlx.Connect("mysql", "root:q123q123@tcp(127.0.0.1:3306)/we_tools")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(createTableStmt)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewRepo(db)

	meme1 := Meme{
		ID:         uint64(node.Generate().Int64()),
		Title:      "Test",
		StorageKey: "test",
		Type:       MemeTypeImage,
		CreatorID:  0,
		Status:     MemeStatusPending,
		LikeCount:  0,
	}

	err = repo.CreateMeme(&meme1)
	if err != nil {
		t.Fatal(err)
	}

	meme2, err := repo.GetMemeById(meme1.ID)
	if err != nil {
		return
	}
	assert.Equal(t, meme1.ID, meme2.ID)
	assert.Equal(t, meme1.Title, meme2.Title)
	assert.Equal(t, meme1.StorageKey, meme2.StorageKey)
	assert.Equal(t, meme1.Type, meme2.Type)
	assert.Equal(t, meme1.CreatorID, meme2.CreatorID)
	assert.Equal(t, meme1.Status, meme2.Status)
	assert.Equal(t, meme1.LikeCount, meme2.LikeCount)
}
