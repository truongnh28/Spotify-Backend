package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/middleware"
	"spotify/services"
)

var __songService services.SongService
var __authenService services.AuthenService
var __playListService services.PlayListService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.SongService:
			__songService = dependency.(services.SongService)
		case services.AuthenService:
			__authenService = dependency.(services.AuthenService)
		case services.PlayListService:
			__playListService = dependency.(services.PlayListService)
		}
	}

	songHandler := NewSongHandler(__songService)
	playListHandler := NewPlayListHandler(__playListService)
	v1 := g.Group("/v1")

	// Authen
	authenRouter := v1.Group("/authen")
	{
		authenRouter.POST("login", login)
		authenRouter.GET("logout", logout)
	}
	songRouter := v1.Group("/songs")
	{
		songRouter.Use(middleware.HTTPAuthentication)
		songRouter.GET("", songHandler.GetAll)
		songRouter.GET("/id/:id", songHandler.GetSongByID)
		songRouter.GET("/name/:name", songHandler.GetSongByName)
		songRouter.GET("/play_list/:id", songHandler.GetSongByPlayListId)
	}
	playListRouter := v1.Group("/playlists")
	{
		playListRouter.GET("", playListHandler.GetAll)
		playListRouter.GET("/id/:id", playListHandler.GetPlayListByID)
		playListRouter.GET("/name/:name", playListHandler.GetPlayListByName)
	}
}
