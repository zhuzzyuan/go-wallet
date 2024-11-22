package wallet

import (
	"go-wallet/render"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

var supportedChainList = []string{"ethereum"}

func bindJSON(c *gin.Context, req any) bool {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, render.BindError(err))
		return false
	}

	return true
}

func bindQuery(c *gin.Context, req any) bool {
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(400, render.BindError(err))
		return false
	}

	return true
}

func isSupportedChain(chain string) bool {
	isSupported := false

	for _, supportedChain := range supportedChainList {
		if strings.EqualFold(chain, supportedChain) {
			isSupported = true
		}
	}

	return isSupported
}

func isValidAddress(raw string) bool {
	addr := common.HexToAddress(raw)
	return strings.EqualFold(addr.Hex(), raw)
}
