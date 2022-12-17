package repositories

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"spotify/dto"
	"spotify/models"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type SongRepository interface {
	GetAllSong(ctx context.Context) ([]models.Song, error)
	GetSongByID(ctx context.Context, id uint) (models.Song, error)
	GetSongByName(ctx context.Context, name string) ([]models.Song, error)
	GetSongByPlayListID(ctx context.Context, id uint) ([]models.Song, error)
	GetSongByArtistID(ctx context.Context, artistId uint) ([]models.Song, error)
	GetSongByAlbumID(ctx context.Context, albumId uint) ([]models.Song, error)
	GetSongLikedByUserID(ctx context.Context, userId uint) ([]models.Song, error)
	AddSong(ctx context.Context, songIn dto.Song) error
	AddSongToPlayList(ctx context.Context, songId, playlistId uint) error
	RemoveSongToPlayList(ctx context.Context, songId, playlistId uint) error
}

type songRepositoryImpl struct {
	database *gorm.DB
}

func NewSongRepository(database *gorm.DB) SongRepository {
	return &songRepositoryImpl{
		database: database,
	}
}

func (s *songRepositoryImpl) GetAllSong(ctx context.Context) ([]models.Song, error) {
	var (
		resp = make([]models.Song, 0)
	)
	err := s.database.WithContext(ctx).Model(models.Song{}).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *songRepositoryImpl) GetSongByID(ctx context.Context, id uint) (models.Song, error) {
	var (
		song = models.Song{}
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Song{}).Where("id = ?", id).First(&song).Error
	if err != nil {
		glog.Errorln("GetSongByID Repository err: ", err)
		return song, err
	}
	return song, nil
}

func (s *songRepositoryImpl) GetSongByName(ctx context.Context, name string) ([]models.Song, error) {
	var (
		song = make([]models.Song, 0)
		db   = s.database.WithContext(ctx)
	)
	err := db.Model(models.Song{}).Where("name like ?", "%"+name+"%").Find(&song).Error
	if err != nil {
		glog.Errorln("GetSongByName Repository err: ", err)
		return song, err
	}
	return song, nil
}

func (s *songRepositoryImpl) GetSongByPlayListID(ctx context.Context, playListId uint) ([]models.Song, error) {
	var (
		songs         = make([]models.Song, 0)
		songIds       = make([]uint, 0)
		playListSongs = make([]models.PlayListSong, 0)
		db            = s.database.WithContext(ctx)
	)
	err := db.Model(&models.PlayListSong{}).Where("playlist_id = ?", playListId).Find(&playListSongs).Error
	if err != nil {
		glog.Errorln("GetSongByPlayListID Repository err: ", err)
		return songs, err
	}
	for _, val := range playListSongs {
		songIds = append(songIds, val.SongID)
	}
	err = db.Model(&models.Song{}).Where("id IN ?", songIds).Find(&songs).Error
	if err != nil {
		glog.Errorln("GetSongByPlayListID Repository err: ", err)
		return songs, err
	}
	return songs, nil
}

func (s *songRepositoryImpl) GetSongByArtistID(ctx context.Context, artistId uint) ([]models.Song, error) {
	var (
		songs = make([]models.Song, 0)
		db    = s.database.WithContext(ctx)
	)

	err := db.Model(&models.Song{}).Joins("inner join artists on artists.id = songs.artist_id").Where("artists.id = ?", artistId).Find(&songs).Error
	if err != nil {
		glog.Errorln("GetSongByArtistID Repository err: ", err)
		return songs, err
	}
	return songs, nil
}

func (s *songRepositoryImpl) GetSongByAlbumID(ctx context.Context, albumId uint) ([]models.Song, error) {
	var (
		songs = make([]models.Song, 0)
		db    = s.database.WithContext(ctx)
	)

	err := db.Model(&models.Song{}).Joins("inner join albums on albums.id = songs.album_id").Where("albums.id = ?", albumId).Find(&songs).Error
	if err != nil {
		glog.Errorln("GetSongByAlbumID Repository err: ", err)
		return songs, err
	}
	return songs, nil
}

func (s *songRepositoryImpl) GetSongLikedByUserID(ctx context.Context, userId uint) ([]models.Song, error) {
	var (
		songs        = make([]models.Song, 0)
		interactions = make([]models.Interaction, 0)
		songIds      = make([]uint, 0)
		db           = s.database.WithContext(ctx)
	)
	err := db.Model(&models.Interaction{}).Joins("inner join accounts on accounts.id = interactions.user_id").Where("interactions.liked = 1 and interactions.user_id = ?", userId).Find(&interactions).Error
	if err != nil {
		glog.Errorln("GetSongLikedByUserID Repository err: ", err)
		return songs, err
	}
	for _, interaction := range interactions {
		songIds = append(songIds, interaction.SongID)
	}
	err = db.Model(&models.Song{}).Where("songs.id IN ?", songIds).Find(&songs).Error
	if err != nil {
		glog.Errorln("GetSongByAlbumID Repository err: ", err)
		return songs, err
	}
	return songs, nil
}

func (s *songRepositoryImpl) AddSong(ctx context.Context, songIn dto.Song) error {
	var (
		song = models.Song{
			Name:        songIn.Name,
			AlbumID:     songIn.AlbumID,
			ArtistID:    songIn.ArtistID,
			Lyrics:      songIn.Lyrics,
			Length:      songIn.Length,
			URL:         songIn.URL,
			YoutubeLink: songIn.YoutubeLink,
			SongCloudId: songIn.SongCloudId,
		}
		db = s.database.WithContext(ctx)
	)
	err := db.Model(&models.Song{}).Create(&song).Error
	if err != nil {
		glog.Errorln("AddSong repository err: ", err)
		return err
	}
	return nil
}

func (s *songRepositoryImpl) AddSongToPlayList(ctx context.Context, songId, playlistId uint) error {
	var (
		playListSong = models.PlayListSong{
			SongID:     songId,
			PlayListID: playlistId,
		}
		db = s.database.WithContext(ctx)
	)
	err := db.Model(&models.PlayListSong{}).Create(&playListSong).Error
	if err != nil {
		glog.Errorln("AddSongToPlayList repository err: ", err)
		return err
	}
	return nil
}

func (s *songRepositoryImpl) RemoveSongToPlayList(ctx context.Context, songId, playlistId uint) error {
	var (
		playListSong = models.PlayListSong{
			SongID:     songId,
			PlayListID: playlistId,
		}
		db = s.database.WithContext(ctx)
	)
	err := db.Model(&models.PlayListSong{}).Where("song_id = ? AND playlist_id = ?", songId, playlistId).Delete(&playListSong).Error
	if err != nil {
		glog.Errorln("RemoveSongToPlayList repository err: ", err)
		return err
	}
	return nil
}
