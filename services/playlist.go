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
	GetPlayListByName(ctx context.Context, name string) ([]dto.PlayList, common.SubReturnCode)
	GetPlayListByUserID(ctx context.Context, userId uint) ([]dto.PlayList, common.SubReturnCode)
	AddPlayList(ctx context.Context, playListIn dto.PlayList) common.SubReturnCode
	UpdatePlayList(ctx context.Context, playListIn dto.PlayList) common.SubReturnCode
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
			CoverImg:   val.CoverImg,
			UserID:     val.UserID,
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

func (s *playListServiceImpl) GetPlayListByName(ctx context.Context, name string) ([]dto.PlayList, common.SubReturnCode) {
	var (
		resp = make([]dto.PlayList, 0)
	)
	playLists, err := s.playListRepo.GetPlayListByName(ctx, name)
	if err != nil {
		glog.Infoln("GetPlayListByName service err: ", err)
		return resp, common.SystemError
	}
	for _, playList := range playLists {
		resp = append(resp, dto.PlayList{
			PlayListID: playList.PlayListID,
			Name:       playList.Name,
			CoverImg:   playList.CoverImg,
			UserID:     playList.UserID,
		})
	}
	return resp, common.OK
}

func (s *playListServiceImpl) GetPlayListByUserID(ctx context.Context, userId uint) ([]dto.PlayList, common.SubReturnCode) {
	var (
		resp = make([]dto.PlayList, 0)
	)
	playLists, err := s.playListRepo.GetPlayListByUserID(ctx, userId)
	if err != nil {
		glog.Infoln("GetPlayListByUserID service err: ", err)
		return resp, common.SystemError
	}
	for _, playList := range playLists {
		resp = append(resp, dto.PlayList{
			PlayListID: playList.PlayListID,
			Name:       playList.Name,
			CoverImg:   playList.CoverImg,
			UserID:     playList.UserID,
		})
	}
	return resp, common.OK
}

func (s *playListServiceImpl) AddPlayList(ctx context.Context, playListIn dto.PlayList) common.SubReturnCode {
	err := s.playListRepo.AddPlayList(ctx, playListIn)
	if err != nil {
		glog.Errorln("Add playlist service err: ", err)
		return common.SystemError
	}
	return common.OK
}

func (s *playListServiceImpl) UpdatePlayList(ctx context.Context, playListIn dto.PlayList) common.SubReturnCode {
	err := s.playListRepo.UpdatePlayList(ctx, playListIn)
	if err != nil {
		glog.Errorln("Update playlist service err: ", err)
		return common.SystemError
	}
	return common.OK
}
