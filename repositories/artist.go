package repositories

import (
	"context"
	"gorm.io/gorm"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ArtistRepository interface {
	GetAllArtist(ctx context.Context) ([]models.Artist, error)
	GetArtistByID(ctx context.Context, id uint) (models.Artist, error)
	GetArtistByName(ctx context.Context, name string) ([]models.Artist, error)
}

type artistRepositoryImpl struct {
	database *gorm.DB
}

func NewArtistRepository(database *gorm.DB) ArtistRepository {
	return &artistRepositoryImpl{
		database: database,
	}
}

func (s *artistRepositoryImpl) GetAllArtist(ctx context.Context) ([]models.Artist, error) {
	resp := make([]models.Artist, 0)
	err := s.database.WithContext(ctx).Model(models.Artist{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *artistRepositoryImpl) GetArtistByID(ctx context.Context, id uint) (models.Artist, error) {
	var (
		Artist = models.Artist{}
		db     = s.database.WithContext(ctx)
	)
	err := db.Model(models.Artist{}).Where("id = ?", id).First(&Artist).Error
	if err != nil {
		return Artist, err
	}
	return Artist, nil
}

func (s *artistRepositoryImpl) GetArtistByName(ctx context.Context, name string) ([]models.Artist, error) {
	var (
		artists = make([]models.Artist, 0)
		db      = s.database.WithContext(ctx)
	)
	err := db.Model(models.Artist{}).Where("name like ?", "%"+name+"%").Find(&artists).Error
	if err != nil {
		return artists, err
	}
	return artists, nil
}
