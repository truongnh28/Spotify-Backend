package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/middleware"
	"spotify/services"
)

var __songService services.SongService
var __authenService services.AuthenService
var __playListService services.PlayListService
var __albumService services.AlbumService
var __artistService services.ArtistService

func InitRoutes(g *gin.RouterGroup, dependencies ...interface{}) {
	for _, dependency := range dependencies {
		switch dependency.(type) {
		case services.SongService:
			__songService = dependency.(services.SongService)
		case services.AuthenService:
			__authenService = dependency.(services.AuthenService)
		case services.PlayListService:
			__playListService = dependency.(services.PlayListService)
		case services.AlbumService:
			__albumService = dependency.(services.AlbumService)
		case services.ArtistService:
			__artistService = dependency.(services.ArtistService)

		}
	}

	songHandler := NewSongHandler(__songService)
	playListHandler := NewPlayListHandler(__playListService)
	albumHandler := NewAlbumHandler(__albumService)
	artistHandler := NewArtistHandler(__artistService)
	v1 := g.Group("/v1")
	v1.Use(middleware.HTTPAuthentication)
	// Authen
	authenRouter := v1.Group("/authen")
	{
		authenRouter.POST("login", login)
		authenRouter.GET("logout", logout)
	}
	songRouter := v1.Group("/songs")
	{
		songRouter.GET("", songHandler.GetAll)
		songRouter.GET("/id/:id", songHandler.GetSongByID)
		songRouter.GET("/name/:name", songHandler.GetSongByName)
		songRouter.GET("/play_list/:id", songHandler.GetSongByPlayListId)
		songRouter.GET("/album/:id", songHandler.GetSongByAlbumID)
		songRouter.GET("/artist/:id", songHandler.GetSongByArtistID)
		songRouter.GET("/interactions/:id", songHandler.GetSongLikedByUserID)
	}
	playListRouter := v1.Group("/playlists")
	{
		playListRouter.GET("", playListHandler.GetAll)
		playListRouter.GET("/id/:id", playListHandler.GetPlayListByID)
		playListRouter.GET("/name/:name", playListHandler.GetPlayListByName)
		playListRouter.GET("/user/:id", playListHandler.GetPlayListByUserID)
	}
	albumRouter := v1.Group("/albums")
	{
		albumRouter.GET("", albumHandler.GetAll)
		albumRouter.GET("/id/:id", albumHandler.GetAlbumByID)
		albumRouter.GET("/name/:name", albumHandler.GetAlbumByName)
	}
	artistRouter := v1.Group("/artist")
	{
		artistRouter.GET("", artistHandler.GetAll)
		artistRouter.GET("/id/:id", artistHandler.GetArtistByID)
		artistRouter.GET("/name/:name", artistHandler.GetArtistByName)
	}
}
