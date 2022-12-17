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
	AddAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode
	UpdateAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode
}

func NewAlbumService(albumRepo repositories.AlbumRepository) AlbumService {
	return &albumServiceImpl{
		albumRepo: albumRepo,
	}
}

type albumServiceImpl struct {
	albumRepo repositories.AlbumRepository
}

func (s *albumServiceImpl) GetAllAlbum(ctx context.Context) ([]dto.Album, common.SubReturnCode) {
	var resp = make([]dto.Album, 0)
	albums, err := s.albumRepo.GetAllAlbum(ctx)
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
	album, err := s.albumRepo.GetAlbumByID(ctx, id)
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
	albums, err := s.albumRepo.GetAlbumByName(ctx, name)
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

func (s *albumServiceImpl) AddAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode {
	err := s.albumRepo.AddAlbum(ctx, albumIn)
	if err != nil {
		glog.Errorln("Add Album service err: ", err)
		return common.SystemError
	}
	return common.OK
}

func (s *albumServiceImpl) UpdateAlbum(ctx context.Context, albumIn dto.Album) common.SubReturnCode {
	err := s.albumRepo.UpdateAlbum(ctx, albumIn)
	if err != nil {
		glog.Errorln("Update Album service err: ", err)
		return common.SystemError
	}
	return common.OK
}
