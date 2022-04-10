package memes

import (
	"github.com/bwmarrin/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
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
    like_count   int unsigned                                        NOT NULL default 0,
    unlike_count int unsigned                                        not null default 0,
    created_at   TIMESTAMP                                           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP                                           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP                                           NULL
) comment '梗图';

-- CREATE TABLE IF NOT EXISTS tags
-- (
--     id         BIGINT unsigned primary key AUTO_INCREMENT,
--     title      VARCHAR(64) NOT NULL,
--     count      int unsigned NOT NULL default 0,
--     created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP    NULL,
--     unique key (title)
-- ) comment '标签';

CREATE TABLE IF NOT EXISTS meme_tag
(
    meme_id      bigint unsigned not null ,
    tag_id       bigint unsigned not null,
    like_count   int unsigned    not null default 0,
    unlike_count int unsigned    not null default 0,
    created_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL,
    primary key (meme_id, tag_id)
) comment '梗图标签关联表';
`

func TestRepo(t *testing.T) {

	node, err := snowflake.NewNode(1)
	assert.NoError(t, err)

	db, err := sqlx.Connect("mysql", "root:q123q123@tcp(127.0.0.1:3306)/we_tools?charset=utf8mb4&parseTime=true")
	if err != nil {
		t.Fatal(err)
	}

	//_, err = db.Exec(createTableStmt)
	//if err != nil {
	//	t.Fatal(err)
	//}

	repo := NewRepo(db)

	err = repo.CreateTag("德国boy")
	if err != nil {
		t.Fatal("create tag failed", err)
	}

	tags, err := repo.GetTagsByNames([]string{"德国boy"})
	if err != nil {
		t.Fatal("get tags by names failed", err)
	}
	assert.Equal(t, 1, len(*tags))
	theTag := (*tags)[0]

	memeId := uint64(node.Generate())
	meme1 := Meme{
		ID:         memeId,
		Title:      "Test",
		StorageKey: "test",
		Type:       MemeTypeImage,
		CreatorID:  0,
		Status:     MemeStatusPending,
		LikeCount:  0,
	}

	createMemeAndTag := func(tx *sqlx.Tx) error {
		createMemeErr := repo.CreateMeme(tx, &meme1)
		if createMemeErr != nil {
			t.Fatal("create meme failed", createMemeErr)
		}
		createMemeTagErr := repo.CreateMemeTag(tx, memeId, []uint64{theTag.ID})
		if createMemeTagErr != nil {
			t.Fatal("create meme tag failed", createMemeTagErr)
		}
		return nil
	}

	err = repo.WithUnitOfWork(createMemeAndTag)
	if err != nil {
		t.Fatal("create meme and meme tag failed", err)
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
