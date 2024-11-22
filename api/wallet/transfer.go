package wallet

import (
	"fmt"
	"go-wallet/api/models"
	"go-wallet/db"
	"go-wallet/render"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	var req models.TransferRequest
	if !bindJSON(c, &req) {
		return
	}

	if !isSupportedChain(req.Chain) {
		c.JSON(200, render.Error(fmt.Errorf("chain:%s is not supported yet", req.Chain)))
		return
	}

	resp, err := db.Transfer(req.Email, req.DestinationEmail, req.Chain, req.CoinType, req.Value)
	if err != nil {
		c.JSON(200, render.Error(err))
		return
	}

	c.JSON(200, render.Success(resp))
}
