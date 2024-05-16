package utils

import "github.com/gin-gonic/gin"

// CheckAccessToken 检查 token
var CheckAccessToken = func() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := ctx.Request.Header

		token := h.Get("x-jike-access-token")

		if token == "" {
			ReturnBadRequest(ctx, nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
