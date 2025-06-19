package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/config/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Wrong method !")
	c.JSON(err.Code, err)
}
