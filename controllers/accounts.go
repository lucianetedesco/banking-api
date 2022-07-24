package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAccount(c *gin.Context) {

	c.JSON(http.StatusCreated, "created")
}
