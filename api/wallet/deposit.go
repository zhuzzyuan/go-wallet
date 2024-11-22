package wallet

import (
	"fmt"
	"go-wallet/api/models"
	"go-wallet/db"
	"go-wallet/render"

	"github.com/gin-gonic/gin"
)

func Deposit(c *gin.Context) {
	var req models.DepositRequest
	if !bindJSON(c, &req) {
		return
	}

	if !isSupportedChain(req.Chain) {
		c.JSON(200, render.Error(fmt.Errorf("chain:%s is not supported yet", req.Chain)))
		return
	}

	resp, err := db.GetAddressByEmailAndChain(req.Email, req.Chain)
	if err != nil {
		c.JSON(200, render.Error(err))
		return
	}

	c.JSON(200, render.Success(resp))
}
