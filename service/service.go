package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/router"
	"github.com/ultrazg/xyz/utils"
	"log"
	"net/http"
)

func Start() error {
	p, d := utils.InitFlag()

	err := utils.CheckPort(p)
	if err != nil {
		return err
	}

	port := fmt.Sprintf("%d", p)

	utils.P(port)

	go func() {
		err := utils.CheckUpgrade()
		if err != nil {
			return
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(Cors())

	router.RegisterRouters(engine)

	log.Printf("server start on %s", port)

	if d {
		err := utils.OpenBrowser("http://localhost:" + port + "/docs")
		if err != nil {
			log.Println("open browser fail")

			return fmt.Errorf("open browser fail")
		}
	}

	err = engine.Run(":" + port)
	if err != nil {
		log.Println("server start fail")

		return fmt.Errorf("server start fail")
	}

	return nil
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
