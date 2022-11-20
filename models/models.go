package models

import (
	"gorm.io/gorm"
	"time"
)

type Albums struct {
	*gorm.Model
	AlbumID  uint      `gorm:"column:album_ID"`
	Name     string    `gorm:"column:name"`
	ArtistID uint      `gorm:"column:artist_ID"`
	Artists  Artists   `gorm:"foreignKey:ArtistID;references:ID"`
	CreateAt time.Time `gorm:"column:create_at"`
	UploadAt time.Time `gorm:"column:upload_at"`
	CoverImg string    `gorm:"column:cover_img"`
}

type Artists struct {
	*gorm.Model
	ArtistID    uint      `gorm:"column:artist_ID"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UploadAt    time.Time `gorm:"column:upload_at"`
	CoverImg    string    `gorm:"column:cover_img"`
}

type Interactions struct {
	*gorm.Model
	InteractionID uint      `gorm:"column:interaction_ID"`
	UserID        uint      `gorm:"column:user_ID"`
	Users         Users     `gorm:"foreignKey:UserID;references:ID"`
	SongID        uint      `gorm:"column:song_ID"`
	Songs         Songs     `gorm:"foreignKey:SongID;references:ID"`
	Liked         uint      `gorm:"column:liked"`
	PlayDuration  time.Time `gorm:"column:play_duration"`
	CreateAt      time.Time `gorm:"column:create_at"`
	UploadAt      time.Time `gorm:"column:upload_at"`
}

type PlayLists struct {
	*gorm.Model
	PlayListID uint      `gorm:"column:playlist_ID"`
	Name       string    `gorm:"column:name"`
	CreateAt   time.Time `gorm:"column:create_at"`
	UploadAt   time.Time `gorm:"column:upload_at"`
	CoverImg   string    `gorm:"column:cover_img"`
	UserID     uint      `gorm:"column:user_ID"`
	Users      Users     `gorm:"foreignKey:UserID;references:ID"`
}

type PlayListSongs struct {
	SongID     uint      `gorm:"column:song_ID"`
	Songs      Songs     `gorm:"foreignKey:SongID;references:ID"`
	PlayListID uint      `gorm:"column:playlist_ID"`
	PlayLists  PlayLists `gorm:"foreignKey:PlayListID;references:ID"`
}

type Songs struct {
	*gorm.Model
	SongID      uint      `gorm:"column:song_ID"`
	Name        string    `gorm:"column:name"`
	AlbumID     uint      `gorm:"column:album_ID"`
	Albums      Albums    `gorm:"foreignKey:AlbumID;references:ID"`
	ArtistID    uint      `gorm:"column:artist_ID"`
	Artists     Artists   `gorm:"foreignKey:ArtistID;references:ID"`
	Lyrics      string    `gorm:"column:lyrics"`
	Length      time.Time `gorm:"column:length"`
	TrackNumber uint      `gorm:"column:track_number"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UploadAt    time.Time `gorm:"column:upload_at"`
	YoutubeLink string    `gorm:"column:youtube_link"`
}

type Users struct {
	*gorm.Model
	UserID        uint      `gorm:"column:user_ID"`
	Name          string    `gorm:"column:name"`
	Email         string    `gorm:"column:email"`
	Password      string    `gorm:"column:password"`
	IsAdmin       uint      `gorm:"column:is_admin"`
	Preferences   string    `gorm:"column:preferences"`
	RememberToken string    `gorm:"column:remember_token"`
	CreateAt      time.Time `gorm:"column:create_at"`
	UploadAt      time.Time `gorm:"column:upload_at"`
	Role          string    `gorm:"column:role"`
	GroupType     uint      `gorm:"column:group_type"`
	Status        string    `gorm:"column:status"`
}

// Status enum
type Status byte

const (
	InitialStatus = iota
	WaitingForScanStatus
	ScanningStatus
	ScannedStatus
)

type Table interface {
	TableName() string
}

func (Albums) TableName() string {
	return "albums"
}

func (Artists) TableName() string {
	return "artists"
}

func (Interactions) TableName() string {
	return "interactions"
}

func (PlayLists) TableName() string {
	return "playlists"
}

func (PlayListSongs) TableName() string {
	return "playlist_songs"
}

func (Songs) TableName() string {
	return "songs"
}

func (Users) TableName() string {
	return "users"
}
