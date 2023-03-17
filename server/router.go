package server

import (
	"github.com/abgeo/domains/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "assets")
	engine.StaticFile("/favicon.ico", "assets/img/favicon.png")

	defaultController := new(controllers.DefaultController)

	engine.GET("/", defaultController.Index)
	engine.POST("/submit", defaultController.Submit)

	return engine
}
