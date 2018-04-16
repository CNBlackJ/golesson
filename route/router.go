package router

import "github.com/gin-gonic/gin"
import "golesson/controller"

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.GET("/news/:id", Hello.Get)
	Router.GET("/news", Hello.List)
	Router.POST("/news", Hello.Post)
	Router.PUT("/news/:id", Hello.Put)
	Router.DELETE("/news", Hello.Destroy)
}