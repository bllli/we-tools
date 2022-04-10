package memes

//type Tag struct {
//	ID    uint64 `db:"id"`
//	Name  string `db:"name"`
//	Count uint64 `db:"count"`
//}

type Meme struct {
	ID         uint64     `db:"id"`
	Title      string     `db:"title"`
	StorageKey string     `db:"storage_key"`
	Type       MemeType   `db:"type"`
	CreatorID  uint64     `db:"creator_id"`
	Status     MemeStatus `db:"status"`
	LikeCount  uint64     `db:"like_count"`
}

//
//type MemeTag struct {
//	ID      uint64 `db:"id"`
//	MemesID uint64 `db:"memes_id"`
//	TagsID  uint64 `db:"tags_id"`
//	Like    uint64 `db:"like"`
//}
