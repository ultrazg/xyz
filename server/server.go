package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	U "github.com/ultrazg/xyz/pkg/utils"
	"github.com/ultrazg/xyz/routes"
	"log"
)

func Start() {
	p, d := U.InitFlag()
	port := fmt.Sprintf("%d", p)

	U.P(port)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

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
