package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianetedesco/banking-api/controllers"
)

func CreateRoutesAndRun() {
	router := gin.Default()

	router.POST("/accounts", controllers.SaveAccount)

	router.Run()
}
