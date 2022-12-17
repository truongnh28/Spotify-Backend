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

type AlbumHandlerImpl struct {
	albumService services.AlbumService
}

func NewAlbumHandler(AlbumService services.AlbumService) *AlbumHandlerImpl {
	return &AlbumHandlerImpl{
		albumService: AlbumService,
	}
}

func (s *AlbumHandlerImpl) GetAll(context *gin.Context) {
	var (
		out = &dto.AlbumResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.albumService.GetAllAlbum(context)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Albums = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *AlbumHandlerImpl) GetAlbumByID(context *gin.Context) {
	var (
		out = &dto.AlbumResponse{}
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
	response, code := s.albumService.GetAlbumByID(context, uint(id))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Albums = append(out.Albums, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *AlbumHandlerImpl) GetAlbumByName(context *gin.Context) {
	var (
		out = &dto.AlbumResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Param("name")
	response, code := s.albumService.GetAlbumByName(context, name)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Albums = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *AlbumHandlerImpl) AddAlbum(context *gin.Context) {
	var (
		out = &dto.AlbumResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Request.FormValue("name")
	img := context.Request.FormValue("img")
	artistId, err := strconv.Atoi(context.Request.FormValue("artist_id"))
	if err != nil {
		glog.Errorln("parse artistId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.InvalidRequest)
		return
	}
	in := dto.Album{
		Name:     name,
		ArtistID: uint(artistId),
		CoverImg: img,
	}
	code := s.albumService.AddAlbum(context, in)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Albums = append(out.Albums, in)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *AlbumHandlerImpl) UpdateAlbum(context *gin.Context) {
	var (
		out = &dto.AlbumResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Request.FormValue("name")
	img := context.Request.FormValue("img")
	artistId, err := strconv.Atoi(context.Request.FormValue("artist_id"))
	if err != nil {
		glog.Errorln("parse artistId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.InvalidRequest)
		return
	}
	albumId, err := strconv.Atoi(context.Request.FormValue("album_id"))
	if err != nil {
		glog.Errorln("parse albumId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.InvalidRequest)
		return
	}
	in := dto.Album{
		AlbumID:  uint(albumId),
		Name:     name,
		ArtistID: uint(artistId),
		CoverImg: img,
	}
	code := s.albumService.UpdateAlbum(context, in)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Albums = append(out.Albums, in)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
