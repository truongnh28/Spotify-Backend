package services

import (
	"context"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type SongService interface {
	GetAllSong(ctx context.Context) ([]dto.Song, common.SubReturnCode)
	GetSongByID(ctx context.Context, id uint) (dto.Song, common.SubReturnCode)
	GetSongByName(ctx context.Context, name string) ([]dto.Song, common.SubReturnCode)
	GetSongByPlayListID(ctx context.Context, playListId uint) ([]dto.Song, common.SubReturnCode)
	GetSongByArtistID(ctx context.Context, artistId uint) ([]dto.Song, common.SubReturnCode)
	GetSongByAlbumID(ctx context.Context, albumId uint) ([]dto.Song, common.SubReturnCode)
}

func NewSongService(songRepo repositories.SongRepository) SongService {
	return &songServiceImpl{
		songRepo: songRepo,
	}
}

type songServiceImpl struct {
	songRepo repositories.SongRepository
}

func (s *songServiceImpl) GetAllSong(ctx context.Context) ([]dto.Song, common.SubReturnCode) {
	var resp = make([]dto.Song, 0)
	songs, err := s.songRepo.GetAllSong(ctx)
	if err != nil {
		glog.Infoln("GetAllSong service err: ", err)
		return resp, common.SystemError
	}
	for _, val := range songs {
		resp = append(resp, dto.Song{
			SongID:      val.SongID,
			Name:        val.Name,
			AlbumID:     val.AlbumID,
			ArtistID:    val.ArtistID,
			Lyrics:      val.Lyrics,
			Length:      val.Length,
			URL:         val.URL,
			YoutubeLink: val.YoutubeLink,
		})
	}
	return resp, common.OK
}

func (s *songServiceImpl) GetSongByID(ctx context.Context, id uint) (dto.Song, common.SubReturnCode) {
	var (
		resp = dto.Song{}
	)
	song, err := s.songRepo.GetSongByID(ctx, id)
	if err != nil {
		glog.Infoln("GetSongByID service err: ", err)
		return resp, common.SystemError
	}
	resp = dto.Song{
		SongID:      song.SongID,
		Name:        song.Name,
		AlbumID:     song.AlbumID,
		ArtistID:    song.ArtistID,
		Lyrics:      song.Lyrics,
		Length:      song.Length,
		URL:         song.URL,
		YoutubeLink: song.YoutubeLink,
	}
	return resp, common.OK
}

func (s *songServiceImpl) GetSongByName(ctx context.Context, name string) ([]dto.Song, common.SubReturnCode) {
	var (
		resp = make([]dto.Song, 0)
	)
	songs, err := s.songRepo.GetSongByName(ctx, name)
	if err != nil {
		glog.Infoln("GetSongByName service err: ", err)
		return resp, common.SystemError
	}
	for _, song := range songs {
		resp = append(resp, dto.Song{
			SongID:      song.SongID,
			Name:        song.Name,
			AlbumID:     song.AlbumID,
			ArtistID:    song.ArtistID,
			Lyrics:      song.Lyrics,
			Length:      song.Length,
			URL:         song.URL,
			YoutubeLink: song.YoutubeLink,
		})
	}
	return resp, common.OK
}

func (s *songServiceImpl) GetSongByPlayListID(ctx context.Context, id uint) ([]dto.Song, common.SubReturnCode) {
	var (
		resp = make([]dto.Song, 0)
	)
	songs, err := s.songRepo.GetSongByPlayListID(ctx, id)
	if err != nil {
		glog.Infoln("GetSongByPlayListID service err: ", err)
		return resp, common.SystemError
	}
	for _, song := range songs {
		resp = append(resp, dto.Song{
			SongID:      song.SongID,
			Name:        song.Name,
			AlbumID:     song.AlbumID,
			ArtistID:    song.ArtistID,
			Lyrics:      song.Lyrics,
			Length:      song.Length,
			URL:         song.URL,
			YoutubeLink: song.YoutubeLink,
		})
	}
	return resp, common.OK
}

func (s *songServiceImpl) GetSongByArtistID(ctx context.Context, artistId uint) ([]dto.Song, common.SubReturnCode) {
	var (
		resp = make([]dto.Song, 0)
	)
	songs, err := s.songRepo.GetSongByArtistID(ctx, artistId)
	if err != nil {
		glog.Infoln("GetSongByArtistID service err: ", err)
		return resp, common.SystemError
	}
	for _, song := range songs {
		resp = append(resp, dto.Song{
			SongID:      song.SongID,
			Name:        song.Name,
			AlbumID:     song.AlbumID,
			ArtistID:    song.ArtistID,
			Lyrics:      song.Lyrics,
			Length:      song.Length,
			URL:         song.URL,
			YoutubeLink: song.YoutubeLink,
		})
	}
	return resp, common.OK
}

func (s *songServiceImpl) GetSongByAlbumID(ctx context.Context, albumId uint) ([]dto.Song, common.SubReturnCode) {
	var (
		resp = make([]dto.Song, 0)
	)
	songs, err := s.songRepo.GetSongByAlbumID(ctx, albumId)
	if err != nil {
		glog.Infoln("GetSongByAlbumID service err: ", err)
		return resp, common.SystemError
	}
	for _, song := range songs {
		resp = append(resp, dto.Song{
			SongID:      song.SongID,
			Name:        song.Name,
			AlbumID:     song.AlbumID,
			ArtistID:    song.ArtistID,
			Lyrics:      song.Lyrics,
			Length:      song.Length,
			URL:         song.URL,
			YoutubeLink: song.YoutubeLink,
		})
	}
	return resp, common.OK
}
