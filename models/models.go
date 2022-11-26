package models

import (
	"gorm.io/gorm"
	"time"
)

type Albums struct {
	*gorm.Model
	AlbumID   uint      `gorm:"column:album_ID"`
	Name      string    `gorm:"column:name"`
	ArtistID  uint      `gorm:"column:artist_ID"`
	Artists   Artists   `gorm:"foreignKey:ArtistID;references:ID"`
	CreateAt  time.Time `gorm:"column:create_at"`
	UploadAt  time.Time `gorm:"column:upload_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	CoverImg  string    `gorm:"column:cover_img"`
}

type Artists struct {
	*gorm.Model
	ArtistID    uint      `gorm:"column:artist_ID"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UploadAt    time.Time `gorm:"column:upload_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
	CoverImg    string    `gorm:"column:cover_img"`
}

type Interactions struct {
	*gorm.Model
	InteractionID uint      `gorm:"column:interaction_ID"`
	UserID        uint      `gorm:"column:user_ID"`
	Users         Accounts  `gorm:"foreignKey:Id;references:ID"`
	SongID        uint      `gorm:"column:song_ID"`
	Songs         Songs     `gorm:"foreignKey:SongID;references:ID"`
	Liked         uint      `gorm:"column:liked"`
	PlayDuration  time.Time `gorm:"column:play_duration"`
	CreateAt      time.Time `gorm:"column:create_at"`
	UploadAt      time.Time `gorm:"column:upload_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
}

type PlayLists struct {
	*gorm.Model
	Id        uint      `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreateAt  time.Time `gorm:"column:create_at"`
	UploadAt  time.Time `gorm:"column:upload_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	CoverImg  string    `gorm:"column:cover_img"`
	UserID    uint      `gorm:"column:user_ID"`
	Users     Accounts  `gorm:"foreignKey:Id;references:ID"`
}

type PlayListSongs struct {
	*gorm.Model
	SongID     uint      `gorm:"column:song_ID"`
	Songs      Songs     `gorm:"foreignKey:SongID;references:ID"`
	PlayListID uint      `gorm:"column:playlist_ID"`
	PlayLists  PlayLists `gorm:"foreignKey:Id;references:ID"`
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
	Length      uint      `gorm:"column:length"`
	TrackNumber uint      `gorm:"column:track_number"`
	CreateAt    time.Time `gorm:"column:create_at"`
	UploadAt    time.Time `gorm:"column:upload_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
	YoutubeLink string    `gorm:"column:youtube_link"`
}

type Accounts struct {
	*gorm.Model
	Id        uint          `gorm:"column:id"`
	UserName  string        `gorm:"index:username_idx_uni,unique"`
	Email     string        `gorm:"column:email"`
	Password  string        `gorm:"column:password"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
	DeletedAt time.Time     `gorm:"column:deleted_at"`
	Role      AccountRole   `gorm:"column:role"`
	GroupType uint          `gorm:"column:group_type"`
	Status    AccountStatus `gorm:"column:status"`
}

type AccountStatus string

const (
	Active  AccountStatus = "Active"
	Blocked AccountStatus = "Blocked"
)

type AccountRole string

const (
	SuperAdmin AccountRole = "SuperAdmin"
	Admin      AccountRole = "Admin"
	User       AccountRole = "User"
)

var (
	priority = map[AccountRole]int{
		SuperAdmin: 0,
		Admin:      -1,
		User:       -2,
	}
)

func (a AccountRole) CanChange(ar AccountRole) bool {
	return priority[a] > priority[ar]
}

func (a AccountRole) IsValid() bool {
	_, ok := priority[a]
	return ok
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

func (Accounts) TableName() string {
	return "accounts"
}
