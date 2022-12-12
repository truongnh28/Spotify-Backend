package dto

type Album struct {
	AlbumID  uint   `json:"album_id"`
	Name     string `json:"name"`
	ArtistID uint   `json:"artist_id"`
	CoverImg string `json:"cover_img"`
}
type AlbumResponse struct {
	Albums []Album `json:"albums"`
	StatusError
}
