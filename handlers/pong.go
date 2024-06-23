package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Pong = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "pong",
	})
}
