package services

import (
	"spotify/dto"
	"spotify/helper/common"
	"spotify/repositories"
)

type SongService interface {
	GetAllSong() ([]dto.Song, common.SubReturnCode)
}

func NewSongService(songRepo repositories.SongRepository) SongService {
	return &songServiceImpl{
		songRepo: songRepo,
	}
}

type songServiceImpl struct {
	songRepo repositories.SongRepository
}

func (s *songServiceImpl) GetAllSong() ([]dto.Song, common.SubReturnCode) {
	var resp = make([]dto.Song, 0)
	songs, err := s.songRepo.GetAllSong()
	if err != nil {
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
			TrackNumber: val.TrackNumber,
			CreateAt:    val.CreateAt,
			UploadAt:    val.UploadAt,
			YoutubeLink: val.YoutubeLink,
		})
	}
	return resp, common.OK
}
