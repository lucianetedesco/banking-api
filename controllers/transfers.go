package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lucianetedesco/banking-api/core"
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/repositories"
	"github.com/lucianetedesco/banking-api/usecases"
	"net/http"
	"strings"
)

func SaveTransfer(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryTransfer := repositories.NewTransferRepository(d.Db)
	useCaseTransfer := usecases.NewTransferUseCase(repositoryTransfer)

	var transfer entities.Transfer

	if err := c.ShouldBind(&transfer); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := getToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	id, err := useCaseTransfer.SaveTransfer(transfer, token)
	if err != nil {
		if err == errors.New("user unauthorized") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if err == errors.New("amount in account must be greater than 0") {
			c.JSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction-id": id})
}

func GetTransfers(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryTransfer := repositories.NewTransferRepository(d.Db)
	useCaseTransfer := usecases.NewTransferUseCase(repositoryTransfer)

	token, err := getToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}

	transfers, err := useCaseTransfer.GetTransfers(token)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, transfers)
}

func getToken(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("no Authorization header provided")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return "", errors.New("could not find bearer token in Authorization header")
	}

	return token, nil
}
