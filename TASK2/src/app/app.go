package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rampo0/bibit/layer/application/services"
	"github.com/rampo0/bibit/layer/infrastructure/rest"
	"github.com/rampo0/bibit/layer/presentation/controllers"
)

var (
	router = gin.Default()
)

func Run() {

	movieHandler := controllers.NewMovieHandler(services.NewMovieService(rest.NewMovieRepository()))

	router.GET("/detail", movieHandler.Detail)
	router.GET("/search", movieHandler.Search)

	router.Run()
}
