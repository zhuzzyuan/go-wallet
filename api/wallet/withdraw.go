package wallet

import (
	"fmt"
	"go-wallet/api/models"
	"go-wallet/db"
	"go-wallet/render"

	"github.com/gin-gonic/gin"
)

func Withdraw(c *gin.Context) {
	var req models.WithdrawRequest
	if !bindJSON(c, &req) {
		return
	}

	if !isSupportedChain(req.Chain) {
		c.JSON(200, render.Error(fmt.Errorf("chain:%s is not supported yet", req.Chain)))
		return
	}

	if !isValidAddress(req.Destination) {
		c.JSON(200, render.Error(fmt.Errorf("invalid address:%s", req.Destination)))
		return
	}

	resp, err := db.Withdraw(req)
	if err != nil {
		c.JSON(200, render.Error(err))
		return
	}

	c.JSON(200, render.Success(resp))
}
