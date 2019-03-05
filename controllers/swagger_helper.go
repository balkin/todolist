package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"golang.org/x/net/webdav"
)

type CountStruct struct {
	Count int `json:"count"`
}

type SimpleTodoItem struct {
	Name string
}

func GinModeReleaseSwagger(h *webdav.Handler) gin.HandlerFunc {
	if gin.Mode() == gin.ReleaseMode {
		return func(c *gin.Context) {
			// Simulate behavior when route unspecified and
			// return 404 HTTP code
			c.String(404, "")
		}
	}
	return ginSwagger.WrapHandler(h)
}
