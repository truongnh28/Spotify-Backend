package services

import (
	"context"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type AlbumService interface {
	GetAllAlbum(ctx context.Context) ([]dto.Album, common.SubReturnCode)
	GetAlbumByID(ctx context.Context, id uint) (dto.Album, common.SubReturnCode)
	GetAlbumByName(ctx context.Context, name string) ([]dto.Album, common.SubReturnCode)
}

func NewAlbumService(AlbumRepo repositories.AlbumRepository) AlbumService {
	return &albumServiceImpl{
		AlbumRepo: AlbumRepo,
	}
}

type albumServiceImpl struct {
	AlbumRepo repositories.AlbumRepository
}

func (s *albumServiceImpl) GetAllAlbum(ctx context.Context) ([]dto.Album, common.SubReturnCode) {
	var resp = make([]dto.Album, 0)
	albums, err := s.AlbumRepo.GetAllAlbum(ctx)
	if err != nil {
		glog.Infoln("GetAllAlbum service err: ", err)
		return resp, common.SystemError
	}
	for _, album := range albums {
		resp = append(resp, dto.Album{
			AlbumID:  album.AlbumID,
			Name:     album.Name,
			ArtistID: album.ArtistID,
			CoverImg: album.CoverImg,
		})
	}
	return resp, common.OK
}

func (s *albumServiceImpl) GetAlbumByID(ctx context.Context, id uint) (dto.Album, common.SubReturnCode) {
	var (
		resp = dto.Album{}
	)
	album, err := s.AlbumRepo.GetAlbumByID(ctx, id)
	if err != nil {
		glog.Infoln("GetAlbumByID service err: ", err)
		return resp, common.SystemError
	}
	resp = dto.Album{
		AlbumID:  album.AlbumID,
		Name:     album.Name,
		ArtistID: album.ArtistID,
		CoverImg: album.CoverImg,
	}
	return resp, common.OK
}

func (s *albumServiceImpl) GetAlbumByName(ctx context.Context, name string) ([]dto.Album, common.SubReturnCode) {
	var (
		resp = make([]dto.Album, 0)
	)
	albums, err := s.AlbumRepo.GetAlbumByName(ctx, name)
	if err != nil {
		glog.Infoln("GetAlbumByName service err: ", err)
		return resp, common.SystemError
	}
	for _, album := range albums {
		resp = append(resp, dto.Album{
			AlbumID:  album.AlbumID,
			Name:     album.Name,
			ArtistID: album.ArtistID,
			CoverImg: album.CoverImg,
		})
	}
	return resp, common.OK
}
