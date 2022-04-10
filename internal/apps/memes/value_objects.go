package memes

type MemeStatus string
type MemeType string

const (
	MemeStatusPending  MemeStatus = "pending"
	MemeStatusApproved MemeStatus = "approved"
	MemeStatusRejected MemeStatus = "rejected"
	MemeStatusDeleted  MemeStatus = "deleted"

	MemeTypeImage MemeType = "image"
	MemeTypeGif   MemeType = "gif"
)

type MemeDto struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Url   string   `json:"image_url"`
}

type TagDto struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
