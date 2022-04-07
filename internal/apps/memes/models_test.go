package memes

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestMemesModel(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Memes{})

	db.Create(&Memes{
		ID:        1,
		Title:     "test",
		Type:      MemeTypeImage,
		Status:    MemeStatusPending,
		CreatorID: 1,
		Url:       "https://www.google.com",
	})

	var memes Memes
	result := db.First(&memes, 1)
	if result.Error != nil {
		t.Errorf("failed to get memes: %v", result.Error)
	}
	assert.Equal(t, uint64(1), memes.ID)
	assert.Equal(t, "test", memes.Title)
	assert.Equal(t, MemeTypeImage, memes.Type)
	assert.Equal(t, MemeStatusPending, memes.Status)
	assert.Equal(t, uint64(1), memes.CreatorID)
	assert.Equal(t, "https://www.google.com", memes.Url)
}
