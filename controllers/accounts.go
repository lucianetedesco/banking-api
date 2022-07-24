package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianetedesco/banking-api/core"
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/repositories"
	"github.com/lucianetedesco/banking-api/usecases"
	"net/http"
)

func SaveAccount(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryAccount := repositories.NewAccountRepository(d.Db)
	useCaseAccount := usecases.NewAccountUseCase(repositoryAccount)

	var account entities.Account

	if err := c.ShouldBind(&account); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := account.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := useCaseAccount.SaveAccount(account)
	if err != nil {
		if err.Error() == "CPF already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account_id": id})
}
