package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	U "github.com/ultrazg/xyz/pkg/utils"
	"github.com/ultrazg/xyz/routes"
	"log"
	"net/http"
)

func Start() {
	p, d := U.InitFlag()
	port := fmt.Sprintf("%d", p)

	U.P(port)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(Cors())

	routes.RegisterRouters(engine)

	log.Printf("server start on %s", port)

	if d {
		err := U.OpenBrowser("http://localhost:" + port + "/doc")
		if err != nil {
			log.Println("open browser fail")
		}
	}

	err := engine.Run(":" + port)
	if err != nil {
		log.Println("server start fail")

		return
	}
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
