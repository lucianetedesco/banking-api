package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianetedesco/banking-api/controllers"
)

func CreateRoutesAndRun() {
	router := gin.Default()

	router.POST("/accounts", controllers.SaveAccount)
	router.GET("/accounts", controllers.GetAccounts)
	router.GET("/accounts/:account_id/balance", controllers.GetBalanceAccount)

	router.Run()
}
