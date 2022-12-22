package repositories

import (
	"context"
	"gorm.io/gorm"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type InteractionRepository interface {
	AddInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error)
	RemoveInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error)
}

type interactionRepositoryImpl struct {
	database *gorm.DB
}

func NewInteractionRepository(database *gorm.DB) InteractionRepository {
	return &interactionRepositoryImpl{
		database: database,
	}
}

func (i *interactionRepositoryImpl) AddInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error) {
	var (
		db    = i.database.WithContext(ctx)
		inter = models.Interaction{}
	)
	db.Begin()
	query := db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).First(&inter)
	err := query.Error
	if err != nil {
		if err.Error() != "record not found" {
			db.Rollback()
			return inter, err
		}
	}
	if err != nil {
		inter = models.Interaction{
			UserID: userId,
			SongID: songId,
			Liked:  1,
		}
		err = db.Model(&models.Interaction{}).Create(&inter).Error
	} else {
		inter.Liked = 1
		err = db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).Update("liked", 1).Error
	}
	if err != nil {
		db.Rollback()
		return inter, err
	}
	db.Commit()
	return inter, nil
}

func (i *interactionRepositoryImpl) RemoveInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error) {
	var (
		db    = i.database.WithContext(ctx)
		inter = models.Interaction{}
	)
	db.Begin()
	query := db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).First(&inter)
	err := query.Error
	if err != nil {
		if err.Error() != "record not found" {
			db.Rollback()
			return inter, err
		}
	}
	if err != nil {
		inter = models.Interaction{
			UserID: userId,
			SongID: songId,
			Liked:  0,
		}
		err = db.Model(&models.Interaction{}).Create(&inter).Error
	} else {
		inter.Liked = 0
		err = db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).Update("liked", 0).Error
	}
	if err != nil {
		db.Rollback()
		return inter, err
	}
	db.Commit()
	return inter, nil
}
