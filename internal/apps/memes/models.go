package memes

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255);unique_index"`
}

type Memes struct {
	gorm.Model
	ID        uint64     `gorm:"primary_key"`
	Title     string     `gorm:"type:varchar(64);not null"`
	Url       string     `gorm:"type:varchar(128);not null"`
	Type      MemeType   `gorm:"not null"`
	CreatorID uint64     `gorm:"not null"`
	Status    MemeStatus `gorm:"not null"`
}

type MemesTags struct {
	gorm.Model
	ID      uint64 `gorm:"primary_key"`
	MemesID uint64 `gorm:"not null"`
	TagsID  uint64 `gorm:"not null"`
}
