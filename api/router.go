package api

import (
	"go-wallet/api/wallet"

	"github.com/gin-gonic/gin"
)

func EnableRouter(r *gin.Engine) {

	wallet.EnableRouter(r.Group("/wallet"))

	// token
	r.POST("/token", GetToken)
	r.POST("/refresh_token", nil) // TODO refresh token
}
