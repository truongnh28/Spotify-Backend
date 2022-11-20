package dto

import "time"

type Song struct {
	SongID      uint      `json:"song_id"`
	Name        string    `json:"name"`
	AlbumID     uint      `json:"album_id"`
	ArtistID    uint      `json:"artist_id"`
	Lyrics      string    `json:"lyrics"`
	Length      time.Time `json:"length"`
	TrackNumber uint      `json:"track_number"`
	CreateAt    time.Time `json:"create_at"`
	UploadAt    time.Time `json:"upload_at"`
	YoutubeLink string    `json:"youtube_link"`
}

type SongResponse struct {
	Songs []Song
	StatusError
}
