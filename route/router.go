package router

import (
	"golesson/controller/hello"
	"golesson/controller/graphql"
	
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.GET("/news/:id", hello.Get)
	// Router.GET("/news", Hello.List)
	Router.POST("/news", hello.Post)
	Router.PUT("/news/:id", hello.Put)
	Router.DELETE("/news", hello.Destroy)

	Router.POST("/graphql", graphql.GraphqlHandler())
	Router.GET("/graphql", graphql.GraphqlHandler())
}