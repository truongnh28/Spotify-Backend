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

type ArtistHandlerImpl struct {
	artistService services.ArtistService
}

func NewArtistHandler(ArtistService services.ArtistService) *ArtistHandlerImpl {
	return &ArtistHandlerImpl{
		artistService: ArtistService,
	}
}

func (s *ArtistHandlerImpl) GetAll(context *gin.Context) {
	var (
		out = &dto.ArtistResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.artistService.GetAllArtist(context)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Artists = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *ArtistHandlerImpl) GetArtistByID(context *gin.Context) {
	var (
		out = &dto.ArtistResponse{}
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
	response, code := s.artistService.GetArtistByID(context, uint(id))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Artists = append(out.Artists, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *ArtistHandlerImpl) GetArtistByName(context *gin.Context) {
	var (
		out = &dto.ArtistResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Param("name")
	response, code := s.artistService.GetArtistByName(context, name)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Artists = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *ArtistHandlerImpl) AddArtist(context *gin.Context) {
	var (
		out = &dto.ArtistResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Request.FormValue("name")
	img := context.Request.FormValue("img")
	desc := context.Request.FormValue("desc")
	in := dto.Artist{
		Name:        name,
		Description: desc,
		CoverImg:    img,
	}
	code := s.artistService.AddArtist(context, in)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Artists = append(out.Artists, in)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *ArtistHandlerImpl) UpdateArtist(context *gin.Context) {
	var (
		out = &dto.ArtistResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Request.FormValue("name")
	img := context.Request.FormValue("img")
	desc := context.Request.FormValue("desc")
	artistId, err := strconv.Atoi(context.Request.FormValue("artist_id"))
	if err != nil {
		glog.Errorln("parse artistId string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	in := dto.Artist{
		Name:        name,
		Description: desc,
		CoverImg:    img,
		ArtistID:    uint(artistId),
	}
	code := s.artistService.UpdateArtist(context, in)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Artists = append(out.Artists, in)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
