package api

import (
	"go-wallet/config"
	"go-wallet/util/color"
	"go-wallet/util/log"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := newRouter()
	EnableRouter(r)

	addr := config.GetAPIEndpoint()
	log.Info(color.BGreenf("API listening at %s", addr))

	err := r.Run(addr)
	if err != nil {
		log.Panic(err)
	}
}

func newRouter() *gin.Engine {
	gin.SetMode(config.Mode())
	if config.Mode() == config.ReleaseMode {
		gin.DisableConsoleColor()
	}

	gin.DefaultErrorWriter = log.GetErrorLogger().Out

	r := gin.Default()

	r.MaxMultipartMemory = 1 << 20 // 1 MiB
	r.HandleMethodNotAllowed = true
	r.RedirectTrailingSlash = false

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Panic(err)
	}

	return r
}
