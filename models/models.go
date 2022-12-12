package models

import (
	"gorm.io/gorm"
	"time"
)

type Album struct {
	*gorm.Model
	AlbumID    uint      `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	ArtistID   uint      `gorm:"column:artist_ID"`
	Artists    Artist    `gorm:"foreignKey:id;references:id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UploadedAt time.Time `gorm:"column:uploaded_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
	CoverImg   string    `gorm:"column:cover_img"`
}

type Artist struct {
	*gorm.Model
	ArtistID    uint      `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UploadedAt  time.Time `gorm:"column:uploaded_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
	CoverImg    string    `gorm:"column:cover_img"`
}

type Interaction struct {
	*gorm.Model
	InteractionID uint      `gorm:"column:id"`
	UserID        uint      `gorm:"column:user_ID"`
	Users         Account   `gorm:"foreignKey:id;references:id"`
	SongID        uint      `gorm:"column:song_ID"`
	Songs         Song      `gorm:"foreignKey:id;references:id"`
	Liked         uint      `gorm:"column:liked"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UploadedAt    time.Time `gorm:"column:uploaded_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at"`
}

type PlayList struct {
	*gorm.Model
	PlayListID uint      `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UploadedAt time.Time `gorm:"column:uploaded_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
	CoverImg   string    `gorm:"column:cover_img"`
	UserID     uint      `gorm:"column:user_ID"`
	Users      Account   `gorm:"foreignKey:id;references:id"`
}

type PlayListSong struct {
	*gorm.Model
	SongID     uint     `gorm:"column:song_ID"`
	Songs      Song     `gorm:"foreignKey:id;references:id"`
	PlayListID uint     `gorm:"column:playlist_ID"`
	PlayLists  PlayList `gorm:"foreignKey:id;references:id"`
}

type Song struct {
	*gorm.Model
	SongID      uint      `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	AlbumID     uint      `gorm:"column:album_ID"`
	Albums      Album     `gorm:"foreignKey:id;references:id"`
	ArtistID    uint      `gorm:"column:artist_ID"`
	Artists     Artist    `gorm:"foreignKey:id;references:id"`
	Lyrics      string    `gorm:"column:lyrics"`
	Length      uint      `gorm:"column:length"`
	URL         string    `gorm:"column:url"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UploadedAt  time.Time `gorm:"column:uploaded_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
	YoutubeLink string    `gorm:"column:youtube_link"`
}

type Account struct {
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

type Table interface {
	TableName() string
}

func (Album) TableName() string {
	return "albums"
}

func (Artist) TableName() string {
	return "artists"
}

func (Interaction) TableName() string {
	return "interactions"
}

func (PlayList) TableName() string {
	return "playlists"
}

func (PlayListSong) TableName() string {
	return "playlist_songs"
}

func (Song) TableName() string {
	return "songs"
}

func (Account) TableName() string {
	return "accounts"
}
