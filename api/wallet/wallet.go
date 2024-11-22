package wallet

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func EnableRouter(rg *gin.RouterGroup) {
	rg.POST("/deposit", Deposit)
	rg.POST("/withdraw", Withdraw)
	rg.POST("/transfer", Transfer)
	rg.GET("/balance", Balance)
	rg.GET("/tx_history", TxHistory)

	rg.GET("/hello", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("hello:%s", time.Now().Format(time.RFC3339)))
	})
}
