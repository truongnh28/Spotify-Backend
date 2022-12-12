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

type PlayListHandlerImpl struct {
	playListService services.PlayListService
}

func NewPlayListHandler(playListService services.PlayListService) *PlayListHandlerImpl {
	return &PlayListHandlerImpl{
		playListService: playListService,
	}
}

func (s *PlayListHandlerImpl) GetAll(context *gin.Context) {
	var (
		out = &dto.PlayListResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.playListService.GetAllPlayList(context)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.PlayLists = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *PlayListHandlerImpl) GetPlayListByID(context *gin.Context) {
	var (
		out = &dto.PlayListResponse{}
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
	response, code := s.playListService.GetPlayListByID(context, uint(id))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.PlayLists = append(out.PlayLists, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *PlayListHandlerImpl) GetPlayListByName(context *gin.Context) {
	var (
		out = &dto.PlayListResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	name := context.Param("name")
	response, code := s.playListService.GetPlayListByName(context, name)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.PlayLists = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *PlayListHandlerImpl) GetPlayListByUserID(context *gin.Context) {
	var (
		out = &dto.PlayListResponse{}
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
	response, code := s.playListService.GetPlayListByUserID(context, uint(id))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.PlayLists = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
