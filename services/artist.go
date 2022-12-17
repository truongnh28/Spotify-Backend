package services

import (
	"context"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ArtistService interface {
	GetAllArtist(ctx context.Context) ([]dto.Artist, common.SubReturnCode)
	GetArtistByID(ctx context.Context, id uint) (dto.Artist, common.SubReturnCode)
	GetArtistByName(ctx context.Context, name string) ([]dto.Artist, common.SubReturnCode)
	AddArtist(ctx context.Context, ArtistIn dto.Artist) common.SubReturnCode
	UpdateArtist(ctx context.Context, ArtistIn dto.Artist) common.SubReturnCode
}

func NewArtistService(ArtistRepo repositories.ArtistRepository) ArtistService {
	return &artistServiceImpl{
		artistRepo: ArtistRepo,
	}
}

type artistServiceImpl struct {
	artistRepo repositories.ArtistRepository
}

func (s *artistServiceImpl) GetAllArtist(ctx context.Context) ([]dto.Artist, common.SubReturnCode) {
	var resp = make([]dto.Artist, 0)
	artists, err := s.artistRepo.GetAllArtist(ctx)
	if err != nil {
		glog.Infoln("GetAllArtist service err: ", err)
		return resp, common.SystemError
	}
	for _, artist := range artists {
		resp = append(resp, dto.Artist{
			ArtistID:    artist.ArtistID,
			Name:        artist.Name,
			Description: artist.Description,
			CoverImg:    artist.CoverImg,
		})
	}
	return resp, common.OK
}

func (s *artistServiceImpl) GetArtistByID(ctx context.Context, id uint) (dto.Artist, common.SubReturnCode) {
	var (
		resp = dto.Artist{}
	)
	artist, err := s.artistRepo.GetArtistByID(ctx, id)
	if err != nil {
		glog.Infoln("GetArtistByID service err: ", err)
		return resp, common.SystemError
	}
	resp = dto.Artist{
		ArtistID:    artist.ArtistID,
		Name:        artist.Name,
		Description: artist.Description,
		CoverImg:    artist.CoverImg,
	}
	return resp, common.OK
}

func (s *artistServiceImpl) GetArtistByName(ctx context.Context, name string) ([]dto.Artist, common.SubReturnCode) {
	var (
		resp = make([]dto.Artist, 0)
	)
	artists, err := s.artistRepo.GetArtistByName(ctx, name)
	if err != nil {
		glog.Infoln("GetArtistByName service err: ", err)
		return resp, common.SystemError
	}
	for _, artist := range artists {
		resp = append(resp, dto.Artist{
			ArtistID:    artist.ArtistID,
			Name:        artist.Name,
			Description: artist.Description,
			CoverImg:    artist.CoverImg,
		})
	}
	return resp, common.OK
}

func (s *artistServiceImpl) AddArtist(ctx context.Context, artistIn dto.Artist) common.SubReturnCode {
	err := s.artistRepo.AddArtist(ctx, artistIn)
	if err != nil {
		glog.Errorln("Add Artist service err: ", err)
		return common.SystemError
	}
	return common.OK
}

func (s *artistServiceImpl) UpdateArtist(ctx context.Context, artistIn dto.Artist) common.SubReturnCode {
	err := s.artistRepo.UpdateArtist(ctx, artistIn)
	if err != nil {
		glog.Errorln("Update Artist service err: ", err)
		return common.SystemError
	}
	return common.OK
}
