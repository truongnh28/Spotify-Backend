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
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = append(out.Songs, songIn)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
