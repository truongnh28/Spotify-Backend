package dto

type Song struct {
	SongID      uint   `json:"song_id"`
	Name        string `json:"name"`
	AlbumID     uint   `json:"album_id"`
	ArtistID    uint   `json:"artist_id"`
	Lyrics      string `json:"lyrics"`
	Length      uint   `json:"length"`
	URL         string `json:"url"`
	YoutubeLink string `json:"youtube_link"`
	SongCloudId string `json:"song_cloud_id"`
}

type SongResponse struct {
	Songs []Song `json:"songs"`
	StatusError
}
