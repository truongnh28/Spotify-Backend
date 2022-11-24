package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/middleware"
	"spotify/services"
)

var __songService services.SongService
var __authenService services.AuthenService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.SongService:
			__songService = dependency.(services.SongService)
		case services.AuthenService:
			__authenService = dependency.(services.AuthenService)
		}
	}

	songHandler := NewSongHandler(__songService)

	v1 := g.Group("/v1")

	// Authen
	authenRouter := v1.Group("/authen")
	{
		authenRouter.POST("login", login)
		authenRouter.GET("logout", logout)
	}
	songRouter := v1.Group("/song")
	{
		songRouter.Use(middleware.HTTPAuthentication)
		songRouter.GET("", songHandler.GetAll)
	}
}
