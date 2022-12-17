package dto

type Interaction struct {
	UserID uint `json:"user_id"`
	SongID uint `json:"song_id"`
	Liked  uint `json:"liked"`
}
type InteractionResponse struct {
	Interactions []Interaction `json:"interactions"`
	StatusError
}
