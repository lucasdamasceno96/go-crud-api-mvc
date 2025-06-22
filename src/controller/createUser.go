package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/config/logger"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/config/validation"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/model/request"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro trying to validate user info", err)
		restErr := validation.ValidationUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
