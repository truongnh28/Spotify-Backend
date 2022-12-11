package repositories

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type SongRepository interface {
	GetAllSong(ctx context.Context) ([]models.Song, error)
	GetSongByID(ctx context.Context, id uint) (models.Song, error)
	GetSongByName(ctx context.Context, name string) ([]models.Song, error)
	GetSongByPlayListID(ctx context.Context, id uint) ([]models.Song, error)
}

type songRepositoryImpl struct {
	database *gorm.DB
}

func NewSongRepository(database *gorm.DB) SongRepository {
	return &songRepositoryImpl{
		database: database,
	}
}

func (s *songRepositoryImpl) GetAllSong(ctx context.Context) ([]models.Song, error) {
	resp := make([]models.Song, 0)
	err := s.database.WithContext(ctx).Model(models.Song{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *songRepositoryImpl) GetSongByID(ctx context.Context, id uint) (models.Song, error) {
	var (
		song = models.Song{}
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Song{}).Where("id = ?", id).First(&song).Error
	if err != nil {
		glog.Errorln("GetSongByID Repository err: ", err)
		return song, err
	}
	return song, nil
}

func (s *songRepositoryImpl) GetSongByName(ctx context.Context, name string) ([]models.Song, error) {
	var (
		song = make([]models.Song, 0)
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Song{}).Where("name like ?", "%"+name+"%").Find(&song).Error
	if err != nil {
		glog.Errorln("GetSongByName Repository err: ", err)
		return song, err
	}
	return song, nil
}

func (s *songRepositoryImpl) GetSongByPlayListID(ctx context.Context, id uint) ([]models.Song, error) {
	var (
		songs = make([]models.Song, 0)
		db    = s.database.WithContext(ctx)
	)

	err := db.Model(&models.Song{}).Joins("inner join playlists_songs on playlists_songs.song_ID = songs.id").Where("songs.id = ?", id).Find(&songs).Error
	if err != nil {
		glog.Errorln("GetSongByPlayListID Repository err: ", err)
		return songs, err
	}
	return songs, nil
}
