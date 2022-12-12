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
