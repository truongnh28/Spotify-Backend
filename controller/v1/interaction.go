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

type InteractionHandlerImpl struct {
	interactionService services.InteractionService
}

func NewInteractionHandler(interactionService services.InteractionService) *InteractionHandlerImpl {
	return &InteractionHandlerImpl{
		interactionService: interactionService,
	}
}

func (s *InteractionHandlerImpl) AddInteraction(context *gin.Context) {
	var (
		out = &dto.InteractionResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	userIds := context.Param("user_id")
	userId, err := strconv.Atoi(userIds)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	songIds := context.Param("song_id")
	songId, err := strconv.Atoi(songIds)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.interactionService.AddInteraction(context, uint(userId), uint(songId))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Interactions = append(out.Interactions, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *InteractionHandlerImpl) RemoveInteraction(context *gin.Context) {
	var (
		out = &dto.InteractionResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	userIds := context.Param("user_id")
	userId, err := strconv.Atoi(userIds)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	songIds := context.Param("song_id")
	songId, err := strconv.Atoi(songIds)
	if err != nil {
		glog.Errorln("parse id string to int err: ", err)
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.interactionService.RemoveInteraction(context, uint(userId), uint(songId))
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Interactions = append(out.Interactions, response)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
