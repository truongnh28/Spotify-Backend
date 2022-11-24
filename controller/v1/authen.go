package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"net/url"
	"spotify/config"
	"spotify/dto"
	"spotify/helper/common"
)

func login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, dto.ErrorResponse{
			Message: "Cannot binding request",
		})
		return
	}

	response, err := __authenService.Login(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(500, dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	originURL, err := url.Parse(ctx.GetHeader("origin"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	originHostName := originURL.Hostname()
	glog.Infoln("originHostName: ", originHostName)
	cfg := config.AuthConfig()
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(
		common.CookieName,
		response.AccessToken,
		int(cfg.ExpiredTime), "/", cfg.CookiePath,
		cfg.CookieSecure, false)
	response.AccessToken = ""

	ctx.JSON(200, response)
}

func logout(ctx *gin.Context) {
	glog.Infoln("set cookie is empty")
	cfg := config.AuthConfig()
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(
		common.CookieName,
		"", 0, "/", cfg.CookiePath,
		cfg.CookieSecure, false)
	ctx.JSON(200, nil)
}
