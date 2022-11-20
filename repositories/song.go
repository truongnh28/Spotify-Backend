package repositories

import (
	"context"
	"gorm.io/gorm"
	"spotify/models"
)

type SongRepository interface {
	GetAllSong() ([]models.Songs, error)
	GetSongByID(ctx context.Context, id uint) (models.Songs, error)
	GetSongByName(ctx context.Context, name string) (models.Songs, error)
}

type songRepositoryImpl struct {
	database *gorm.DB
}

func NewSongRepository(database *gorm.DB) SongRepository {
	return &songRepositoryImpl{
		database: database,
	}
}

func (s *songRepositoryImpl) GetAllSong() ([]models.Songs, error) {
	resp := make([]models.Songs, 0)
	err := s.database.Model(models.Songs{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *songRepositoryImpl) GetSongByID(ctx context.Context, id uint) (models.Songs, error) {
	var (
		song models.Songs
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Songs{}).Where("id = ?", id).First(song).Error
	if err != nil {
		return song, err
	}
	return song, nil
}

func (s *songRepositoryImpl) GetSongByName(ctx context.Context, name string) (models.Songs, error) {
	var (
		song models.Songs
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Songs{}).Where("name like ?", "%"+name+"%").First(song).Error
	if err != nil {
		return song, err
	}
	return song, nil
}
