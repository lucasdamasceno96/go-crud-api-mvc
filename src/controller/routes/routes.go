package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdamasceno96/go-crud-api-mvc/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createUser", controller.CreateUser)
	r.PUT("/updateUserById/:userId", controller.UpdateUserById)
	r.DELETE("/deleteUserById/:userId", controller.DeleteUserById)
}
