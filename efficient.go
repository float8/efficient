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
	var router Engine
	if len(Config.Middleware) > 0 {
		router = gin.New()
		router.Use(Config.Middleware...)
	} else {
		router = gin.Default()
	}
	registerRouters(router)
	router.Run(Config.Addr)
}
