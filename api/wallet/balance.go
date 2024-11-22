package wallet

import (
	"go-wallet/api/models"
	"go-wallet/db"
	"go-wallet/render"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {
	var req models.BalanceRequest
	if !bindQuery(c, &req) {
		return
	}

	resp, err := db.GetBalances(req.Email)
	if err != nil {
		c.JSON(200, render.Error(err))
		return
	}

	c.JSON(200, render.Success(resp))
}
