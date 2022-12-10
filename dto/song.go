package dto

type Song struct {
	SongID      uint   `json:"song_id"`
	Name        string `json:"name"`
	AlbumID     uint   `json:"album_id"`
	ArtistID    uint   `json:"artist_id"`
	Lyrics      string `json:"lyrics"`
	Length      uint   `json:"length"`
	TrackNumber uint   `json:"track_number"`
	YoutubeLink string `json:"youtube_link"`
}

type SongResponse struct {
	Songs []Song
	StatusError
}
