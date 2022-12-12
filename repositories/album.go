package repositories

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type AlbumRepository interface {
	GetAllAlbum(ctx context.Context) ([]models.Album, error)
	GetAlbumByID(ctx context.Context, id uint) (models.Album, error)
	GetAlbumByName(ctx context.Context, name string) ([]models.Album, error)
	GetAlbumByArtistID(ctx context.Context, artistId uint) ([]models.Album, error)
}

type albumRepositoryImpl struct {
	database *gorm.DB
}

func NewAlbumRepository(database *gorm.DB) AlbumRepository {
	return &albumRepositoryImpl{
		database: database,
	}
}

func (s *albumRepositoryImpl) GetAllAlbum(ctx context.Context) ([]models.Album, error) {
	albums := make([]models.Album, 0)
	err := s.database.WithContext(ctx).Model(models.Album{}).Find(&albums).Error
	if err != nil {
		return albums, err
	}
	return albums, nil
}

func (s *albumRepositoryImpl) GetAlbumByID(ctx context.Context, id uint) (models.Album, error) {
	var (
		album = models.Album{}
		db    = s.database.WithContext(ctx)
	)
	err := db.Model(models.Album{}).Where("id = ?", id).First(&album).Error
	if err != nil {
		return album, err
	}
	return album, nil
}

func (s *albumRepositoryImpl) GetAlbumByName(ctx context.Context, name string) ([]models.Album, error) {
	var (
		albums = make([]models.Album, 0)
		db     = s.database.WithContext(ctx)
	)
	err := db.Model(models.Album{}).Where("name like ?", "%"+name+"%").Find(&albums).Error
	if err != nil {
		return albums, err
	}
	return albums, nil
}

func (s *albumRepositoryImpl) GetAlbumByArtistID(ctx context.Context, artistId uint) ([]models.Album, error) {
	var (
		albums = make([]models.Album, 0)
		db     = s.database.WithContext(ctx)
	)

	err := db.Model(&models.Album{}).Joins("inner join artists on artists.id = albums.artist_ID").Where("artists.id = ?", artistId).Find(&albums).Error
	if err != nil {
		glog.Errorln("GetAlbumByArtistID Repository err: ", err)
		return albums, err
	}
	return albums, nil
}
