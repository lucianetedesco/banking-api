package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianetedesco/banking-api/core"
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/repositories"
	"github.com/lucianetedesco/banking-api/usecases"
	"net/http"
	"strconv"
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

func GetAccounts(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryAccount := repositories.NewAccountRepository(d.Db)
	useCaseAccount := usecases.NewAccountUseCase(repositoryAccount)

	accounts, err := useCaseAccount.GetAllAccounts()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func GetBalanceAccount(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryAccount := repositories.NewAccountRepository(d.Db)
	useCaseAccount := usecases.NewAccountUseCase(repositoryAccount)

	accountId := c.Param("account_id")

	u64, _ := strconv.ParseUint(accountId, 10, 32)
	uAccountId := uint(u64)

	balance, err := useCaseAccount.GetBalanceAccount(uAccountId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return

	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func AuthUser(c *gin.Context) {
	d := core.GetDatabaseConnectionInstance()
	repositoryAccount := repositories.NewAccountRepository(d.Db)
	useCaseAccount := usecases.NewAccountUseCase(repositoryAccount)

	var login entities.Login

	if err := c.ShouldBind(&login); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := useCaseAccount.GetAccount(login)
	if err != nil {
		if err.Error() == "account not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "incorrect secret" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
