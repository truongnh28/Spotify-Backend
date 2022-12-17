package v1

import (
	"github.com/gin-gonic/gin"
	"spotify/services"
)

var __songService services.SongService
var __authenService services.AuthenService
var __playListService services.PlayListService
var __albumService services.AlbumService
var __artistService services.ArtistService
var __interactionService services.InteractionService
var __mediaService services.MediaService
var __accountService services.AccountService

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
		case services.InteractionService:
			__interactionService = dependency.(services.InteractionService)
		case services.MediaService:
			__mediaService = dependency.(services.MediaService)
		case services.AccountService:
			__accountService = dependency.(services.AccountService)
		}
	}

	songHandler := NewSongHandler(__songService, __mediaService)
	playListHandler := NewPlayListHandler(__playListService)
	albumHandler := NewAlbumHandler(__albumService)
	artistHandler := NewArtistHandler(__artistService)
	interactionHandler := NewInteractionHandler(__interactionService)
	accountHandler := NewAccountHandler(__accountService)
	v1 := g.Group("/v1")
	// Authen
	authenRouter := v1.Group("/authen")
	{
		authenRouter.POST("login", login)
		authenRouter.GET("logout", logout)
	}
	accountRouter := v1.Group("/accounts")
	{
		accountRouter.POST("/register", accountHandler.CreateAccount)
	}
	songRouter := v1.Group("/songs")
	{
		//songRouter.Use(middleware.HTTPAuthentication)
		songRouter.GET("", songHandler.GetAll)
		songRouter.GET("/id/:id", songHandler.GetSongByID)
		songRouter.GET("/name/:name", songHandler.GetSongByName)
		songRouter.GET("/play_list/:id", songHandler.GetSongByPlayListId)
		songRouter.GET("/album/:id", songHandler.GetSongByAlbumID)
		songRouter.GET("/artist/:id", songHandler.GetSongByArtistID)
		songRouter.GET("/interactions/:id", songHandler.GetSongLikedByUserID)
		songRouter.POST("/add", songHandler.AddSong)
		songRouter.POST("/add_playlist", songHandler.AddSongToPlayList)
		songRouter.DELETE("/remove_playlist", songHandler.RemoveSongToPlayList)
	}
	playListRouter := v1.Group("/playlists")
	{
		//songRouter.Use(middleware.HTTPAuthentication)
		playListRouter.GET("", playListHandler.GetAllPlayList)
		playListRouter.GET("/id/:id", playListHandler.GetPlayListByID)
		playListRouter.GET("/name/:name", playListHandler.GetPlayListByName)
		playListRouter.GET("/user/:id", playListHandler.GetPlayListByUserID)
		playListRouter.POST("/add", playListHandler.AddPlayList)
		playListRouter.POST("/update", playListHandler.UpdatePlayList)
	}
	albumRouter := v1.Group("/albums")
	{
		//albumRouter.Use(middleware.HTTPAuthentication)
		albumRouter.GET("", albumHandler.GetAll)
		albumRouter.GET("/id/:id", albumHandler.GetAlbumByID)
		albumRouter.GET("/name/:name", albumHandler.GetAlbumByName)
		albumRouter.POST("/add", albumHandler.AddAlbum)
		albumRouter.POST("/update", albumHandler.UpdateAlbum)
	}
	artistRouter := v1.Group("/artist")
	{
		//artistRouter.Use(middleware.HTTPAuthentication)
		artistRouter.GET("", artistHandler.GetAll)
		artistRouter.GET("/id/:id", artistHandler.GetArtistByID)
		artistRouter.GET("/name/:name", artistHandler.GetArtistByName)
		artistRouter.POST("/add", artistHandler.AddArtist)
		artistRouter.POST("/update", artistHandler.UpdateArtist)
	}
	interactionRouter := v1.Group("/interaction")
	{
		//interactionRouter.Use(middleware.HTTPAuthentication)
		interactionRouter.POST("/add/:user_id/:song_id", interactionHandler.AddInteraction)
		interactionRouter.DELETE("/remove/:user_id/:song_id", interactionHandler.RemoveInteraction)
	}
}
