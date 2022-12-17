package services

import (
	"context"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type InteractionService interface {
	AddInteraction(ctx context.Context, userId, songId uint) (dto.Interaction, common.SubReturnCode)
	RemoveInteraction(ctx context.Context, userId, songId uint) (dto.Interaction, common.SubReturnCode)
}

func NewInteractionService(InteractionRepo repositories.InteractionRepository) InteractionService {
	return &interactionServiceImpl{
		interactionRepo: InteractionRepo,
	}
}

type interactionServiceImpl struct {
	interactionRepo repositories.InteractionRepository
}

func (i *interactionServiceImpl) AddInteraction(ctx context.Context, userId, songId uint) (dto.Interaction, common.SubReturnCode) {
	var (
		resp = dto.Interaction{}
	)
	inter, err := i.interactionRepo.AddInteraction(ctx, userId, songId)
	if err != nil {
		glog.Infoln("AddInteraction service err: ", err)
		return resp, common.SystemError
	}
	resp.Liked = inter.Liked
	resp.UserID = inter.UserID
	resp.SongID = inter.SongID
	return resp, common.OK
}

func (i *interactionServiceImpl) RemoveInteraction(ctx context.Context, userId, songId uint) (dto.Interaction, common.SubReturnCode) {
	var (
		resp = dto.Interaction{}
	)
	inter, err := i.interactionRepo.RemoveInteraction(ctx, userId, songId)
	if err != nil {
		glog.Infoln("RemoveInteraction service err: ", err)
		return resp, common.SystemError
	}
	resp.Liked = inter.Liked
	resp.UserID = inter.UserID
	resp.SongID = inter.SongID
	return resp, common.OK
}
