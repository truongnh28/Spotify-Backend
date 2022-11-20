package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/dto"
	"spotify/helper"
	"spotify/helper/common"
	"spotify/services"
)

type SongHandlerImpl struct {
	songService services.SongService
}

func NewSongHandler(songService services.SongService) *SongHandlerImpl {
	return &SongHandlerImpl{
		songService: songService,
	}
}

func (s *SongHandlerImpl) GetAll(context *gin.Context) {
	var (
		out = &dto.SongResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.songService.GetAllSong()
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Songs = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
