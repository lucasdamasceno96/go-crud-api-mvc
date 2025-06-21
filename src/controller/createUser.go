package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/config/rest_err"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There some incorrect fields, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
