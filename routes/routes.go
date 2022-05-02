package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routeEngine := gin.Default()
	routeEngine.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<div><h1>Welcome to Prime Number Distributed Sytem.</h1><h2>Everything is fine...</h2><h3>Made in &#128151; with Go </h3></div>"))
	})

	return routeEngine
}
