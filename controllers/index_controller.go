package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func IndexController(ctx *gin.Context) {
	time.Sleep(time.Minute)
	ctx.String(200, "index page")
}

func PingController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"pong": 1})
}
