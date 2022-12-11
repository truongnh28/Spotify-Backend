package services

import (
	"context"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type PlayListService interface {
	GetAllPlayList(ctx context.Context) ([]dto.PlayList, common.SubReturnCode)
	GetPlayListByID(ctx context.Context, id uint) (dto.PlayList, common.SubReturnCode)
	GetPlayListByName(ctx context.Context, name string) (dto.PlayList, common.SubReturnCode)
}

func NewPlayListService(playListRepo repositories.PlayListRepository) PlayListService {
	return &playListServiceImpl{
		playListRepo: playListRepo,
	}
}

type playListServiceImpl struct {
	playListRepo repositories.PlayListRepository
}

func (s *playListServiceImpl) GetAllPlayList(ctx context.Context) ([]dto.PlayList, common.SubReturnCode) {
	var resp = make([]dto.PlayList, 0)
	playLists, err := s.playListRepo.GetAllPlayList(ctx)
	if err != nil {
		glog.Infoln("GetAllPlayList service err: ", err)
		return resp, common.SystemError
	}
	for _, val := range playLists {
		resp = append(resp, dto.PlayList{
			PlayListID: val.PlayListID,
			Name:       val.Name,
			CoverImg:   "",
			UserID:     0,
		})
	}
	return resp, common.OK
}

func (s *playListServiceImpl) GetPlayListByID(ctx context.Context, id uint) (dto.PlayList, common.SubReturnCode) {
	var (
		resp = dto.PlayList{}
	)
	playList, err := s.playListRepo.GetPlayListByID(ctx, id)
	if err != nil {
		glog.Infoln("GetPlayListByID service err: ", err)
		return resp, common.SystemError
	}
	resp = dto.PlayList{
		PlayListID: playList.PlayListID,
		Name:       playList.Name,
		CoverImg:   playList.CoverImg,
		UserID:     playList.UserID,
	}
	return resp, common.OK
}

func (s *playListServiceImpl) GetPlayListByName(ctx context.Context, name string) (dto.PlayList, common.SubReturnCode) {
	var (
		resp = dto.PlayList{}
	)
	PlayList, err := s.playListRepo.GetPlayListByName(ctx, name)
	if err != nil {
		glog.Infoln("GetPlayListByName service err: ", err)
		return resp, common.SystemError
	}
	resp = dto.PlayList{
		PlayListID: PlayList.PlayListID,
		Name:       PlayList.Name,
	}
	return resp, common.OK
}
