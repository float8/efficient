package efficient

import (
	"github.com/gin-gonic/gin"
)

type (
	Context    = *gin.Context
	Engine     = *gin.Engine
	Middleware = gin.HandlerFunc
)

func Run() {
	if !Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = Logger.Out
	gin.DefaultErrorWriter = Logger.Out

	var router Engine
	router = gin.New()
	router.Use(ginLogger)
	router.Use(Config.Middleware...)
	registerRouters(router)

	if err := router.Run(Config.Addr); err != nil {
		Log.Fatal(err)
	}
}
