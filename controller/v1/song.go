package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"spotify/dto"
	"spotify/helper"
	"spotify/helper/common"
	"spotify/services"
	"strconv"
)

type SongHandlerImpl struct {
	songService  services.SongService
	mediaService services.MediaService
}

func NewSongHandler(songService services.SongService, mediaService services.MediaService) *SongHandlerImpl {
	return &SongHandlerImpl{
		songService:  songService,
		mediaService: mediaService,
	}
}

func (s *SongHandlerImpl) GetAll(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.songService.GetAllSong(context)
	if code != common.OK {
		glog.Errorln("GetAll controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongByID(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	ids := context.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.songService.GetSongByID(context, uint(id))
	if code != common.OK {
		glog.Errorln("GetSongByID controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = append(out.Songs, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongByName(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Param("name")
	response, code := s.songService.GetSongByName(context, name)
	if code != common.OK {
		glog.Errorln("GetSongByName controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongByPlayListId(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	ids := context.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.songService.GetSongByPlayListID(context, uint(id))
	if code != common.OK {
		glog.Errorln("GetSongByPlayListId controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongByArtistID(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	ids := context.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.songService.GetSongByArtistID(context, uint(id))
	if code != common.OK {
		glog.Errorln("GetSongByArtistID controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongByAlbumID(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	ids := context.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.songService.GetSongByAlbumID(context, uint(id))
	if code != common.OK {
		glog.Errorln("GetSongByAlbumID controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) GetSongLikedByUserID(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	ids := context.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.songService.GetSongLikedByUserID(context, uint(id))
	if code != common.OK {
		glog.Errorln("GetSongLikedByUserID controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) AddSong(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		glog.Errorln("Get file from request err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.InvalidRequest)
		return
	}
	name := context.Request.FormValue("name")
	lyric := context.Request.FormValue("lyric")
	ytb := context.Request.FormValue("ytb")
	albumId, err := strconv.Atoi(context.Request.FormValue("album_id"))
	if err != nil {
		glog.Errorln("parse albumId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	artistId, err := strconv.Atoi(context.Request.FormValue("artist_id"))
	if err != nil {
		glog.Errorln("parse artistID string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	length, err := strconv.Atoi(context.Request.FormValue("length"))
	if err != nil {
		glog.Errorln("parse length string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	cldResp, code := s.mediaService.Upload(dto.UploadIn{
		File: file,
	})
	songIn := dto.Song{
		Name:        name,
		AlbumID:     uint(albumId),
		ArtistID:    uint(artistId),
		Lyrics:      lyric,
		Length:      uint(length),
		URL:         cldResp.URL,
		YoutubeLink: ytb,
		SongCloudId: cldResp.AssetID,
	}
	code = s.songService.AddSong(context, songIn)
	if code != common.OK {
		glog.Errorln("AddSong controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = append(out.Songs, songIn)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) AddSongToPlayList(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	songId, err := strconv.Atoi(context.Request.FormValue("song_id"))
	if err != nil {
		glog.Errorln("parse songId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	playlistId, err := strconv.Atoi(context.Request.FormValue("playlist_id"))
	if err != nil {
		glog.Errorln("parse playlistId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	code := s.songService.AddSongToPlayList(context, uint(songId), uint(playlistId))
	if code != common.OK {
		glog.Errorln("AddSongToPlayList controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = nil
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *SongHandlerImpl) RemoveSongToPlayList(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	songId, err := strconv.Atoi(context.Request.FormValue("song_id"))
	if err != nil {
		glog.Errorln("parse songId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	playlistId, err := strconv.Atoi(context.Request.FormValue("playlist_id"))
	if err != nil {
		glog.Errorln("parse playlistId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	code := s.songService.RemoveSongToPlayList(context, uint(songId), uint(playlistId))
	if code != common.OK {
		glog.Errorln("RemoveSongToPlayList controller err: ", code)
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = nil
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
