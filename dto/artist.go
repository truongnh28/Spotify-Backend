package dto

type Artist struct {
	ArtistID    uint   `json:"artist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CoverImg    string `json:"coverImg"`
}
type ArtistResponse struct {
	Artists []Artist `json:"artists"`
	StatusError
}
