package repositories

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type PlayListRepository interface {
	GetAllPlayList(ctx context.Context) ([]models.PlayList, error)
	GetPlayListByID(ctx context.Context, id uint) (models.PlayList, error)
	GetPlayListByName(ctx context.Context, name string) ([]models.PlayList, error)
	GetPlayListByUserID(ctx context.Context, userId uint) ([]models.PlayList, error)
}

type playlistRepositoryImpl struct {
	database *gorm.DB
}

func NewPlayListRepository(database *gorm.DB) PlayListRepository {
	return &playlistRepositoryImpl{
		database: database,
	}
}

func (s *playlistRepositoryImpl) GetAllPlayList(ctx context.Context) ([]models.PlayList, error) {
	resp := make([]models.PlayList, 0)
	err := s.database.WithContext(ctx).Model(models.PlayList{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *playlistRepositoryImpl) GetPlayListByID(ctx context.Context, id uint) (models.PlayList, error) {
	var (
		playList = models.PlayList{}
		db       = s.database.WithContext(ctx)
	)
	err := db.Model(models.PlayList{}).Where("id = ?", id).First(&playList).Error
	if err != nil {
		return playList, err
	}
	return playList, nil
}

func (s *playlistRepositoryImpl) GetPlayListByName(ctx context.Context, name string) ([]models.PlayList, error) {
	var (
		playLists = make([]models.PlayList, 0)
		db        = s.database.WithContext(ctx)
	)
	err := db.Model(models.PlayList{}).Where("name like ?", "%"+name+"%").Find(&playLists).Error
	if err != nil {
		return playLists, err
	}
	return playLists, nil
}

func (s *playlistRepositoryImpl) GetPlayListByUserID(ctx context.Context, userId uint) ([]models.PlayList, error) {
	var (
		playLists = make([]models.PlayList, 0)
		db        = s.database.WithContext(ctx)
	)

	err := db.Model(&models.PlayList{}).Joins("inner join accounts on accounts.artist_ID = playlists.id").Where("accounts.id = ?", userId).Find(&playLists).Error
	if err != nil {
		glog.Errorln("GetPlayListByArtistID Repository err: ", err)
		return playLists, err
	}
	return playLists, nil
}
