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
		inter = models.Interaction{
			UserID: userId,
			SongID: songId,
			Liked:  1,
		}
	)
	err := db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).FirstOrCreate(&inter).Error
	if err != nil {
		return models.Interaction{}, err
	}
	return inter, err
}

func (i *interactionRepositoryImpl) RemoveInteraction(ctx context.Context, userId, songId uint) (models.Interaction, error) {
	var (
		db    = i.database.WithContext(ctx)
		inter = models.Interaction{
			UserID: userId,
			SongID: songId,
			Liked:  0,
		}
	)
	err := db.Model(&models.Interaction{}).Where("user_id = ? AND song_id = ?", userId, songId).FirstOrCreate(&inter).Error
	if err != nil {
		return models.Interaction{}, err
	}
	return inter, err
}
