package efficient

import (
	"github.com/gin-gonic/gin"
)

type (
	Context    = *gin.Context
	Engine     = *gin.Engine
	Middleware = gin.HandlerFunc
)

func WebRun() {
	if !EConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var router Engine
	if len(EConfig.Middleware) > 0 {
		router = gin.New()
		router.Use(EConfig.Middleware...)
	} else {
		router = gin.Default()
	}
	registerRouters(router)
	router.Run(EConfig.Addr)
}

func CMDRun() {

}
